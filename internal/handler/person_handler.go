package handler

import (
	"go_crud/internal/domain"
	"go_crud/internal/service"
	"go_crud/internal/usecase"
	"net/http"
	"strconv"

	_ "go_crud/docs"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type APIMessage struct {
	Message string `json:"message"`
}

type PersonHandler struct {
	usecase *usecase.PersonUsecase
}

func NewPersonHandler(uc *usecase.PersonUsecase) *PersonHandler {
	return &PersonHandler{usecase: uc}
}

// @Summary Create a new person
// @Description Create a new person record in the database
// @Tags Person
// @Accept json
// @Produce json
// @Param person body service.CreatePersonRequest true "Create Person Request"
// @Success 201 {object} domain.Person
// @Failure 400 {object} APIMessage "error": "Invalid request format or validation error"
// @Failure 500 {object} APIMessage "error": "Internal Server Error"
// @Router /person [post]
func (h *PersonHandler) CreatePerson(c *gin.Context) {
	var req service.CreatePersonRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		validationErrors := make(map[string]string)

		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range errs {
				switch fieldErr.Field() {
				case "Name":
					validationErrors["name"] = "Name is required and must be between 2 and 100 characters."
				case "Age":
					validationErrors["age"] = "Age is required and must be between 0 and 120."
				}
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid request format"})
			return
		}

		c.JSON(http.StatusBadRequest, gin.H{"errors": validationErrors})
		return
	}

	person := domain.Person{
		Name:    req.Name,
		Age:     req.Age,
		Hobbies: req.Hobbies,
	}

	if err := h.usecase.CreatePerson(&person); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Database error: Unable to create person"})
		return
	}

	c.JSON(http.StatusCreated, person)
}

// @Summary Get all persons
// @Description Retrieve all persons with pagination and sorting options
// @Tags Person
// @Accept json
// @Produce json
// @Param page query int false "Page number" default(1)
// @Param limit query int false "Records per page" default(10)
// @Param sortedBy query string false "Field to sort by" default(name)
// @Param sortedOrder query string false "Sort order" default(asc)
// @Success 200 {object} APIMessage "data": []domain.Person, "current_page": int, "total_pages": int, "total_records": int64
// @Failure 500 {object} APIMessage "error": "Internal Server Error"
// @Router /person [get]
func (h *PersonHandler) GetAllPersons(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	sortedBy := c.DefaultQuery("sortedBy", "name")
	sortedOrder := c.DefaultQuery("sortedOrder", "asc")

	persons, totalRecords, err := h.usecase.GetAllPersons(page, limit, sortedBy, sortedOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := (totalRecords + int64(limit) - 1) / int64(limit)

	c.JSON(http.StatusOK, gin.H{
		"data":          persons,
		"current_page":  page,
		"total_pages":   totalPages,
		"total_records": totalRecords,
	})
}

// GetPersonByID godoc
// @Summary Get person by ID
// @Description Retrieve a person by their unique ID.
// @Tags Person
// @Param id path string true "Person ID"
// @Success 200 {object} domain.Person
// @Failure 404 {object} APIMessage "error": "Person not found"
// @Router /person/{id} [get]
func (h *PersonHandler) GetPersonByID(c *gin.Context) {
	id := c.Param("id")
	personUUID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Person ID"})
		return
	}
	person, err := h.usecase.GetPersonByID(personUUID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, person)
}

// UpdatePerson godoc
// @Summary Update person by ID
// @Description Update a person's details by their unique ID.
// @Tags Person
// @Param id path string true "Person ID"
// @Param person body service.CreatePersonRequest true "Update Person Request"
// @Success 200 {object} domain.Person
// @Failure 400 {object} APIMessage "error": "Invalid request"
// @Failure 404 {object} APIMessage "error": "Person not found"
// @Router /person/{id} [put]
func (h *PersonHandler) UpdatePerson(c *gin.Context) {
	id := c.Param("id")
	personUUID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Person ID"})
		return
	}
	var person domain.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	updatedPerson, err := h.usecase.UpdatePerson(personUUID, &person)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.JSON(http.StatusOK, updatedPerson)
}

// DeletePerson godoc
// @Summary Delete person by ID
// @Description Delete a person by their unique ID.
// @Tags Person
// @Param id path string true "Person ID"
// @Success 204 {string} string "No Content"
// @Failure 404 {object} APIMessage "error": "Person not found"
// @Router /person/{id} [delete]
func (h *PersonHandler) DeletePerson(c *gin.Context) {
	id := c.Param("id")
	personUUID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Person ID"})
		return
	}
	err = h.usecase.DeletePerson(personUUID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Person not found"})
		return
	}
	c.Status(http.StatusNoContent)
}
