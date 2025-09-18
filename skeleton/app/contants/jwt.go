package contants

import "github.com/golang-jwt/jwt/v5"

type CustomClaims struct {
	Uid      int64            `json:"uid"`
	Source   string           `json:"source"`
	Subject  string           `json:"sub,omitempty"`
	Audience jwt.ClaimStrings `json:"aud,omitempty"`
	*jwt.RegisteredClaims
}
