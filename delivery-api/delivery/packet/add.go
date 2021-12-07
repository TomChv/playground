package packet

import "context"

// CreateDTO Contains data to create a new packet
type CreateDTO struct {
	Owner       string `json:"owner"`
	Receiver    string `json:"receiver"`
	Destination string `json:"destination"`
}

// AddRO contains a message of confirmation
type AddRO struct {
	Message string `json:"message"`
}

// Add a packet
//encore:api public method=POST path=/packet
func Add(_ context.Context, data *CreateDTO) (*AddRO, error) {
	_, err := db.Exec(`INSERT INTO packet (owner, receiver, destination) VALUES ($1, $2, $3)`, data.Owner, data.Receiver, data.Destination)
	if err != nil {
		return nil, err
	}

	return &AddRO{"packet successfully created"}, nil
}
