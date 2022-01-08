package repository

//This file contains mock objects created for tests.

//FakeConnection is a mock for connection with database.
func FakeConnection() (con MongoDB) {
	MongoConnect()
	return con
}

//FakeEnvironment is a mock of an environment.
type FakeEnvironment struct{}

//StartEnvironment is a mock of the StartEnvironment func.
func (e FakeEnvironment) StartEnvironment(con MongoDB) (err error) { return }

//FakeDevice is a mock of a Device.
type FakeDevice struct{}

//StartDevice is a mock of the StartDevice func.
func (o FakeDevice) StartDevice(con MongoDB) (err error) { return }

//FakeData is a mock of a Data.
type FakeData struct{}

//SaveData is a mock of the SaveData func.
func (d FakeData) SaveData(con MongoDB) (err error) { return }

//FakeAction is a mock of an action.
type FakeAction struct{}

//DelegateAction is a mock of the DelegateAction func.
func (a FakeAction) DelegateAction(con MongoDB) (err error) { return }
