package repository

import (
	"go_crud/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PersonRepository struct {
	db *gorm.DB
}

func NewPersonRepository(db *gorm.DB) *PersonRepository {
	return &PersonRepository{db: db}
}

func (repo *PersonRepository) Create(person *domain.Person) error {
	return repo.db.Create(person).Error
}

func (repo *PersonRepository) GetByID(id uuid.UUID) (*domain.Person, error) {
	var person domain.Person
	err := repo.db.First(&person, "id = ?", id).Error
	return &person, err
}

func (repo *PersonRepository) Update(person *domain.Person) error {
	return repo.db.Save(person).Error
}

func (repo *PersonRepository) Delete(id uuid.UUID) error {
	return repo.db.Delete(&domain.Person{}, "id = ?", id).Error
}

func (repo *PersonRepository) GetAll(page, limit int, sortedBy, sortedOrder string) ([]domain.Person, int64, error) {
	var persons []domain.Person
	var totalRecords int64

	// pagination and sorting
	query := repo.db.Model(&domain.Person{}).Count(&totalRecords)
	if sortedBy != "" && sortedOrder != "" {
		query = query.Order(sortedBy + " " + sortedOrder)
	}
	err := query.Offset((page - 1) * limit).Limit(limit).Find(&persons).Error

	return persons, totalRecords, err
}
