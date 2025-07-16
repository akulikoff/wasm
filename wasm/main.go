package main

import (
	"syscall/js"
)

var todos []string

func addTodo(this js.Value, args []js.Value) interface{} {
	if len(args) > 0 {
		todos = append(todos, args[0].String())
	}
	return nil
}

func getTodos(this js.Value, args []js.Value) interface{} {
	arr := js.Global().Get("Array").New()
	for i, todo := range todos {
		arr.SetIndex(i, todo)
	}
	return arr
}

func main() {
	js.Global().Set("addTodo", js.FuncOf(addTodo))
	js.Global().Set("getTodos", js.FuncOf(getTodos))
	select {} // Блокируем main, чтобы wasm не завершился
} 