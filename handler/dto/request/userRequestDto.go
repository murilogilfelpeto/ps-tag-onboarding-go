package request

type UserRequestDto struct {
	FirstName string `json:"firstName" binding:"required"`
	LastName  string `json:"lastName" binding:"required"`
	Email     string `json:"email" binding:"required,email"`
	Age       int    `json:"age" binding:"required,gt=18" required:"cannot be null and must be greater than 18"`
}
