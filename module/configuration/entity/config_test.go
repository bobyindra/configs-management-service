package entity_test

import (
	"testing"
	"time"

	"github.com/bobyindra/configs-management-service/module/configuration/entity"
	"github.com/stretchr/testify/assert"
)

func TestEntity_Config(t *testing.T) {
	t.Parallel()
	// Given set data to entity config
	data := &entity.Config{
		Id:           1,
		Name:         "Robert",
		ConfigValues: "Hello",
		Version:      1,
		CreatedAt:    time.Now().UTC(),
		ActorId:      2,
	}

	// When call ToResponse() function
	res := data.ToResponse()

	// Then all data should return correctly
	assert.Equal(t, data.Id, res.Id)
	assert.Equal(t, data.Name, res.Name)
	assert.Equal(t, data.ConfigValues, res.ConfigValues)
	assert.Equal(t, data.Version, res.Version)
	assert.Equal(t, data.CreatedAt, res.CreatedAt)
	assert.Equal(t, data.ActorId, res.ActorId)
}
