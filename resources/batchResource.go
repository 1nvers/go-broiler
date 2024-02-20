package resources

import (
	"github.com/oneaushaf/go-broiler/models"
)

type BatchResource struct {
	ID      	     uint
	StartingDate     string
	FinishedDate     string
	InitialQty       uint
	CurrentQty       uint
	Deceased         uint
	Finished 		 bool
}

func BatchDefaultResource(batch models.Batch) BatchResource {
	var result BatchResource

	result.ID = batch.ID
	result.StartingDate = batch.CreatedAt.Format("2006-01-02")
	result.InitialQty = batch.InitialQty
	result.CurrentQty = batch.CurrentQty
	result.Deceased = batch.InitialQty - batch.CurrentQty
	result.Finished = batch.Finished
	if (result.Finished) {
		result.FinishedDate = batch.UpdatedAt.Format("2006-01-02")
	}
 	
	return result
}