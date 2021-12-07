package deliveryman

import "context"

// CreateDTO Contains data to create a new deliveryman
type CreateDTO struct {
	Name string `json:"name"`
}

// AddRO contains a message of confirmation
type AddRO struct {
	Message string `json:"message"`
}

// Add a deliveryman
//encore:api auth method=POST path=/deliveryman
func Add(_ context.Context, data *CreateDTO) (*AddRO, error) {
	_, err := db.Exec("INSERT INTO deliveryman (name) VALUES($1)", data.Name)
	if err != nil {
		return nil, err
	}
	return &AddRO{"deliveryman successfully created"}, nil
}
