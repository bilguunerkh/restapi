package main

import (
	"github.com/bilguunerkh/restapi"
	"net/http"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"title"`
	Completed bool   `json:"completed"`
}

var todos = []todo{
	{ID: "1", Item: "2023/08/02-nii nisleg tosoh", Completed: false},
	{ID: "2", Item: "toosoo archih", Completed: false},
	{ID: "3", Item: "rest api bichih daalgavar", Completed: false},
}

func getTodos(context *bilguunerkh.Context) {
	context.IndentedJSON(http.StatusOK, todos)
}

func main() {
	var router = bilguunerkh.Default()
	router.GET("/todos", getTodos)
	router.Run("https://localhost:8000")

}
