package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndex(t *testing.T) {
	vs := []string{"first name", "last name", "company"}
	r := Index(vs, "company")

	assert.Equal(t, 2, r)
}
