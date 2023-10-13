package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteDuplicatesCte(t *testing.T) {
	fr := newFramework(t)

	startMessages := fr.repository.GetAll()
	assert.Equal(t, 10, len(startMessages))

	duplicates := fr.repository.GetDuplicatesCte()
	assert.Equal(t, 5, len(duplicates))

	affectedRows := fr.repository.DeleteDuplicatesCte()
	assert.Equal(t, 5, affectedRows)

	finalMessages := fr.repository.GetAll()
	assert.Equal(t, 5, len(finalMessages))
}
