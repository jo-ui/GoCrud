package usecase

import (
	"go_crud/internal/domain"
	"go_crud/internal/repository"

	"github.com/google/uuid"
)

type PersonUsecase struct {
	repo *repository.PersonRepository
}

func NewPersonUsecase(repo *repository.PersonRepository) *PersonUsecase {
	return &PersonUsecase{repo: repo}
}

func (uc *PersonUsecase) CreatePerson(person *domain.Person) error {
	person.ID = uuid.New()
	return uc.repo.Create(person)
}

func (uc *PersonUsecase) GetPersonByID(id uuid.UUID) (*domain.Person, error) {
	return uc.repo.GetByID(id)
}

func (u *PersonUsecase) UpdatePerson(id uuid.UUID, updatedData *domain.Person) (*domain.Person, error) {
	person, err := u.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	person.Name = updatedData.Name
	person.Age = updatedData.Age
	person.Hobbies = updatedData.Hobbies

	if err := u.repo.Update(person); err != nil {
		return nil, err
	}
	return person, nil
}

func (uc *PersonUsecase) DeletePerson(id uuid.UUID) error {
	return uc.repo.Delete(id)
}

func (uc *PersonUsecase) GetAllPersons(page, limit int, sortedBy, sortedOrder string) ([]domain.Person, int64, error) {
	return uc.repo.GetAll(page, limit, sortedBy, sortedOrder)
}
