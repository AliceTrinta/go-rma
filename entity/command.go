package entity

//A Command is an something that a resource can perform.
type Command struct {
	ID          string `json:"_id" json:"_id"`
	Description string `json:"description" bson:"description"`
}
