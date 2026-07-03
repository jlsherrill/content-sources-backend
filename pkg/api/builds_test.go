package api

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildResponseJSON(t *testing.T) {
	response := BuildResponse{
		Results: []BuildItem{
			{
				Group:     "io.smallrye.reactive",
				Name:      "smallrye-mutiny-vertx-core",
				Version:   "3.16.0",
				Release:   "rhlw-3002",
				Filename:  "smallrye-mutiny-vertx-core-3.16.0.rhlw-3002.pom",
				CreatedAt: "2024-02-01T14:20:00Z",
			},
		},
		Total:  1,
		Limit:  100,
		Offset: 0,
	}

	jsonData, err := json.Marshal(response)
	assert.Nil(t, err)
	assert.Contains(t, string(jsonData), "io.smallrye.reactive")
	assert.Contains(t, string(jsonData), "rhlw-3002")
}
