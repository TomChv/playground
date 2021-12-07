package packet

import (
	"context"
	"encoding/json"

	"encore.app/delivery/database"
	"encore.app/delivery/database/entities"
)

// GetAllRO contains the number of packets returned and packets
type GetAllRO struct {
	Number   int               `json:"number"`
	Packages []entities.Packet `json:"packages"`
}

// GetAll return all packets from the API
//encore:api public method=GET path=/packet
func GetAll(_ context.Context) (*GetAllRO, error) {
	result := []database.Result{}
	// On step 1, it could be SELECT * FROM packet
	err := db.Select(&result, "SELECT ROW_TO_JSON(results) AS result"+
		" FROM ( SELECT * FROM packets_delivery ) AS results")
	if err != nil {
		return nil, err
	}

	packets := []entities.Packet{}
	for _, e := range result {
		packet := entities.Packet{}
		err = json.Unmarshal([]byte(e.Result), &packet)
		if err != nil {
			return nil, err
		}
		packets = append(packets, packet)
	}

	return &GetAllRO{
		len(packets),
		packets,
	}, nil
}

// Get return a package by his unique identifier
//encore:api public method=GET path=/packet/:id
func Get(_ context.Context, id string) (*entities.Packet, error) {
	result := database.Result{}
	// On step 1, it could be SELECT * FROM packet WHERE id=$1
	err := db.Get(&result, "SELECT ROW_TO_JSON(result) AS result FROM ( SELECT * FROM packets_delivery WHERE id=$1) as result", id)
	if err != nil {
		return nil, err
	}

	packet := entities.Packet{}
	err = json.Unmarshal([]byte(result.Result), &packet)
	if err != nil {
		return nil, err
	}

	return &packet, nil
}
