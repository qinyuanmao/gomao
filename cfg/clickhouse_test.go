package cfg

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestClickhouse(t *testing.T) {
	LoadConfig("./config")
	_, err := NewClickHouseDB("clickhouse")
	assert.Nil(t, err)
}
