package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {

	listInterface := New()

	assert.NotNil(t, listInterface)

	listObj, ok := listInterface.(*listContainer)

	assert.True(t, ok)
	assert.NotNil(t, listObj)
	assert.Nil(t, listObj.first)
	assert.Nil(t, listObj.last)
}
