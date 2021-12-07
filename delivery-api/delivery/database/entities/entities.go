package entities

type Packet struct {
	ID          string       `json:"id"`
	Owner       string       `json:"owner"`
	Receiver    string       `json:"receiver"`
	Destination string       `json:"destination"`
	Status      PacketStatus `json:"status"`
	DeliveryMan *DeliveryMan `json:"deliveryMan,omitempty"`
}

type DeliveryMan struct {
	ID      string            `json:"id"`
	Name    string            `json:"name"`
	Status  DeliveryManStatus `json:"status"`
	Packets []*Packet         `json:"packets,omitempty"`
}
