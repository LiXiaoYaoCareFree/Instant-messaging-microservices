package jwts

import (
	"fmt"
	"testing"
)

func TestGenToken(t *testing.T) {
	token, err := GenToken(JwtPayLoad{
		UserID:   1,
		Role:     1,
		Username: "lly",
	}, "12345", 8)
	fmt.Println(token, err)
}

func TestParseToken(t *testing.T) {
	payload, err := ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ1c2VybmFtZSI6ImxseSIsInJvbGUiOjEsImV4cCI6MTc0Mzg1NjY0MH0.NEofc7eLal3uLEkz3R2D3p_hBDYBw9hLEjJBfpAb0YI", "12345")
	fmt.Println(payload, err)
}
