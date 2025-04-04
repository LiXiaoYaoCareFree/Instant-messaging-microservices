package jwts

import (
	"fmt"
	"testing"
)

func TestGenToken(t *testing.T) {
	token, err := GenToken(JwtPayLoad{
		UserID:   1,
		Role:     1,
		Username: "fengfeng",
	}, "12345", 8)
	fmt.Println(token, err)
}

func TestParseToken(t *testing.T) {
	payload, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjogMiwidXNlcm5hbWUiOiJmZW5nZmVuZyIsInJvbGUiOjEsImV4cCI6MTcwOTMyNTEwMX0=.O75mXE-COqplMozF_uL27GHNsghqWUG1sq8L0FpIyHw", "12345")
	fmt.Println(payload, err)
}
