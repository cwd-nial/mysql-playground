package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteDuplicatesJoin(t *testing.T) {
	fr := newFramework(t)

	startMessages := fr.repository.GetAll()
	assert.Equal(t, 10, len(startMessages))

	duplicates := fr.repository.GetDuplicatesJoin()
	assert.Equal(t, 5, len(duplicates))

	affectedRows := fr.repository.DeleteDuplicatesJoin()
	assert.Equal(t, 5, affectedRows)

	finalMessages := fr.repository.GetAll()
	assert.Equal(t, 5, len(finalMessages))
}
