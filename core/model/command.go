package model

type Command struct {
	ID string `json:"_id" json:"_id"`
	Description string `json:"description" bson:"description"`
}
