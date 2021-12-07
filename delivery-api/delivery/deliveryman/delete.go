package deliveryman

import "context"

// DeleteRO contains a message of confirmation
type DeleteRO struct {
	Message string `json:"message"`
}

// Delete a deliveryman by his unique identifier
//encore:api auth method=DELETE path=/deliveryman/:id
func Delete(_ context.Context, id string) (*DeleteRO, error) {
	_, err := db.Exec("DELETE FROM deliveryman WHERE id=$1", id)
	if err != nil {
		return nil, err
	}
	return &DeleteRO{"deliveryman successfully deleted"}, nil
}
