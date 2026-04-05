// Задание 1: Оборачивание ошибок и errors.Is
//
// В Go ошибки можно оборачивать - добавлять контекст на каждом уровне.
// Это как пересылать посылку: каждый добавляет свою наклейку поверх.
//
// Как оборачивать:
//   err := fmt.Errorf("service: %w", originalErr)
//   //                          ^^
//   //  %w (а не %v!) - сохраняет оригинальную ошибку внутри
//   //  без %w - errors.Is не сможет найти оригинал
//
// Как проверять через обёртки:
//   errors.Is(err, ErrNotFound)  →  true, даже если err обёрнута 3 раза
//
// ──────────────────────────────────────────────────────────────────────────────
//
// Объяви три sentinel error:
//   var ErrNotFound   = errors.New("not found")
//   var ErrForbidden  = errors.New("forbidden")
//   var ErrBadRequest = errors.New("bad request")
//
// Напиши функцию getOrder(userID, orderID int) error:
//   - если userID == 0   → вернуть fmt.Errorf("getOrder: %w", ErrBadRequest)
//   - если userID == 99  → вернуть fmt.Errorf("getOrder: %w", ErrForbidden)
//   - если orderID == 0  → вернуть fmt.Errorf("getOrder: %w", ErrNotFound)
//   - иначе nil
//
// Напиши функцию handleOrder(userID, orderID int), которая:
//   - вызывает getOrder
//   - оборачивает ошибку ещё раз: fmt.Errorf("handleOrder: %w", err)
//   - проверяет через errors.Is каждый тип и выводит разное сообщение:
//       ErrBadRequest → "400: неверный запрос"
//       ErrForbidden  → "403: нет доступа"
//       ErrNotFound   → "404: заказ не найден"
//       иначе         → "500: внутренняя ошибка"
//   - если ошибки нет - "200: заказ получен"
//
// Ожидаемый вывод:
//   200: заказ получен
//   400: неверный запрос
//   403: нет доступа
//   404: заказ не найден
//
// Запусти: go run main.go

package main

import (
	"errors"
	"fmt"
)

// TODO: объяви три sentinel error

// TODO: напиши функцию getOrder(userID, orderID int) error

// TODO: напиши функцию handleOrder(userID, orderID int)

func main() {
	// TODO: вызови handleOrder(1, 42)  → успех
	// TODO: вызови handleOrder(0, 42)  → bad request
	// TODO: вызови handleOrder(99, 42) → forbidden
	// TODO: вызови handleOrder(1, 0)   → not found

	_ = fmt.Println
	_ = errors.Is
}
