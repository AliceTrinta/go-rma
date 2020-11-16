package model

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"
)

func TestAction_DelegateAction(t *testing.T) {
	var fake FakeAction
	con := FakeConnection()
	err := fake.DelegateAction(con)
	if err != nil{
		log.Fatal(err)
	}
}

func TestSendAction(t *testing.T) {
	var buf bytes.Buffer
	log.SetOutput(&buf)
	defer func() {
		log.SetOutput(os.Stderr)
	}()
	sendAction("135432567", "35143657", "235254222111")
	t.Log(buf.String())
	var x = buf.String()
	if !strings.Contains(x, "135432567"){
		t.Fatal("x value unexpected, != 135432567", x)
	}
	if !strings.Contains(x, "35143657"){
		t.Fatal("x value unexpected, != 35143657", x)
	}
	if !strings.Contains(x, "235254222111"){
		t.Fatal("x value unexpected, != 235254222111", x)
	}
}
