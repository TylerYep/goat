package helloworld

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Tyler"
	msg, err := SayHello(name)
	assert.Nil(t, err)
	assert.Equal(t, msg, "Hello Tyler!", "Says hello to Tyler.")
}
