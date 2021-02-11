package model

import "log"

//An Environment is a representation of the environment were the physical Device is.
type Environment struct {
	UUID        string   `json:"UUID" bson:"UUID"`
	Name        string   `json:"name" bson:"name"`
	Description string   `json:"description" bson:"description"`
	Devices     []Device `json:"devices" bson:"devices"`
}

//IoTEnvironment gather all the functions that an environment can implement.
type IoTEnvironment interface {
	StartEnvironment(con MongoDB) (err error)
}

//EnvironmentInstance is an instance of an IoTEnvironment
var EnvironmentInstance IoTEnvironment

//StartEnvironment function will register a new environment in the database.
func (e Environment) StartEnvironment(con MongoDB) (err error) {
	//Verifying if the Environment already exist.
	_, err = con.GetEnvironmentByUUID(e.UUID)
	if err != nil {
		log.Println("Registering new Environment")
		//If does not exist, register a new environment in the database.
		err = con.CreateEnvironment(e)
		if err != nil {
			log.Println("Environment was not registered ")
			return
		}
		return
	}
	log.Println("This environment is already registered")
	return
}
