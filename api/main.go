// main.go

package main

import (
	"github.com/gouser/api/app"
)

func main() {
	app.StartApp()
	// // CORSの設定
	// app.Use(cors.New(cors.Config{
	// 	// 認証にcookieなどの情報を必要とするかどうか
	// 	AllowCredentials: true,
	// }))
}
