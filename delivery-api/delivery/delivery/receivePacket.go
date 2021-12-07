package delivery

import (
	"context"

	"encore.app/delivery/database/entities"
	"encore.app/delivery/delivery/utils"
	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
)

type ReceivePacketDTO struct {
	PacketID      string
	DeliveryManID string
}

type ReceivePacketRO struct {
	Message string `json:"message"`
}

// Check that packet & deliveryman are available
func checkReceiveStatus(params *ReceivePacketDTO) error {
	// Create a fake auth context to bypass auth from Getter
	ctx := auth.WithContext(context.Background(), "", nil)

	exist, err := utils.CheckPacketStatus(ctx, params.PacketID, entities.TRAVELING)
	if err != nil {
		return nil
	}

	if !exist {
		return &errs.Error{
			Code:    errs.Unavailable,
			Message: "packet is not traveling",
		}
	}

	return nil
}

// ReceivePacket receive a packet previously sent
//encore:api public method=PUT path=/receive
func ReceivePacket(_ context.Context, params *ReceivePacketDTO) (*ReceivePacketRO, error) {
	err := checkReceiveStatus(params)
	if err != nil {
		return nil, err
	}

	// Update packet
	_, err = db.Exec("UPDATE packet SET status = 'Received', deliveryman_id = $1 WHERE id = $2", params.DeliveryManID, params.PacketID)
	if err != nil {
		return nil, err
	}

	// Update deliveryman
	_, err = db.Exec("UPDATE deliveryman SET status = 'Available' WHERE id = $1", params.DeliveryManID)
	if err != nil {
		return nil, err
	}

	return &ReceivePacketRO{"packet successfully received"}, nil
}
