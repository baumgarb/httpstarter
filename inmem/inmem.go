package inmem

import (
	"sync"

	"github.com/baumgarb/httpstarter"
	"github.com/baumgarb/httpstarter/inmem/rand"
)

type enitity interface {
	Id() int
}

type inMemoryStore[T enitity] struct {
	sync.RWMutex

	setIdFn  func(T, int) T
	entities []T
}

func NewInMemoryStore[T enitity](seed []T, setIdFn func(T, int) T) *inMemoryStore[T] {
	return &inMemoryStore[T]{
		entities: seed,
		setIdFn:  setIdFn,
	}
}

func (s *inMemoryStore[T]) GetAll() ([]T, error) {
	s.RLock()
	defer s.RUnlock()

	// TODO: deep copy to avoid side effects?
	return s.entities, nil
}

func (s *inMemoryStore[T]) GetByID(id int) (T, error) {
	s.RLock()
	defer s.RUnlock()

	for _, e := range s.entities {
		if e.Id() == id {
			return e, nil
		}
	}
	var e T
	return e, httpstarter.ErrEntityNotFound
}

func (s *inMemoryStore[T]) Add(new T) (int, error) {
	s.Lock()
	defer s.Unlock()

	new = s.setIdFn(new, rand.NewID())
	s.entities = append(s.entities, new)
	return new.Id(), nil
}

func (s *inMemoryStore[T]) Update(new T) error {
	s.RLock()
	defer s.RUnlock()

	for i, e := range s.entities {
		if e.Id() == new.Id() {
			s.entities[i] = new
			return nil
		}
	}
	return httpstarter.ErrEntityNotFound
}
