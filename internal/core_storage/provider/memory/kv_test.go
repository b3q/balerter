package memory

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMemory_Put(t *testing.T) {
	m := &storageKV{kv: map[string]string{}}

	err := m.Put("foo", "bar")
	require.NoError(t, err)

	err = m.Put("foo", "bar2")
	require.Error(t, err)

	v, ok := m.kv["foo"]
	assert.True(t, ok)
	assert.Equal(t, "bar", v)
}

func TestMemory_Upsert(t *testing.T) {
	m := &storageKV{kv: map[string]string{}}

	err := m.Upsert("foo", "bar")
	require.NoError(t, err)

	err = m.Upsert("foo", "bar2")
	require.NoError(t, err)

	v, ok := m.kv["foo"]
	assert.True(t, ok)
	assert.Equal(t, "bar2", v)
}

func TestMemory_Delete(t *testing.T) {
	m := &storageKV{kv: map[string]string{}}

	m.kv["foo"] = "bar"

	err := m.Delete("foo2")
	require.Error(t, err)

	err = m.Delete("foo")
	require.NoError(t, err)

	_, ok := m.kv["foo"]
	assert.False(t, ok)

	err = m.Delete("foo")
	require.Error(t, err)
}

func TestMemory_Get(t *testing.T) {
	m := &storageKV{kv: map[string]string{}}

	_, err := m.Get("foo")
	require.Error(t, err)

	m.kv["foo"] = "bar"

	v, err := m.Get("foo")
	require.NoError(t, err)

	assert.Equal(t, "bar", v)
}

func TestMemory_All(t *testing.T) {
	m := &storageKV{kv: map[string]string{
		"f1": "v1",
		"f2": "v2",
	}}

	result, err := m.All()
	require.NoError(t, err)
	assert.Equal(t, 2, len(result))

	v, ok := result["f1"]
	assert.True(t, ok)
	assert.Equal(t, "v1", v)

	v, ok = result["f2"]
	assert.True(t, ok)
	assert.Equal(t, "v2", v)
}
