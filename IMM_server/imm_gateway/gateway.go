package main

import (
	"IMM_server/common/etcd"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"io"
	"net/http"
	"regexp"
	"strings"
)

func gateway(res http.ResponseWriter, req *http.Request) {
	// 匹配请求前缀  /api/user/xx
	regex, _ := regexp.Compile(`/api/(.*?)/`)
	addrList := regex.FindStringSubmatch(req.URL.Path)
	if len(addrList) != 2 {
		res.Write([]byte("err"))
		return
	}
	service := addrList[1]

	addr := etcd.GetServiceAddr(config.Etcd, service+"_api")
	if addr == "" {
		fmt.Println("不匹配的服务", service)
		res.Write([]byte("err"))
		return
	}

	remoteAddr := strings.Split(req.RemoteAddr, ":")
	fmt.Println(remoteAddr)

	// 请求认证服务地址
	authAddr := etcd.GetServiceAddr(config.Etcd, "auth_api")
	authUrl := fmt.Sprintf("http://%s/api/auth/authentication", authAddr)
	authReq, _ := http.NewRequest("POST", authUrl, nil)
	authReq.Header = req.Header
	authReq.Header.Set("ValidPath", req.URL.Path)
	authRes, err := http.DefaultClient.Do(authReq)
	if err != nil {
		res.Write([]byte("认证服务错误"))
		return
	}

	type Response struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
	}
	var authResponse Response
	byteData, _ := io.ReadAll(authRes.Body)
	authErr := json.Unmarshal(byteData, &authResponse)
	if authErr != nil {
		logx.Error(authErr)
		res.Write([]byte("认证服务错误"))
		return
	}

	// 认证不通过
	if authResponse.Code != 0 {
		res.Write(byteData)
		return
	}

	url := fmt.Sprintf("http://%s%s", addr, req.URL.String())
	fmt.Println(url)

	proxyReq, err := http.NewRequest(req.Method, url, req.Body)
	if err != nil {
		logx.Error(err)
		res.Write([]byte("err"))
		return
	}
	fmt.Println(proxyReq)
	proxyReq.Header.Set("X-Forwarded-For", remoteAddr[0])
	response, ProxyErr := http.DefaultClient.Do(proxyReq)
	if ProxyErr != nil {
		fmt.Println(ProxyErr)
		res.Write([]byte("服务异常"))
		return
	}
	io.Copy(res, response.Body)
}

var configFile = flag.String("f", "settings.yaml", "the config file")

type Config struct {
	Addr string
	Etcd string
}

var config Config

func main() {
	flag.Parse()
	conf.MustLoad(*configFile, &config)

	// 回调函数
	http.HandleFunc("/", gateway)
	fmt.Printf("gateway running %s\n", config.Addr)
	// 绑定服务
	http.ListenAndServe(config.Addr, nil)
}
