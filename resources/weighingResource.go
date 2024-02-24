package resources

import "github.com/oneaushaf/go-broiler/models"

type WeighingResource struct {
	ID            uint
	Deceased      uint
	Image         string
	Age           uint
	AverageWeight float64
}

func WeighingDefaultResource(Ranch models.Weighing) WeighingResource {
	var result WeighingResource

	result.ID = Ranch.ID
	result.Deceased = Ranch.Deceased
	result.Image = Ranch.Image
	result.Age = Ranch.Age
	result.AverageWeight = Ranch.AverageWeight

	return result
}
