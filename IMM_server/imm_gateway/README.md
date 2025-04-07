
---

## 🌐 主体结构

```go
func main() {
    ...
    http.HandleFunc("/", gateway) // 所有请求都交由 gateway 处理
    ...
}
```

网关监听的是配置文件 `settings.yaml` 中定义的地址（默认是 `127.0.0.1:8080`），通过 `gateway` 函数统一接收和处理 HTTP 请求。

---

## 🧠 核心逻辑：`gateway(res, req)`

```go
func gateway(res http.ResponseWriter, req *http.Request)
```

这个函数做了三件事：

### 1. **提取服务名称**

```go
regex, _ := regexp.Compile(`/api/(.*?)/`)
addrList := regex.FindStringSubmatch(req.URL.Path)
```

从 URL 路径中解析服务名，例如：`/api/user/login` 会提取出 `user`，然后查找 etcd 上注册的 `user_api` 地址：

```go
addr := etcd.GetServiceAddr(config.Etcd, service+"_api")
```

---

### 2. **调用认证服务**

```go
authUrl := fmt.Sprintf("http://%s/api/auth/authentication", authAddr)
if !auth(authUrl, res, req) {
    return
}
```

通过将请求原始 Header 转发给认证服务 `/api/auth/authentication`，由认证服务决定是否放行。

> 认证服务的实现对应 `authenticationHandler` 和 `AuthenticationLogic`，会去校验 JWT token 是否有效。

---

### 3. **请求代理转发**

```go
proxyUrl := fmt.Sprintf("http://%s%s", addr, req.URL.String())
proxy(proxyUrl, res, req)
```

认证通过后，继续调用目标服务的接口（例如 user_api），并将结果返回给客户端。

---

## 🔒 认证处理 `auth(...)`

```go
func auth(authAddr string, res http.ResponseWriter, req *http.Request) (ok bool)
```

- 构造一个 POST 请求，转发原始 header（包含 token）给认证服务。
- 根据认证服务返回的 code 是否为 0 决定是否通过。
- 若认证失败，返回自定义 JSON 错误响应（见 `FilResponse`）。

---

## 🔁 请求转发 `proxy(...)`

```go
func proxy(proxyAddr string, res http.ResponseWriter, req *http.Request)
```

- 将原请求的 Body 读取后，构造一个新的 HTTP 请求并转发到目标服务。
- 传递原始 Header，但移除认证中添加的 `ValidPath`。
- 返回服务响应内容给客户端。

---

## ⚙️ 配置文件解析

```go
type Config struct {
    Addr string
    Etcd string
    Log  logx.LogConf
}
```

用来加载 `settings.yaml`，包括服务监听地址、etcd 注册中心地址、日志格式等。

---

## ✅ 总结

`gateway.go` 实现的是一个 go-zero 微服务网关，其功能包含：

| 功能             | 描述 |
|------------------|------|
| 路由解析         | 根据路径中 `/api/xxx/` 动态解析服务名 |
| 服务发现         | 通过 etcd 获取对应服务地址 |
| 请求认证         | 所有请求在转发前先进行 token 验证 |
| 请求转发         | 将请求转发至目标服务并响应客户端 |
| 错误统一处理     | 错误都以统一结构体 `BaseResponse` 返回 |

如果你还想进一步增强功能，可以考虑：
- 支持白名单路径跳过认证；
- 支持负载均衡（多个服务实例）；
- 支持路由缓存；
- 支持请求限流/熔断。
