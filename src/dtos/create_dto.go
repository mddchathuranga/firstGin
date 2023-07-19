package dtos

type CreateDTO struct {
	Username string `json:"username"`
	Address  string `json:"address"`
	Mobile   string `json:"mobile"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
}
