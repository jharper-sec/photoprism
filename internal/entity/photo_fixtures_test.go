package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPhotoMap_Get(t *testing.T) {
	t.Run("get existing photo", func(t *testing.T) {
		r := PhotoFixtures.Get("19800101_000002_D640C559")
		assert.Equal(t, "pt9jtdre2lvl0yh7", r.PhotoUID)
		assert.Equal(t, "19800101_000002_D640C559", r.PhotoName)
		assert.IsType(t, Photo{}, r)
	})
	t.Run("get not existing photo", func(t *testing.T) {
		r := PhotoFixtures.Get("TestName")
		assert.Equal(t, "TestName", r.PhotoName)
		assert.IsType(t, Photo{}, r)
	})
}

func TestPhotoMap_Pointer(t *testing.T) {
	t.Run("get existing photo pointer", func(t *testing.T) {
		r := PhotoFixtures.Pointer("19800101_000002_D640C559")
		assert.Equal(t, "pt9jtdre2lvl0yh7", r.PhotoUID)
		assert.Equal(t, "19800101_000002_D640C559", r.PhotoName)
		assert.IsType(t, &Photo{}, r)
	})
	t.Run("get not existing photo pointer", func(t *testing.T) {
		r := PhotoFixtures.Pointer("TestName2")
		assert.Equal(t, "TestName2", r.PhotoName)
		assert.IsType(t, &Photo{}, r)
	})
}
