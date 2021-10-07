package lesson4

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"src/github.com/golang/glog"
)

func main() {

	//定义一个http的多路复用对象
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthzHandler)

	//服务监听端口
	http.ListenAndServe(":80", mux)
}

func healthzHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(200)
	io.WriteString(res, "OK")
	//记录请求IP地址
	tarckLog(req)
}

//定义根路径handler
func rootHandler(res http.ResponseWriter, req *http.Request) {

	fmt.Println("root handler processing.")

	//把request头的参数写到response中
	for k, _ := range req.Header {
		v := req.Header.Get(k)
		res.Header().Set(k, v)
	}

	//获取当前系统的环境变量,写到response的header中
	vers := os.Getenv("VERSION")
	res.Header().Set("VERSION", vers)

	//记录请求IP地址
	tarckLog(req)
}

func tarckLog(req *http.Request) {
	ipAddr := req.RemoteAddr
	glog.V(1).Infoln(fmt.Printf("client ip:%s", ipAddr))
}
