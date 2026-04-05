// Задание 5: Несколько ошибок сразу - errors.Join и multierror
//
// Иногда нужно собрать сразу несколько ошибок.
// Например: провалидировать форму и вернуть все проблемы, а не только первую.
//
// Go 1.20 добавил errors.Join - соединяет ошибки в одну:
//
//   err := errors.Join(
//       errors.New("поле name пустое"),
//       errors.New("поле email: нет @"),
//   )
//   fmt.Println(err)
//   // поле name пустое
//   // поле email: нет @
//
// Важно: errors.Join(nil, nil) → nil  (если все ошибки nil - результат nil)
//
// errors.Is и errors.As работают через errors.Join:
//   err := errors.Join(ErrNotFound, someOtherError)
//   errors.Is(err, ErrNotFound)  // true!
//
// ──────────────────────────────────────────────────────────────────────────────
//
// Создай тип ValidationError:
//   type ValidationError struct {
//       Field   string
//       Message string
//   }
//   func (e *ValidationError) Error() string → `поле "name": не должно быть пустым`
//
// Напиши функцию validateOrder(name string, amount int, email string) error:
//   Собирает все ошибки через errors.Join (не возвращает при первой!):
//   - name пустой   → &ValidationError{Field: "name", Message: "не должно быть пустым"}
//   - amount <= 0   → &ValidationError{Field: "amount", Message: "должно быть больше 0"}
//   - нет "@" в email → &ValidationError{Field: "email", Message: "неверный формат"}
//   Возвращает errors.Join всех найденных ошибок.
//
// Напиши функцию printErrors(err error):
//   - если err == nil → fmt.Println("валидация прошла успешно")
//   - иначе выводит каждую отдельную ошибку через errors.As в цикле.
//
//   Подсказка - как разобрать errors.Join:
//     type multiErr interface { Unwrap() []error }
//     var me multiErr
//     if errors.As(err, &me) {
//         for _, e := range me.Unwrap() {
//             var vErr *ValidationError
//             if errors.As(e, &vErr) {
//                 fmt.Printf("  поле %q: %s\n", vErr.Field, vErr.Message)
//             }
//         }
//     }
//
// Ожидаемый вывод:
//
//   validateOrder("Аня", 5, "anya@mail.ru"):
//   валидация прошла успешно
//
//   validateOrder("", -1, "anya-without-at"):
//   ошибки валидации:
//     поле "name": не должно быть пустым
//     поле "amount": должно быть больше 0
//     поле "email": неверный формат
//
//   validateOrder("Боря", 0, ""):
//   ошибки валидации:
//     поле "amount": должно быть больше 0
//     поле "email": неверный формат
//
// Запусти: go run main.go

package main

import (
	"errors"
	"fmt"
	"strings"
)

// TODO: создай ValidationError { Field, Message string }
// TODO: реализуй Error() string

// TODO: напиши validateOrder(name string, amount int, email string) error

// TODO: напиши printErrors(err error)

func main() {
	// TODO: раскомментируй когда реализуешь validateOrder и printErrors
	//
	// fmt.Println(`validateOrder("Аня", 5, "anya@mail.ru"):`)
	// printErrors(validateOrder("Аня", 5, "anya@mail.ru"))
	//
	// fmt.Println()
	// fmt.Println(`validateOrder("", -1, "anya-without-at"):`)
	// printErrors(validateOrder("", -1, "anya-without-at"))
	//
	// fmt.Println()
	// fmt.Println(`validateOrder("Боря", 0, ""):`)
	// printErrors(validateOrder("Боря", 0, ""))

	_ = errors.Join
	_ = strings.Contains
	_ = fmt.Println
}
