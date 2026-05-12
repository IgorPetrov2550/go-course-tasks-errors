// Задание 3: Цепочка ошибок repo → service → handler
//
// В реальных приложениях ошибка проходит через несколько слоёв.
// Каждый слой добавляет контекст через fmt.Errorf("%w", err).
// Верхний слой может найти исходную ошибку через errors.Is / errors.As.
//
// Схема:
//   handler → service → repository
//
//   repository возвращает ErrNotFound или ошибку с данными
//   service оборачивает: fmt.Errorf("service.GetUser id=%d: %w", id, err)
//   handler смотрит что пришло и выбирает HTTP-статус
//
// ──────────────────────────────────────────────────────────────────────────────
//
// Объяви:
//   var ErrNotFound = errors.New("not found")
//
// Создай тип User:
//   type User struct { ID int; Name string }
//
// Создай тип DBError - ошибка от базы данных:
//   type DBError struct {
//       Code    int    // код ошибки postgres, например 23505
//       Detail  string // детали
//   }
//   func (e *DBError) Error() string → `db error code=23505: нарушение уникальности`
//
// Напиши userRepository.FindByID(id int) (*User, error):
//   - id == 0  → return nil, fmt.Errorf("repo.FindByID: %w", ErrNotFound)
//   - id == -1 → return nil, fmt.Errorf("repo.FindByID: %w", &DBError{Code: 23505, Detail: "нарушение уникальности"})
//   - иначе   → return &User{ID: id, Name: "Аня"}, nil
//
// Напиши userService.GetUser(id int) (*User, error):
//   - вызывает repo.FindByID
//   - при ошибке: return nil, fmt.Errorf("service.GetUser id=%d: %w", id, err)
//
// Напиши handleGetUser(id int):
//   - вызывает service.GetUser
//   - если errors.Is(err, ErrNotFound)  → fmt.Println("404: пользователь не найден")
//   - если errors.As(err, &dbErr)       → fmt.Printf("500: db error code=%d\n", dbErr.Code)
//   - если err != nil                   → fmt.Println("500: внутренняя ошибка:", err)
//   - если nil                          → fmt.Printf("200: %s (id=%d)\n", user.Name, user.ID)
//
// Ожидаемый вывод:
//   200: Аня (id=7)
//   404: пользователь не найден
//   500: db error code=23505
//
// Обрати внимание:
//   errors.Is и errors.As находят оригинальную ошибку через две обёртки!
//   fmt.Println(err) выведет: service.GetUser id=0: repo.FindByID: not found
//   но errors.Is(err, ErrNotFound) всё равно вернёт true
//
// Запусти: go run main.go

package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

type User struct {
	ID   int
	Name string
}

// TODO: создай DBError { Code int; Detail string }

type DBError struct {
	Code   int
	Detail string
}

// TODO: реализуй Error() string

func (e *DBError) Error() string {
	return fmt.Sprintf("db error code=%d: %s", e.Code, e.Detail)
}

// TODO: создай userRepository с методом FindByID(id int) (*User, error)
type userRepository struct{}

func (r *userRepository) FindByID(id int) (*User, error) {
	// Напиши userRepository.FindByID(id int) (*User, error):
	//   - id == 0  → return nil, fmt.Errorf("repo.FindByID: %w", ErrNotFound)
	//   - id == -1 → return nil, fmt.Errorf("repo.FindByID: %w", &DBError{Code: 23505, Detail: "нарушение уникальности"})
	//   - иначе   → return &User{ID: id, Name: "Аня"}, nil
	//
	if id == 0 {
		return nil, fmt.Errorf("repo.FindByID: %w", ErrNotFound)
	}
	if id == -1 {
		return nil, fmt.Errorf("repo.FindByID: %w", &DBError{Code: 23505, Detail: "нарушение уникальности"})
	}
	return &User{ID: id, Name: "Аня"}, nil
}

// TODO: создай userService с полем repo и методом GetUser(id int) (*User, error)

type userService struct {
	repo *userRepository
}

func (r *userService) GetUser(id int) (*User, error) {
	//   - вызывает repo.FindByID
	//   - при ошибке: return nil, fmt.Errorf("service.GetUser id=%d: %w", id, err)
	if user, err := r.repo.FindByID(id); err != nil {
		return nil, fmt.Errorf("service.GetUser id=%d: %w", id, err)
	} else {
		return user, nil
	}
}

// TODO: напиши handleGetUser(id int)

func handleGetUser(id int) {
	//   - вызывает service.GetUser
	//   - если errors.Is(err, ErrNotFound)  → fmt.Println("404: пользователь не найден")
	//   - если errors.As(err, &dbErr)       → fmt.Printf("500: db error code=%d\n", dbErr.Code)
	//   - если err != nil                   → fmt.Println("500: внутренняя ошибка:", err)
	//   - если nil                          → fmt.Printf("200: %s (id=%d)\n", user.Name, user.ID)
	repo := &userRepository{}
	svc := &userService{repo: repo}
	user, err := svc.GetUser(id)
	var dbErr *DBError
	switch {
	case errors.Is(err, ErrNotFound):
		fmt.Println("404: пользователь не найден")

	case errors.As(err, &dbErr):
		fmt.Printf("500: db error code=%d\n", dbErr.Code)

	case err != nil:
		fmt.Println("500: внутренняя ошибка:", err)

	default:
		fmt.Printf("200: %s (id=%d)\n", user.Name, user.ID)
	}
}
func main() {
	// TODO: вызови handleGetUser(7)  → 200
	// TODO: вызови handleGetUser(0)  → 404
	// TODO: вызови handleGetUser(-1) → 500 db error

	handleGetUser(7)
	handleGetUser(0)
	handleGetUser(-1)
}
