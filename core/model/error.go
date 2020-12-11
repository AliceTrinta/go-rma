package model

type InvalidJSONInput struct{}

/* A personalized error to indicate some message received through
 the IoT network that does not match with a Device, a Data or an Action */
func (I *InvalidJSONInput) Error() string {
	return "This input isn't a Device, Data or Action!"
}
