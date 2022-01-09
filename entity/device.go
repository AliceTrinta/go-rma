package entity

//Device corresponds to an IoTObject at the Resource Management Layer (RMA).
type Device struct {
	UUID               string     `json:"UUID" bson:"UUID"`
	GatewayUUID        string     `json:"gateway_UUID" bson:"gatewayUUID"`
	Name               string     `json:"name" bson:"name"`
	Description        string     `json:"description" bson:"description"`
	CycleDelayInMillis int        `json:"cycleDelayInMillis" bson:"cycleDelayInMillis"`
	ResourceList       []Resource `json:"resourceList" bson:"resourceList"`
	Status             bool       `json:"status" bson:"status"`
}
