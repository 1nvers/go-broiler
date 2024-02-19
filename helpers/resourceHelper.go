package helpers

import "github.com/oneaushaf/go-broiler/models"

type UserResource struct{
	ID uint
	FirstName string
	LastName string
	Phone string
	Email string
	UserType string
}

func DefaultUserResource(user models.User)(UserResource){
	var result UserResource;

	result.ID = user.ID
	result.FirstName = user.FirstName
	result.LastName = user.LastName
	result.Phone = user.Phone
	result.Email = user.Email
	result.UserType = user.UserType

	return result
}