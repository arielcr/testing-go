package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	total := Add(1, 3)

	assert.NotNil(t, total)

	assert.Equal(t, 4, total)

}

func TestSubtract(t *testing.T) {

	total := Subtract(1, 3)

	assert.NotNil(t, total)

	assert.Equal(t, -2, total)

}