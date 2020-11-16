package services

func (m *MongoDB) GetDeviceByID(id string) (Device, error) {
	var device Device
	err := Db.C("device").FindId(id).One(&device)
	return device, err
}

func (m *MongoDB) CreateDevice(device Device) error {
	err := Db.C("device").Insert(&device)
	return err
}


func (m *MongoDB) CreateData(data Data) error {
	err := Db.C("data").Insert(data)
	return err
}

func (m *MongoDB) CreateAction(action Action) error {
	err := Db.C("action").Insert(action)
	return err
}
