package repositories

import (
	"qropen-backend/internal/core/domain"
	"qropen-backend/pkg/database"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(postgresDB *database.PostgresDB) UserRepository {
	return UserRepository{db: postgresDB.DB()}
}

func (r *UserRepository) FindByUsername(username string) (domain.User, error) {
	var user domain.User
	if err := r.db.Where("username = ?", username).First(&user).Error; err != nil {
		return domain.User{}, err
	}
	return user, nil
}

func (r *UserRepository) CreateUser(user domain.User) error {
	return r.db.Create(&user).Error
}
