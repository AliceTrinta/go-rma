package entity

import (
	"time"
)

//A Data is an information sent by a resource.
type Data struct {
	ID           int       `json:"_id" bson:"_id"`
	Instant      time.Time `json:"instant" bson:"instant"`
	DeviceUUID   string    `json:"UUID" bson:"UUID"`
	ResourceName string    `json:"resourceName" bson:"resourceName"`
	Value        string    `json:"value" bson:"value"`
}

// //The IoTData Gather all the functions that a data can implement.
// type IoTData interface {
// 	SaveData(con MongoDB) (err error)
// }

// //DataInstance is an instance of IoTData
// var DataInstance IoTData

// //SaveData function will register the data received through an IoT network in the database.
// func (d Data) SaveData(con MongoDB) (err error) {
// 	//Creating a data in the database
// 	err = con.CreateData(d)
// 	if err != nil {
// 		log.Println("It was not possible to save data ")
// 		return
// 	}
// 	log.Println("Message registered Successfully!")
// 	return

// }
