package deliveryman

import (
	"context"
	"encoding/json"

	"encore.app/delivery/database"
	"encore.app/delivery/database/entities"
)

// GetAllRO contains the number of deliverymans returned and deliverymans
type GetAllRO struct {
	Number int                    `json:"number"`
	GetAll []entities.DeliveryMan `json:"deliverymans"`
}

// GetAll return all deliverymans from the API
//encore:api auth method=GET path=/deliveryman
func GetAll(_ context.Context) (*GetAllRO, error) {
	result := []database.Result{}
	// On step 1, it could be SELECT * FROM deliveryman
	err := db.Select(&result, "SELECT ROW_TO_JSON(results) AS result "+
		"FROM ( SELECT *  FROM delivery_packets ) AS results")
	if err != nil {
		return nil, err
	}

	deliverymans := []entities.DeliveryMan{}
	for _, e := range result {
		deliveryman := entities.DeliveryMan{}
		err = json.Unmarshal([]byte(e.Result), &deliveryman)
		if err != nil {
			return nil, err
		}
		deliverymans = append(deliverymans, deliveryman)
	}

	return &GetAllRO{
		len(deliverymans),
		deliverymans,
	}, nil
}

// Get return a deliveryman selected by his unique identifier
//encore:api auth method=GET path=/deliveryman/:id
func Get(_ context.Context, id string) (*entities.DeliveryMan, error) {
	result := database.Result{}
	// On step 1, it could be SELECT * FROM deliveryman WHERE id=$1
	err := db.Get(&result, "SELECT ROW_TO_JSON(result) AS result FROM ( SELECT * FROM delivery_packets WHERE id=$1) as result", id)
	if err != nil {
		return nil, err
	}

	deliveryman := entities.DeliveryMan{}
	err = json.Unmarshal([]byte(result.Result), &deliveryman)
	if err != nil {
		return nil, err
	}

	return &deliveryman, nil
}
