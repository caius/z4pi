package app

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAppNew(t *testing.T) {
	newApp, err := NewApp("/dev/null", 8080)
	assert.Nil(t, err)

	assert.Equal(t, newApp.KbusPath, "/dev/null")
	assert.Equal(t, newApp.HttpPort, 8080)
}
