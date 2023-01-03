package util

import (
	"github.com/labstack/echo-jwt/v4"
)

func init() {
	echojwt.WithConfig(
		echojwt.Config{
		SigningKey: []byte("secret"),	
		},
	)
}
