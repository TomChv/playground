package utils

import (
	"context"

	"encore.app/delivery/database/entities"
	"encore.app/delivery/deliveryman"
	"encore.app/delivery/packet"
)

// CheckPacketStatus compare the targeted packet status with a given status
// If status match, it returns true, else false
func CheckPacketStatus(ctx context.Context, id string, status entities.PacketStatus) (bool, error) {
	pkt, err := packet.Get(ctx, id)
	if err != nil {
		return false, err
	}

	if pkt.Status != status {
		return false, nil
	}
	return true, nil
}

// CheckDeliverymanStatus compare the targeted deliveryman status with a given status
// If status match, it returns true, else false
func CheckDeliverymanStatus(ctx context.Context, id string, status entities.DeliveryManStatus) (bool, error) {
	d, err := deliveryman.Get(ctx, id)
	if err != nil {
		return false, err
	}

	if d.Status != status {
		return false, nil
	}
	return true, nil
}
