package helloworld

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestHelloName(t *testing.T) {
	name := "Tyler"
	msg, err := SayHello(name)
	require.Nil(t, err)
	require.Equal(t, msg, "Hello Tyler!", "Says hello to Tyler.")
}
