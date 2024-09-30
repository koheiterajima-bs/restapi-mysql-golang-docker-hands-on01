// Goの標準ライブラリnet/httpを用いてREST APIを実装

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/koheiterajima-bs/restapi-mysql-golang-docker-hands-on01/backend/handlers"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// httpリクエストが送られてきたときに、それを処理するためのルートとハンドラ関数を設定
	http.HandleFunc("/todos", handlers.HandleTodos) // handlers.HandleTodos関数がリクエストを処理する
	http.HandleFunc("/todos/", handlers.HandleTodo)

	fmt.Println("Server is running on port 8080...")

	// HTTPサーバーはポート8080でリクエストの待ち受けを開始する
	log.Fatal(http.ListenAndServe(":8080", nil))
}
