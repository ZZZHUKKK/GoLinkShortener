package link

import "demo/linker/pkg/db"

type LinkRepository struct {
	Database *db.Db
}

func NewLinkRepo(databse *db.Db) *LinkRepository {
	return &LinkRepository{
		Database: databse,
	}
}

func (repo *LinkRepository) Create(link *LinkModel) (*LinkModel, error) {
	result := repo.Database.DB.Table("link_models").Create(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) Update(link *LinkModel) (*LinkModel, error) {
	result := repo.Database.DB.Table("link_models").Updates(link)
	if result.Error != nil {
		return nil, result.Error
	}
	return link, nil
}

func (repo *LinkRepository) GetByHash(hash string) (*LinkModel, error) {
	var link LinkModel
	result := repo.Database.DB.First(&link, "hash = ?", hash)
	if result.Error != nil {
		return nil, result.Error
	}
	return &link, nil
}

func (repo *LinkRepository) CheckByID(id uint) error {
	var link LinkModel
	result := repo.Database.DB.First(&link, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *LinkRepository) Delete(id uint) error {
	result := repo.Database.DB.Table("link_models").Delete(&LinkModel{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
