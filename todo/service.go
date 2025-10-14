package todo

import "fmt"

var todos []Todo = []Todo{
	{id: 1, title: "first work", complete: false},
	{id: 2, title: "second work", complete: true},
}

func AddTodo(title string) string{
	for _, v := range todos{
		if v.title == title{
			return fmt.Sprintf("Туду с таким текстом уже существует")
		}
	}
	var todo Todo = Todo {
		id: len(todos)+1,
		title: title,
		complete: false,
	}
	todos = append(todos, todo)
	return fmt.Sprintf("Добавлено дело")
}

func GetAllTodo() []Todo{
	return todos
}

func CompleteTodo(id int) bool{
	for i, comp := range todos{
		if comp.id == id{
			todos[i].complete = true
		}
	}
	return true
}

func DeleteTodo(id int) bool{
	for i, delTodo := range todos{
		if delTodo.id == id{
			todos = append(todos[:i], todos[i+1:]...)
		}
	}
	return true
}

