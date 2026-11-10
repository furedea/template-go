package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreeting(t *testing.T) {
	t.Parallel()

	assert.Equal(t, "Hello, World!", greeting())
}
