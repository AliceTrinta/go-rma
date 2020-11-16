package model

func FakeConnection () (con MongoDB){
	MongoConnect()
	return con
}

type FakeDevice struct {}

func (o FakeDevice)StartDevice (con MongoDB) (err error){return}

type FakeData struct {}

func (d FakeData)SaveData (con MongoDB) (err error) {return}

type FakeAction struct {}

func (a FakeAction)DelegateAction (con MongoDB) (err error) {return}
