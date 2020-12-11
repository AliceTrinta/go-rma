package model

//A Resource represents a sensor or an actuator of a Device.
type Resource struct {
	ID string             `json:"_id" bson:"_id"`
	Name string           `json:"name" bson:"name"`
	Description string    `json:"description" bson:"description"`
	Port string           `json:"port" bson:"port"`
	DataUnit string       `json:"dataUnit" bson:"dataUnit"`
	WaitTimeInMillis int  `json:"waitTimeInMillis" bson:"waitTimeInMillis"`
	CommandList []Command `json:"commandList" bson:"commandList"`
}
