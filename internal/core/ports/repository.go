package ports

import (
	postgres "qropen-backend/internal/adapters/repositories/postgres"
	"qropen-backend/pkg/database"
)

type Repositories struct {
	UserRepo postgres.UserRepository
}

func NewRepositories(db *database.PostgresDB) Repositories {
	return Repositories{
		UserRepo: postgres.NewUserRepository(db),
	}
}
