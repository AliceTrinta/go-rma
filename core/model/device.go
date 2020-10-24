package model

import (
	"encoding/json"
	"time"
)

type Device struct {
	ID string `json:"_id" bson:"_id"`
	UUID string `json:"UUID" bson:"UUID"`
	GatewayUUID string `json:"gateway_UUID" bson:"gatewayUUID"`
	Name string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	CycleDelayInMillis json.Number `json:"cycleDelayInMillis" bson:"cycleDelayInMillis"`
	ResourceList []Resource `json:"resourceList" bson:"resourceList"`
	LastUpdate time.Time `json:"lastUpdate" bson:"lastUpdate"`
}
