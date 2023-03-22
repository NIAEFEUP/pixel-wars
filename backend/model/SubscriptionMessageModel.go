package model

import (
	"bytes"
	"encoding/binary"
)

// PixelColorUpdateMessage contains the message model between browser and the websockets server
type PixelColorUpdateMessage struct {
	Color uint8  `bson:"color"`
	PosX  uint16 `bson:"PosX"`
	PosY  uint16 `bson:"PosY"`
}

// PixelColorUpdatePubSubMessage is the message model for redis PubSub
type PixelColorUpdatePubSubMessage struct {
	ClientUUID string                  `bson:"uuid"`
	Message    PixelColorUpdateMessage `bson:"message"`
}

// EncodePixelColorUpdateMessage encodes the data into a simple binary format
func EncodePixelColorUpdateMessage(buf *bytes.Buffer, data PixelColorUpdateMessage) error {
	if err := binary.Write(buf, binary.BigEndian, data.PosX); err != nil {
		return err
	}
	if err := binary.Write(buf, binary.BigEndian, data.PosY); err != nil {
		return err
	}
	if err := buf.WriteByte(data.Color); err != nil {
		return err
	}
	return nil
}

// DecodePixelColorUpdateMessage decodes the data from the browser into a usable struct
func DecodePixelColorUpdateMessage(buf []uint8) (PixelColorUpdateMessage, error) {
	value := PixelColorUpdateMessage{}
	readerPosX := bytes.NewReader(buf[0:2])
	readerPosY := bytes.NewReader(buf[2:5])

	if err := binary.Read(readerPosX, binary.BigEndian, &value.PosX); err != nil {
		return value, err
	}
	if err := binary.Read(readerPosY, binary.BigEndian, &value.PosY); err != nil {
		return value, err
	}
	value.Color = buf[len(buf)-1]
	return value, nil
}
