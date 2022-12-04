package repository

import (
	"blogs/infrastructure"
	"blogs/models"
)

type PostRepository struct {
	db infrastructure.Database
}

func NewPostRepository(db infrastructure.Database) PostRepository {
	return PostRepository{
		db: db,
	}
}

func (p PostRepository) Save(post models.Post) error {
	return p.db.DB.Create(&post).Error
}

func (p PostRepository) FindAll(post models.Post, keyword string) (*[]models.Post, int64, error) {
	var posts []models.Post
	var totalRows int64 = 0

	queryBuilder := p.db.DB.Order("created_at desc").Model(&models.Post{})

	if keyword != "" {
		queryKeyword := "'%" + keyword + "%'"
		queryBuilder = queryBuilder.Where(p.db.DB.Where("post.title LIKE " + queryKeyword))
	}
	// Tìm đối tượng với ID theo Post và trả về mảng và đếm số phần tử
	err := queryBuilder.Where(post).Find(&posts).Count(&totalRows).Error

	return &posts, totalRows, err
}

func (p PostRepository) Update(post models.Post) error {
	return p.db.DB.Save(&post).Error
}

func (p PostRepository) Find(post models.Post) (models.Post, error) {
	var item models.Post
	err := p.db.DB.Debug().Model(&models.Post{}).Where(&post).Take(&item).Error
	return item, err
}

func (p PostRepository) Delete(post models.Post) error {
	return p.db.DB.Delete(&post).Error
}
