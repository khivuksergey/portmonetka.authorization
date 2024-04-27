package storage

import "github.com/khivuksergey/portmonetka.authorization/internal/core/port/repository"

type IDB interface {
	InitRepositoryManager() *repository.Manager
	Close() error
}
