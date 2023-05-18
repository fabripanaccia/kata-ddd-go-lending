package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIdempotency_Get(t *testing.T) {
	a := assert.New(t)
	t.Run("Success", func(t *testing.T) {
		dom := ddd{}
		resp := dom.ValePorUnMetodo()
		a.Equal("Vale por una respuesta", resp)
	})
}
