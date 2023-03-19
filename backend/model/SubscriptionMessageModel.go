package model

import (
	"bytes"
	"encoding/binary"
)

// SubscriptionMessage contains the message model between browser and the websockets server
type SubscriptionMessage struct {
	Colors uint8  `bson:"color"`
	PosX   uint16 `bson:"PosX"`
	PosY   uint16 `bson:"PosY"`
}

// InternalSubscriptionMessage is the message model for redis PubSub
type InternalSubscriptionMessage struct {
	ClientUUID string              `bson:"uuid"`
	Message    SubscriptionMessage `bson:"message"`
}

// EncodeSubscriptionMessage encodes the data into a simple binary format
func EncodeSubscriptionMessage(buf *bytes.Buffer, data SubscriptionMessage) error {
	if err := binary.Write(buf, binary.BigEndian, data.PosX); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, data.PosY); err != nil {
		return err
	}
	if err := buf.WriteByte(data.Colors); err != nil {
		return err
	}
	return nil
}

// DecodeSubscriptionMessage decodes the data from the browser into a usable struct
func DecodeSubscriptionMessage(buf []uint8) (SubscriptionMessage, error) {
	value := SubscriptionMessage{}
	readerPosX := bytes.NewReader(buf[0:2])
	readerPosY := bytes.NewReader(buf[2:5])

	if err := binary.Read(readerPosX, binary.BigEndian, &value.PosX); err != nil {
		return value, err
	}
	if err := binary.Read(readerPosY, binary.BigEndian, &value.PosY); err != nil {
		return value, err
	}
	value.Colors = buf[len(buf)-1]
	return value, nil
}
