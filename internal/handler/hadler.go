package handler

import (
	"github.com/gorilla/sessions"
	"github.com/kyosu-1/ims/internal/repository"
)

type Handlers struct {
	userRepository     repository.UserRepository
	categoryRepository repository.CategoryRepository
	itemRepository     repository.ItemRepository
	sessionStore       sessions.Store
}

func NewHandlers(
	userRepository repository.UserRepository,
	categoryRepository repository.CategoryRepository,
	itemRepository repository.ItemRepository,
	sessionStore sessions.Store,
) *Handlers {
	return &Handlers{
		userRepository:     userRepository,
		categoryRepository: categoryRepository,
		itemRepository:     itemRepository,
		sessionStore:       sessionStore,
	}
}
