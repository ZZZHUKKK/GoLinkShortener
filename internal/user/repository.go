package user

import "demo/linker/pkg/db"

type UserRepository struct {
	Database *db.Db
}

func NewUserRepo(databse *db.Db) *UserRepository {
	return &UserRepository{
		Database: databse,
	}
}

func (repo *UserRepository) Create(user *UserModel) (*UserModel, error) {
	result := repo.Database.DB.Table("user_models").Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) error {
	var user UserModel
	result := repo.Database.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
