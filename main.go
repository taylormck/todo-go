package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func main() {
	InitDatabase()
	defer DB.Close()

	engine := gin.Default()
	engine.LoadHTMLGlob("templates/*")

	engine.GET("/", func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"todos": ReadTodoList(),
		})
	})

	engine.POST("/todos", func(ctx *gin.Context) {
		title := ctx.PostForm("title")
		status := ctx.PostForm("status")
		id, _ := CreateTodo(title, status)

		ctx.HTML(http.StatusOK, "task.html", gin.H{
			"title":  title,
			"status": status,
			"id":     id,
		})
	})

	engine.DELETE("/todos/:id", func(ctx *gin.Context) {
		param := ctx.Param("id")
		id, _ := strconv.ParseInt(param, 10, 64)
		DeleteTodo(id)
	})

	engine.Run(":3000")
}
