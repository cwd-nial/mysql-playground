package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteDuplicates(t *testing.T) {
	fr := newFramework(t)

	_, messages := fr.repository.GetAll()
	assert.Equal(t, 4, len(messages))

	_, duplicates := fr.repository.GetDuplicatesJoin()
	assert.Equal(t, 2, len(duplicates))
}
