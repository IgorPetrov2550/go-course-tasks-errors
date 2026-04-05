// Задание 4: HTTP-ошибки - http.Error vs JSON-ответ
//
// Два способа вернуть ошибку из HTTP-хендлера:
//
// Способ 1 - http.Error:
//   http.Error(w, "not found", http.StatusNotFound)
//   Content-Type: text/plain; charset=utf-8
//   Body: not found\n
//
// Способ 2 - вручную JSON:
//   w.Header().Set("Content-Type", "application/json")
//   w.WriteHeader(http.StatusNotFound)
//   json.NewEncoder(w).Encode(map[string]string{"error": "not found"})
//   Content-Type: application/json
//   Body: {"error":"not found"}
//
// Почему для API выбирают Способ 2:
//   - клиент (фронт, мобилка) ожидает JSON, а не plain text
//   - формат ошибок одинаковый: всегда {"error": "..."}
//   - можно добавить поля: {"error": "...", "code": "NOT_FOUND", "field": "id"}
//   - http.Error подходит для простых утилит/отладки
//
// ──────────────────────────────────────────────────────────────────────────────
//
// Напиши HTTP-сервер с двумя хендлерами:
//
// GET /v1/plain/{id}
//   - если id == "0" → http.Error(w, "not found", 404)
//   - иначе → http.Error(w, "user: "+id, 200)
//     (Подсказка: http.Error всегда text/plain)
//
// GET /v1/json/{id}
//   - используй writeJSON(w, status, body) и writeError(w, status, message)
//   - если id == "0" → writeError(w, 404, "not found")
//   - иначе → writeJSON(w, 200, map[string]string{"name": "user-" + id})
//
// Напиши хелперы:
//
//   func writeJSON(w http.ResponseWriter, status int, body any) {
//       w.Header().Set("Content-Type", "application/json")
//       w.WriteHeader(status)
//       json.NewEncoder(w).Encode(body)
//   }
//
//   func writeError(w http.ResponseWriter, status int, message string) {
//       writeJSON(w, status, map[string]string{"error": message})
//   }
//
// Запусти: go run main.go
// Проверь в терминале:
//
//   curl -i http://localhost:8080/v1/plain/0
//   curl -i http://localhost:8080/v1/plain/42
//   curl -i http://localhost:8080/v1/json/0
//   curl -i http://localhost:8080/v1/json/42
//
// Обрати внимание на Content-Type в ответе:
//   /v1/plain  → Content-Type: text/plain; charset=utf-8
//   /v1/json   → Content-Type: application/json
//
// Ожидаемые ответы:
//   GET /v1/plain/0   → 404, text/plain,       body: not found
//   GET /v1/plain/42  → 200, text/plain,        body: user: 42
//   GET /v1/json/0    → 404, application/json,  body: {"error":"not found"}
//   GET /v1/json/42   → 200, application/json,  body: {"name":"user-42"}

package main

import (
	"fmt"
	"net/http"
	// TODO: добавь "encoding/json" когда будешь реализовывать writeJSON
)

// TODO: напиши writeJSON(w http.ResponseWriter, status int, body any)

// TODO: напиши writeError(w http.ResponseWriter, status int, message string)

func main() {
	mux := http.NewServeMux()

	// TODO: зарегистрируй GET /v1/plain/{id}
	// TODO: зарегистрируй GET /v1/json/{id}

	fmt.Println("сервер запущен: http://localhost:8080")
	fmt.Println("попробуй: curl -i http://localhost:8080/v1/plain/0")
	fmt.Println("          curl -i http://localhost:8080/v1/json/42")

	http.ListenAndServe(":8080", mux)
}
