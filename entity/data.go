package entity

//A Data is an information sent by a resource.
type Data struct {
	Instant      string `json:"instant" bson:"instant"`
	DeviceUUID   string `json:"deviceUUID" bson:"deviceUUID"`
	ResourceName string `json:"resourceName" bson:"resourceName"`
	Value        string `json:"value" bson:"value"`
}
