package delivery

import (
	"context"

	"encore.app/delivery/database/entities"
	"encore.app/delivery/delivery/utils"
	"encore.dev/beta/auth"
	"encore.dev/beta/errs"
)

type SendPacketDTO struct {
	PacketID      string
	DeliveryManID string
}

type SendPacketRO struct {
	Message string `json:"message"`
}

// Check that packet & deliveryman are available
func checkSendStatus(params *SendPacketDTO) error {
	// Create a fake auth context to bypass auth from Getter
	ctx := auth.WithContext(context.Background(), "", nil)

	exist, err := utils.CheckPacketStatus(ctx, params.PacketID, entities.SEND)
	if err != nil {
		return err
	}

	if !exist {
		return &errs.Error{
			Code:    errs.Unavailable,
			Message: "packet has been already sent",
		}
	}

	exist, err = utils.CheckDeliverymanStatus(ctx, params.DeliveryManID, entities.AVAILABLE)
	if err != nil {
		return err
	}

	if !exist {
		return &errs.Error{
			Code:    errs.Unavailable,
			Message: "deliveryman is already on a mission",
		}
	}

	return nil
}

// SendPacket send a packet selected by his unique and assign it to a deliveryman
//encore:api public method=PUT path=/send
func SendPacket(_ context.Context, params *SendPacketDTO) (*SendPacketRO, error) {
	err := checkSendStatus(params)
	if err != nil {
		return nil, err
	}

	// Update packet
	_, err = db.Exec("UPDATE packet SET status = 'Traveling', deliveryman_id = $1 WHERE id = $2", params.DeliveryManID, params.PacketID)
	if err != nil {
		return nil, err
	}

	// Update deliveryman
	_, err = db.Exec("UPDATE deliveryman SET status = 'Working' WHERE id = $1", params.DeliveryManID)
	if err != nil {
		return nil, err
	}

	return &SendPacketRO{"packet successfully sent"}, nil
}
