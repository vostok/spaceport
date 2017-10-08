package integration

import (
	"encoding/binary"

	restruct "gopkg.in/restruct.v1"
)

// AirlockMessage represents body of complete gate message
type AirlockMessage struct {
	Version     int16
	EventGroups EventGroupList
}

// EventGroupList is group of events
type EventGroupList struct {
	Size int32
	List []EventGroup
}

// EventGroup combines events by API key
type EventGroup struct {
	KeySize     int32 // Size of routing key
	RoutingKey  string
	RecordsSize int32
	Records     []EventRecord
}

// EventRecord links byte message with Timestamp
type EventRecord struct {
	Timestamp int64
	Data      ByteArray
}

// ByteArray enriches standart array with size
type ByteArray struct {
	Size  int32
	Bytes []byte
}

type keyTest struct {
	APIKey     string
	RoutingKey string
}

// GetBytes formats structure into byte array
func (a AirlockMessage) GetBytes() []byte {
	buf, _ := restruct.Pack(binary.LittleEndian, &a)
	return buf
}
