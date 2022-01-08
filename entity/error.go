package entity

//InvalidJSONInput indicates that some message received through the IoT network that does not match with a Device, a Data or an Action
type InvalidJSONInput struct{}

func (I *InvalidJSONInput) Error() string {
	return "This input isn't a Device, Data or Action!"
}
