package memory

import (
	"fmt"

	"github.com/kmjayadeep/go-hexagonal-example/shortner"
	"github.com/pkg/errors"
)

type memoryRepository struct {
	store map[string]shortner.Redirect
}

func NewMemoryRepository() (*memoryRepository, error) {
	store := make(map[string]shortner.Redirect)
	return &memoryRepository{
		store: store,
	}, nil
}

func (m *memoryRepository) Find(code string) (*shortner.Redirect, error) {
	r, ok := m.store[code]

	fmt.Println("finding key", code)

	if !ok {
		return nil, errors.Wrap(shortner.ErrRedirectNotFound, "repository.Redirect.Find")
	}

	fmt.Println("found key", code, r)

	return &r, nil
}

func (m *memoryRepository) Store(redirect *shortner.Redirect) error {
	m.store[redirect.Code] = *redirect
	fmt.Println("stored", redirect, "store", m.store)
	return nil
}
