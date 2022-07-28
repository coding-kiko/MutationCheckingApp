package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNotMutantResponse(t *testing.T) {
	err := NewNotMutant()
	resp := CreateResponse(err)
	assert.NotEmpty(t, resp)
	assert.Equal(t, resp.Error.Code, 403)
	assert.NotEmpty(t, resp.Error)

}

func TestCreateDuplicateDnaResponse(t *testing.T) {
	err := NewDuplicateDna(true)
	resp := CreateResponse(err)
	assert.NotEmpty(t, resp)
	assert.Equal(t, resp.Error.Code, 409)
	assert.NotEmpty(t, resp.Error)
}
