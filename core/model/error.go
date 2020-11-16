package model

type InvalidJSONInput struct{}

func (I *InvalidJSONInput) Error() string {
	return "This input isn't a Device, Data or Action!"
}
