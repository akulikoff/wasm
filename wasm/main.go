package main

import (
	"syscall/js"
)

func fetchTodos(this js.Value, args []js.Value) interface{} {
	promise := js.Global().Get("Promise").New(js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		resolve := args[0]
		js.Global().Call("fetch", "/todos").Call("then", js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
			resp := args[0]
			return resp.Call("json")
		})).Call("then", js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
			json := args[0]
			resolve.Invoke(json)
			return nil
		}))
		return nil
	}))
	return promise
}

func postTodo(this js.Value, args []js.Value) interface{} {
	text := args[0].String()
	promise := js.Global().Get("Promise").New(js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		resolve := args[0]
		js.Global().Call("fetch", "/todos", map[string]interface{}{
			"method": "POST",
			"headers": map[string]interface{}{"Content-Type": "application/json"},
			"body": js.Global().Get("JSON").Call("stringify", map[string]interface{}{"text": text}),
		}).Call("then", js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
			resolve.Invoke(js.Undefined())
			return nil
		}))
		return nil
	}))
	return promise
}

func main() {
	js.Global().Set("fetchTodos", js.FuncOf(fetchTodos))
	js.Global().Set("postTodo", js.FuncOf(postTodo))
	select {}
} 