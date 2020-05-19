package gce

import (
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func getGCE(t *testing.T, controllerID string) *gceClient {
	err := os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/home/bcox/google/milpa-207719-fb309bd448d5.json")
	assert.NoError(t, err)
	c, err := NewGCEClient(controllerID, "bcoxtest", "milpa-207719", WithVPCName("default"), WithZone("us-central1-a"), WithSubnetName("default"))
	assert.NoError(t, err)
	return c
}

func TestWaitForBackoff(t *testing.T) {
	tests := []struct {
		i   int
		exp time.Duration
	}{
		{i: 0, exp: 1},
		{i: 1, exp: 1},
		{i: 3, exp: 3},
		{i: 4, exp: 5},
		{i: 5, exp: 5},
		{i: 6, exp: 5},
	}
	for _, tc := range tests {
		res := waitBackoff(tc.i)
		assert.Equal(t, tc.exp*time.Second, res)
	}
}