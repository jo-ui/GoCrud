package service

type CreatePersonRequest struct {
	Name    string   `json:"name" binding:"required,min=2,max=100" example:"John Doe"`
	Age     int      `json:"age" binding:"required,gte=0,lte=120" example:"25"`
	Hobbies []string `json:"hobbies" binding:"required"`
}
