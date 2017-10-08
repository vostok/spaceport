package integration

import (
	"bytes"
	"net/http"
	"time"
)

func SendBody(payloadURL, apiKeyField, apiKey, rountingKey string) *http.Response {
	body := createCorrectBody(rountingKey, 10).GetBytes()
	client := &http.Client{}
	req, _ := http.NewRequest("POST", payloadURL, bytes.NewReader(body))
	req.Header.Set(apiKeyField, apiKey)
	resp, _ := client.Do(req)
	return resp
}

func createCorrectBody(routingKey string, recordSize int32) AirlockMessage {
	arr := ByteArray{}
	arr.Size = recordSize
	arr.Bytes = make([]byte, recordSize)

	record := EventRecord{}
	record.Timestamp = time.Now().Unix()
	record.Data = arr

	eventGroup := EventGroup{}
	eventGroup.KeySize = int32(len(routingKey))
	eventGroup.RoutingKey = routingKey
	eventGroup.RecordsSize = 1
	eventGroup.Records = []EventRecord{record}

	groups := EventGroupList{}
	groups.Size = 1
	groups.List = []EventGroup{eventGroup}

	message := AirlockMessage{}
	message.Version = 1
	message.EventGroups = groups
	return message
}
