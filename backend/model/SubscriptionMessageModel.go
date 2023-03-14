package model

// SubscriptionMessage is the message model for redis PubSub
type SubscriptionMessage struct {
	ClientUUID string  `json:"uuid"`
	Colors     []uint8 `json:"colors"`
}
