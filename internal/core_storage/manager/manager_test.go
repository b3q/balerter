package manager

import (
	"github.com/balerter/balerter/internal/core_storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestManager_Get(t *testing.T) {
	m := &Manager{map[string]core_storage.CoreStorage{
		"foo": nil,
	}}

	_, err := m.Get("foo")
	assert.NoError(t, err)

	_, err = m.Get("bar")
	require.Error(t, err)
	assert.Equal(t, "storage not found", err.Error())
}
