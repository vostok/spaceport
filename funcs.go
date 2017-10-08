package integration

import (
	"bytes"
	"net/http"
	"time"

	airlock "github.com/vostok/airlock-client-go"
)

// SendBody sends message to airlock
func SendBody(payloadURL, apiKeyField, apiKey, rountingKey string) *http.Response {
	body := createCorrectBody(rountingKey, 10).GetBytes()
	client := &http.Client{}
	req, _ := http.NewRequest("POST", payloadURL, bytes.NewReader(body))
	req.Header.Set(apiKeyField, apiKey)
	resp, _ := client.Do(req)
	return resp
}

func createCorrectBody(routingKey string, recordSize int32) airlock.Message {
	arr := airlock.ByteArray{}
	arr.Size = recordSize
	arr.Bytes = make([]byte, recordSize)

	record := airlock.EventRecord{}
	record.Timestamp = time.Now().Unix()
	record.Data = arr

	eventGroup := airlock.EventGroup{}
	eventGroup.KeySize = int32(len(routingKey))
	eventGroup.RoutingKey = routingKey
	eventGroup.RecordsSize = 1
	eventGroup.Records = []airlock.EventRecord{record}

	groups := airlock.EventGroupList{}
	groups.Size = 1
	groups.List = []airlock.EventGroup{eventGroup}

	message := airlock.Message{}
	message.Version = 1
	message.EventGroups = groups
	return message
}
