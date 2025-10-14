package main

import (
	"fmt"

	"guthub.com/IsmaCHup/goModTodolisRecover/todo"
)

func main() {
	fmt.Println(todo.AddTodo("second Todo"))

	fmt.Println(todo.AddTodo("second Todo"))

	fmt.Println(todo.GetAllTodo())

	fmt.Println(todo.CompleteTodo(2))

	fmt.Println(todo.GetAllTodo())


	fmt.Println(todo.DeleteTodo(1))
	fmt.Println(todo.GetAllTodo())

}


