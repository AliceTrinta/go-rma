package entity

//An Environment is a representation of the environment were the physical Device is.
type Environment struct {
	UUID        string   `json:"UUID" bson:"UUID"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	Devices     []Device `json:"devices" bson:"devices"`
}
