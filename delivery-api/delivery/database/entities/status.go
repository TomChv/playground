package entities

import "fmt"

// IStatus define method to manipulate Status
type IStatus interface {
	fmt.Stringer
}

// PacketStatus Enumerate kind of delivery status
type PacketStatus string

// PacketStatus values
const (
	SEND      PacketStatus = "Send"
	TRAVELING PacketStatus = "Traveling"
	RECEIVED  PacketStatus = "Received"
)

// Get the packet status as string
func (s PacketStatus) String() string {
	return string(s)
}

// DeliveryManStatus Enumerate kind of delivery status
type DeliveryManStatus string

// DeliveryManStatus values
const (
	AVAILABLE DeliveryManStatus = "Available"
	WORKING   DeliveryManStatus = "Working"
)

// Get the delivery status as string
func (s DeliveryManStatus) String() string {
	return string(s)
}
