package resources

import "github.com/oneaushaf/go-broiler/models"

type RanchResource struct {
	ID      uint
	Code    string
	Capacity  uint
}

func RanchDefaultResource(Ranch models.Ranch) RanchResource {
	var result RanchResource

	result.ID = Ranch.ID
	result.Code = Ranch.Code
	result.Capacity = Ranch.Capacity
	
	return result
}