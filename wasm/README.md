# ToDo List на Go + WebAssembly

## Описание

Приложение демонстрирует работу ToDo-листа с хранением задач на сервере (Go backend) и клиентской частью на Go (WebAssembly).

---

## Структура проекта

- `main.go` — Go-код для WebAssembly (клиент)
- `server.go` — Go backend-сервер (API + раздача статики)
- `index.html` — фронтенд
- `main.wasm` — собранный wasm-файл (создаётся автоматически)
- `wasm_exec.js` — JS-рантайм для Go WASM (скопируйте из вашей установки Go)

---

## Запуск

### 1. Сборка WebAssembly

Выполните в папке проекта:

**Windows PowerShell:**

```powershell
$env:GOOS="js"; $env:GOARCH="wasm"; go build -o main.wasm main.go
```

**cmd.exe:**

```cmd
set GOOS=js
set GOARCH=wasm
go build -o main.wasm main.go
```

---

### 2. Запуск backend-сервера

```sh
cd <папка_проекта>
go run server.go
```

Сервер будет доступен на http://localhost:8080

---

### 3. Открытие приложения

Откройте в браузере:

```
http://localhost:8080/index.html
```

---

## Важно

- Файл `wasm_exec.js` возьмите из вашей установки Go (`<GOROOT>/misc/wasm/wasm_exec.js`).
- Все файлы (`index.html`, `main.go`, `main.wasm`, `wasm_exec.js`, `server.go`) должны лежать в одной папке.
- Для работы приложения должны быть запущены оба процесса: backend и wasm (main.wasm).

---

## Пример

```sh
cd C:/FE/wasm/wasm
$env:GOOS="js"; $env:GOARCH="wasm"; go build -o main.wasm main.go
go run server.go
```

---

## Возможные проблемы

- Если не открывается страница — проверьте, что сервер запущен из той же папки, где лежит index.html.
- Если не собирается wasm — проверьте, что используете Go 1.12+ и файл main.go не содержит серверного кода.
- Если не найден wasm_exec.js — скопируйте его из вашей установки Go.
