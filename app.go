package main

import (
	"github.com/kataras/iris"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Serve using a host:port form.
var addr = iris.Addr("localhost:4000")

// Posts struct
type Posts struct {
	Title string `json:"title"`
	Content  string `json:"content"`
}

func main() {
	app := iris.New()

	app.Get("/wp-json/wp/v2/posts", iris.Gzip, posts)
	app.Get("/wp-json/wp/v2/posts/{id:int}", iris.Gzip, post)

	app.Run(addr)
}

func checkErr(err error){
    if err != nil {
        panic(err)
    }
}

// fetch single post from wordpress database
func post(ctx iris.Context){
	var id = ctx.Params().Get("id")
	var row = Posts{}

	db, err := sql.Open("mysql", "root:root@/wp_test?charset=utf8")
	checkErr(err)

	err = db.QueryRow("SELECT post_title, post_content FROM wp_posts WHERE id=? AND post_status = 'publish' AND post_type = 'post' LIMIT 1", id).Scan(&row.Title, &row.Content)
	checkErr(err)

    ctx.JSON(row)
}

// fetch mutli posts list from wordpress database
func posts(ctx iris.Context){
	var res = []Posts{}

	db, err := sql.Open("mysql", "root:root@/wp_test?charset=utf8")
	checkErr(err)

	rows, err := db.Query("SELECT post_title, post_content FROM wp_posts WHERE post_status = 'publish' AND post_type = 'post'")

	checkErr(err)
	defer rows.Close()

    for i:=0; rows.Next(); i++ {
		var row = Posts{}
        err := rows.Scan(&row.Title, &row.Content)
		checkErr(err)
		res = append(res, row)
    }
	ctx.JSON(res)
}
