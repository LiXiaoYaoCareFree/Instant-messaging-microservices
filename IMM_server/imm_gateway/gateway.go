package main

import (
	"IMM_server/common/etcd"
	"bytes"
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

type BaseResponse struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func FilResponse(msg string, res http.ResponseWriter) {
	response := BaseResponse{Code: 7, Msg: msg}
	byteData, _ := json.Marshal(response)
	res.Write(byteData)
}

func auth(authAddr string, res http.ResponseWriter, req *http.Request) (ok bool) {
	authReq, _ := http.NewRequest("POST", authAddr, nil)

	authReq.Header = req.Header
	authReq.Header.Set("ValidPath", req.URL.Path)
	authRes, err := http.DefaultClient.Do(authReq)
	if err != nil {
		logx.Error(err)
		FilResponse("认证服务错误", res)
		return
	}

	type Response struct {
		Code int    `json:"code"`
		Msg  string `json:"msg"`
		Data *struct {
			UserID uint `json:"userID"`
			Role   int  `json:"role"`
		} `json:"data"`
	}
	var authResponse Response
	byteData, _ := io.ReadAll(authRes.Body)
	authErr := json.Unmarshal(byteData, &authResponse)
	if authErr != nil {
		logx.Error(authErr)
		FilResponse("认证服务错误", res)
		return
	}

	// 认证不通过
	if authResponse.Code != 0 {
		res.Write(byteData)
		return
	}
	if authResponse.Data != nil {
		req.Header.Set("User-ID", fmt.Sprintf("%d", authResponse.Data.UserID))
		req.Header.Set("Role", fmt.Sprintf("%d", authResponse.Data.Role))
	}
	return true
}

func proxy(proxyAddr string, res http.ResponseWriter, req *http.Request) {
	byteData, _ := io.ReadAll(req.Body)
	proxyReq, err := http.NewRequest(req.Method, proxyAddr, bytes.NewBuffer(byteData))

	if err != nil {
		logx.Error(err)
		FilResponse("err", res)
		return
	}

	proxyReq.Header = req.Header
	proxyReq.Header.Del("ValidPath")
	response, ProxyErr := http.DefaultClient.Do(proxyReq)
	if ProxyErr != nil {
		logx.Error(ProxyErr)
		FilResponse("服务异常", res)
		return
	}
	io.Copy(res, response.Body)
}

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
		logx.Errorf("%s 不匹配的服务", service)
		FilResponse("err", res)
		return
	}

	remoteAddr := strings.Split(req.RemoteAddr, ":")
	// 请求认证服务地址
	authAddr := etcd.GetServiceAddr(config.Etcd, "auth_api")
	authUrl := fmt.Sprintf("http://%s/api/auth/authentication", authAddr)
	proxyUrl := fmt.Sprintf("http://%s%s", addr, req.URL.String())

	// 打印日志
	logx.Infof("%s %s", remoteAddr[0], proxyUrl)

	if !auth(authUrl, res, req) {
		return
	}

	proxy(proxyUrl, res, req)
}

var configFile = flag.String("f", "settings.yaml", "the config file")

type Config struct {
	Addr string
	Etcd string
	Log  logx.LogConf
}

var config Config

func main() {
	flag.Parse()
	conf.MustLoad(*configFile, &config)

	logx.SetUp(config.Log)

	// 回调函数
	http.HandleFunc("/", gateway)
	fmt.Printf("gateway running %s\n", config.Addr)
	// 绑定服务
	http.ListenAndServe(config.Addr, nil)
}
