package deliveryman

import (
	"context"
	"fmt"

	"encore.app/delivery/database/entities"
	"encore.app/delivery/database/utils"
)

// UpdateDTO contains data to update a deliveryman
type UpdateDTO struct {
	Name   string                     `json:"owner,omitempty"`
	Status entities.DeliveryManStatus `json:"status,omitempty"`
}

// UpdateRO contains a message of confirmation
type UpdateRO struct {
	Message string `json:"message"`
}

// Update deliveryman selected by his unique identifier with given properties
//encore:api auth method=PUT path=/deliveryman/:id
func Update(_ context.Context, id string, data *UpdateDTO) (*UpdateRO, error) {
	expr, args := utils.BuildUpdateArgs(id, data.getValues())
	_, err := db.Exec(fmt.Sprintf("UPDATE deliveryman SET %v WHERE id=$1", expr), args...)
	if err != nil {
		return nil, err
	}
	return &UpdateRO{"deliveryman successfully updated"}, nil
}

// getValues Create a map of values from UpdateDTO than can be easily
// manipulate
func (d *UpdateDTO) getValues() *utils.Values {
	ret := utils.Values{
		"name":   d.Name,
		"status": d.Status.String(),
	}

	for i, e := range ret {
		if e == "" {
			delete(ret, i)
		}
	}
	return &ret
}
