package storage

import "github.com/khivuksergey/portmonetka.authorization/internal/core/port/repository"

//go:generate mockgen -source=storage.go -destination=../../../adapter/storage/gorm/mock/mock_storage.go -package=mock
type IDB interface {
	InitRepository() *repository.Manager
	Close() error
}
