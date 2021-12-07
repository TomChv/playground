package packet

import "context"

// DeleteRO contains a message of confirmation
type DeleteRO struct {
	Message string `json:"message"`
}

// Delete a packet by his unique identifier
//encore:api public method=DELETE path=/packet/:id
func Delete(_ context.Context, id string) (*DeleteRO, error) {
	_, err := db.Exec("DELETE FROM packet WHERE id=$1", id)
	if err != nil {
		return nil, err
	}

	return &DeleteRO{"packet successfully deleted"}, nil
}
