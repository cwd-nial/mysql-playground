package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteDuplicatesJoin(t *testing.T) {
	fr := newFramework(t)

	startMessages := fr.repository.GetAll()
	assert.Equal(t, 15, len(startMessages))

	duplicates := fr.repository.GetDuplicatesJoin()
	assert.Equal(t, 9, len(duplicates))

	affectedRows := fr.repository.DeleteDuplicatesJoin()
	assert.Equal(t, 9, affectedRows)

	finalMessages := fr.repository.GetAll()
	assert.Equal(t, 6, len(finalMessages))
}
