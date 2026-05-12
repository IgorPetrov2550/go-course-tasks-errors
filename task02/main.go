// Задание 2: Типизированная ошибка и errors.As
//
// errors.Is - проверяет "это та же ошибка?".
// errors.As - спрашивает "это ошибка такого типа? Дай её мне!"
//
// Зачем errors.As?
//
//   У ошибки могут быть дополнительные поля - статус, поле, код.
//   errors.As позволяет достать эти поля даже если ошибка обёрнута.
//
//   Пример - НЕ работает:
//     var vErr *ValidationError
//     if err.(*ValidationError) ...  // паника если тип другой
//                                    // и не проходит через обёртки
//
//   Пример - работает:
//     var vErr *ValidationError
//     if errors.As(err, &vErr) {     // безопасно, проходит через %w-обёртки
//         fmt.Println(vErr.Field)
//     }
//
// ──────────────────────────────────────────────────────────────────────────────
//
// Создай тип ValidationError:
//   type ValidationError struct {
//       Field   string  // имя поля, например "email"
//       Message string  // что не так, например "неверный формат"
//   }
//
// Реализуй метод Error() string - строка вида:
//   поле "email": неверный формат
//
// Создай тип NotFoundError:
//   type NotFoundError struct {
//       Resource string  // например "user"
//       ID       int
//   }
//
// Реализуй Error() string - строка вида:
//   user с id=42 не найден
//
// Напиши функцию processRequest(action string) error:
//   - action == "validate" → вернуть fmt.Errorf("processRequest: %w",
//                              ValidationError{Field: "email", Message: "нет символа @"})
//   - action == "find"     → вернуть fmt.Errorf("processRequest: %w",
//                              &NotFoundError{Resource: "user", ID: 42})
//   - иначе                → nil
//
// В main() вызови processRequest три раза и через errors.As разбери каждую ошибку:
//   processRequest("validate"):
//     → поле "email": нет символа @  (достать ValidationError, вывести Field и Message)
//   processRequest("find"):
//     → user с id=42 не найден  (достать *NotFoundError, вывести Resource и ID)
//   processRequest("ok"):
//     → успех
//
// Ожидаемый вывод:
//   ошибка: processRequest: поле "email": нет символа @
//   это ValidationError - поле: email, причина: нет символа @
//
//   ошибка: processRequest: user с id=42 не найден
//   это NotFoundError - ресурс: user, id: 42
//
//   успех
//
// Запусти: go run main.go

package main

import (
	"errors"
	"fmt"
)

// TODO: создай тип ValidationError { Field, Message string }
// TODO: реализуй метод Error() string - `поле "email": нет символа @`

type ValidationError struct {
	Field   string
	Message string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf("поле %q: %s", v.Field, v.Message)
}

// TODO: создай тип NotFoundError { Resource string; ID int }

type NotFoundError struct {
	Resource string
	ID       int
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s c id=%d не найден", e.Resource, e.ID)
}

// TODO: реализуй метод Error() string - `user с id=42 не найден`

// TODO: напиши processRequest(action string) error
func processRequest(action string) error {
	//   - action == "validate" → вернуть fmt.Errorf("processRequest: %w",
	//                              ValidationError{Field: "email", Message: "нет символа @"})
	//   - action == "find"     → вернуть fmt.Errorf("processRequest: %w",
	//                              &NotFoundError{Resource: "user", ID: 42})
	//   - иначе                → nil
	//
	if action == "validate" {
		return fmt.Errorf("processRequest: %w", ValidationError{Field: "email", Message: "нет символа @"})
	}
	if action == "find" {
		return fmt.Errorf("processRequest: %w", &NotFoundError{Resource: "user", ID: 42})
	}
	return nil
}

func main() {
	// TODO: вызови processRequest("validate"), processRequest("find"), processRequest("ok")

	actions := []string{"validate", "find", "ok"}
	for _, action := range actions {
		if err := processRequest(action); err != nil {
			fmt.Printf("ошибка: %v\n", err)

			var vErr ValidationError
			var nErr *NotFoundError
			// В main() вызови processRequest три раза и через errors.As разбери каждую ошибку:
			//   processRequest("validate"):
			//     → поле "email": нет символа @  (достать ValidationError, вывести Field и Message)
			//   processRequest("find"):
			//     → user с id=42 не найден  (достать *NotFoundError, вывести Resource и ID)
			//   processRequest("ok"):
			//     → успех
			if errors.As(err, &vErr) {
				fmt.Printf("это ValidationError - поле: %s, причина: %s\n\n", vErr.Field, vErr.Message)
			} else if errors.As(err, &nErr) {
				fmt.Printf("это NotFoundError - ресурс: %s, id: %d\n\n", nErr.Resource, nErr.ID)
			}
		} else {
			fmt.Println("успех")
		}
	}
	// TODO: для каждого случая используй errors.As чтобы разобрать ошибку

}
