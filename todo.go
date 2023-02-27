package httpstarter

import (
	"fmt"
	"time"
)

var EmptyTodo = Todo{}

type Todo struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Done bool   `json:"done"`
	ChangeTracking
}

func (t Todo) Id() int {
	return t.ID
}

type todoStore interface {
	GetAll() ([]Todo, error)
	GetByID(id int) (Todo, error)
	Add(t Todo) (id int, err error)
	Update(t Todo) error
}

type TodoService struct {
	store todoStore
}

func NewTodoService(ts todoStore) *TodoService {
	return &TodoService{
		store: ts,
	}
}

func (ts *TodoService) GetAll(u User) ([]Todo, error) {
	all, err := ts.store.GetAll()
	if err != nil {
		return []Todo{}, fmt.Errorf("todosvc: error when retrieving all: %w", err)
	}
	return all, nil
}

func (ts *TodoService) Add(t Todo, u User) (Todo, error) {
	// TODO: proper input validation
	t.LastModifiedAt = int(time.Now().UnixMilli())
	t.LastModifiedBy = u.Email

	id, err := ts.store.Add(t)
	if err != nil {
		return EmptyTodo, fmt.Errorf("todosvc: error when adding new '%v': %w", t.Name, err)
	}
	t.ID = id
	return t, nil
}

func (ts *TodoService) Update(id int, new Todo, u User) error {
	if id != new.ID {
		return NewInputValidationError("todosvc: provided IDs to not match (%v vs %v)", id, new.ID)
	}

	t, err := ts.store.GetByID(id)
	if err != nil {
		return err
	}

	// TODO: proper input validation
	t.Name = new.Name
	t.Done = new.Done
	t.LastModifiedAt = int(time.Now().UnixMilli())
	t.LastModifiedBy = u.Email
	err = ts.store.Update(t)
	if err != nil {
		return fmt.Errorf("todosvc: error when updating todo %v: %w", id, err)
	}
	return nil
}
