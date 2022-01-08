package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//Testing the DelegateAction func.
func TestAction_DelegateAction(t *testing.T) {
	var fake FakeAction
	con := FakeConnection()
	err := fake.DelegateAction(con)
	assert.Nil(t, err)
}

// //Testing the sendAction func.
// func TestSendAction(t *testing.T) {
// 	var buf bytes.Buffer
// 	log.SetOutput(&buf)
// 	defer func() {
// 		log.SetOutput(os.Stderr)
// 	}()
// 	sendAction("135432567", "35143657", "235254222111")
// 	t.Log(buf.String())
// 	var x = buf.String()
// 	assert.Contains(t, x, "135432567")
// 	assert.Contains(t, x, "35143657")
// 	assert.Contains(t, x, "235254222111")
// }
