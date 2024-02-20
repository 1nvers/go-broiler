package resources

import "github.com/oneaushaf/go-broiler/models"

type FarmResource struct {
	ID      uint
	Code    string
	Adress  string
}

func FarmDefaultResource(farm models.Farm) FarmResource {
	var result FarmResource

	result.ID = farm.ID
	result.Code = farm.Code
	result.Adress = farm.Adress
	
	return result
}