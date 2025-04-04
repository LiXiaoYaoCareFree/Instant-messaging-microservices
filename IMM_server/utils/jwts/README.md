这份 JWT 的实现代码是一个使用 `github.com/golang-jwt/jwt/v4` 库生成和解析 Token 的示例。

---

## 🧱 数据结构部分（在 `enter.go` 中）

```go
type JwtPayLoad struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
}
```

这是你自己定义的 Payload，用来存储用户信息，比如用户 ID、用户名和角色权限。

```go
type CustomClaims struct {
	JwtPayLoad
	jwt.RegisteredClaims
}
```

你自定义的 `CustomClaims` 结构体，嵌入了上面定义的 `JwtPayLoad`，并组合了官方定义的标准声明 `RegisteredClaims`（例如过期时间 `ExpiresAt`）。

---

## 🔐 Token 生成函数 `GenToken`

```go
func GenToken(payload JwtPayLoad, accessSecret string, expires int) (string, error) {
	claim := CustomClaims{
		JwtPayLoad: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(expires))),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString([]byte(accessSecret))
}
```

这个函数主要完成以下几个步骤：

1. 构造自定义的 `CustomClaims`。
2. 设置过期时间（当前时间 + 传入的小时数）。
3. 使用 `HS256` 签名算法创建 token。
4. 调用 `SignedString` 生成最终的字符串 token，传入的是密钥 `accessSecret`。

---

## 🔍 Token 解析函数 `ParseToken`

```go
func ParseToken(tokenStr string, accessSecret string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessSecret), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
```

这段代码：

1. 使用 `jwt.ParseWithClaims` 解析 token，传入的类型是 `&CustomClaims{}`。
2. 提供一个回调函数返回签名密钥 `accessSecret`。
3. 检查解析结果是否合法，并返回 `CustomClaims`。

---

## 🧪 测试代码 `jwt_test.go`

```go
func TestGenToken(t *testing.T) {
	token, err := GenToken(JwtPayLoad{
		UserID:   1,
		Role:     1,
		Username: "fengfeng",
	}, "12345", 8)
	fmt.Println(token, err)
}
```

测试生成 token，使用密钥 "12345"，有效期 8 小时。

```go
func TestParseToken(t *testing.T) {
	payload, err := ParseToken("xxx", "12345")
	fmt.Println(payload, err)
}
```

测试解析一个 token，注意里面的 token 是已经硬编码的字符串。如果你用 `TestGenToken` 生成后，拷贝那个 token 来这里测试会更准确。

---
