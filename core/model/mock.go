package model

//This file contains mock objects created for tests.

//Creating a mock for connection with database.
func FakeConnection () (con MongoDB){
	MongoConnect()
	return con
}

//Creating a mock of an environment.
type FakeEnvironment struct {}

//Creating a mock of the StartEnvironment func.
func (e FakeEnvironment)StartEnvironment (con MongoDB) (err error) {return}

//Creating a mock of a Device.
type FakeDevice struct {}

//Creating a mock of the StartDevice func.
func (o FakeDevice)StartDevice (con MongoDB) (err error){return}

//Creating a mock of a Data.
type FakeData struct {}

//Creating a mock of the SaveData func.
func (d FakeData)SaveData (con MongoDB) (err error) {return}

//Creating a mock of an action.
type FakeAction struct {}

//Creating a mock of the DelegateAction func.
func (a FakeAction)DelegateAction (con MongoDB) (err error) {return}
