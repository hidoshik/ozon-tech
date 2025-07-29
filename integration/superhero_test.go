package integration_test

import (
	"integration"
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestComparePowers(t *testing.T) {
	res, err := integration.ComparePower("ironman", "wolverine")
	require.NoError(t, err, "Recieved invalid power value")
	assert.Equal(t, res, "ironman")
}