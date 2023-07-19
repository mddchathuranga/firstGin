package models

import (
	"github.com/dasun/first_api/src/dtos"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Address  string `json:"address"`
	Mobile   string `json:"mobile"`
	Age      int    `json:"age"`
	Email    string `json:"email"`
}

func MapToGetDTO(user User) dtos.GetDTO {
	return dtos.GetDTO{
		ID:       user.ID,
		Username: user.Username,
		Address:  user.Address,
		Mobile:   user.Mobile,
		Age:      user.Age,
		Email:    user.Email,
	}
}

func MapToGetDTOs(users []User) []dtos.GetDTO {
	getDTOs := make([]dtos.GetDTO, len(users))
	for i, user := range users {
		getDTOs[i] = MapToGetDTO(user)
	}
	return getDTOs
}

func MapToUser(createDTO dtos.CreateDTO) User {
	return User{
		Username: createDTO.Username,
		Address:  createDTO.Address,
		Mobile:   createDTO.Mobile,
		Age:      createDTO.Age,
		Email:    createDTO.Email,
	}
}

func MapToUserFromUpdateDTO(updateDTO dtos.UpdateDTO) User {
	return User{
		Username: updateDTO.Username,
		Address:  updateDTO.Address,
		Mobile:   updateDTO.Mobile,
		Age:      updateDTO.Age,
		Email:    updateDTO.Email,
	}
}
