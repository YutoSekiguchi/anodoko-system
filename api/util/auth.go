package util

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"github.com/labstack/echo/v4"

	"google.golang.org/api/option"
)

// Private APIを認証付きにする
func Private(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Firebase SDK のセットアップ
		opt := option.WithCredentialsFile(os.Getenv("JWT_SECRET_PATH"))
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		auth, err := app.Auth(context.Background())
		if err != nil {
			fmt.Printf("error: %v\n", err)
			os.Exit(1)
		}
		// クライアントから送られてきた JWT 取得
		authHeader := c.Request().Header.Get("Authorization")
		idToken := strings.Replace(authHeader, "Bearer ", "", 1)
		// JWT の検証
		_, err = auth.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			return c.JSON(http.StatusUnauthorized, "unauthorized")
		} else {
			return next(c)
		}
	}
}