package pkg

import (
	"encoding/json"

	pb "github.com/dueruen/WasteChain/service/shipment/gen/proto"
)

type DataEvent struct {
	Event      pb.ShipmentEvent
	OwnerID    string
	ReceiverID string
	TimeStamp  string
	Location   string
}

func mapData(historyItem *pb.HistoryItem) *DataEvent {
	newDataEvent := &DataEvent{
		Event:      historyItem.Event,
		OwnerID:    historyItem.OwnerID,
		ReceiverID: historyItem.ReceiverID,
		TimeStamp:  historyItem.TimeStamp,
		Location:   historyItem.Location,
	}

	return newDataEvent
}

func ToByte(event *pb.HistoryItem) []byte {
	data, _ := json.Marshal(mapData(event))
	return data
}
