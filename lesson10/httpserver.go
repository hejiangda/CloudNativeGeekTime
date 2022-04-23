package main

import (
	"context"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"
	"time"
)

//接收客户端 request，并将 request 中带的 header 写入 response header
//读取当前系统的环境变量中的 VERSION 配置，
//并写入 response header
//Server 端记录访问日志包括客户端 IP，
//HTTP 返回码，输出到 server 端的标准输出
//当访问 localhost/healthz 时，
//应返回200

func index(w http.ResponseWriter, r *http.Request) {
	// 添加0～2秒延时
	timer := NewTimer()
	defer timer.ObserveTotal()
	dura := rand.Intn(2000)
	time.Sleep(time.Duration(dura) * time.Millisecond)

	//读取当前系统的环境变量中的 VERSION 配置，
	//并写入 response header
	os.Setenv("VERSION", "v1")
	version := os.Getenv("VERSION")

	w.Header().Set("VERSION", version)

	for k, v := range r.Header {
		for _, vv := range v {
			//fmt.Println(vv)
			w.Header().Set(k, vv)
		}
	}

	//Server 端记录访问日志包括客户端 IP，
	//HTTP 返回码，输出到 server 端的标准输出
	clientIP := ClientIP(r)
	log.Println("clientIP:", clientIP)
	log.Println("client response code:", 200)
	fmt.Fprintf(w, "dura:%v\n", dura)
}

func getCurrentIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}

// ClientIP 尽最大努力实现获取客户端 IP 的算法。
//解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

//当访问 localhost/healthz 时，
//应返回200

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "working")
}

func main() {

	Register()

	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)
	mux.Handle("/metrics", promhttp.Handler())

	PORT := os.Getenv("httpport")
	srv := &http.Server{Addr: ":" + PORT, Handler: mux}

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	wg := sync.WaitGroup{}

	go func() {
		wg.Add(1)
		sig := <-sigCh
		log.Println("Got signal ", sig, ", attempting graceful shutdown!")
		if err := srv.Shutdown(context.TODO()); err != nil {
			panic(err)
		}
		wg.Done()
	}()
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
	wg.Wait()
}
