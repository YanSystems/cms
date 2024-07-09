package utils

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnectToDB(t *testing.T) {
	t.Run("Environment variable not set", func(t *testing.T) {
		os.Setenv("YAN_CMS_DB_URI", "")
		client, err := ConnectToDB()
		assert.Nil(t, client)
		assert.EqualError(t, err, "set your 'YAN_CMS_DB_URI' environment variable")
	})

	t.Run("Connection error", func(t *testing.T) {
		os.Setenv("YAN_CMS_DB_URI", "invalid_uri")
		client, err := ConnectToDB()
		assert.Nil(t, client)
		assert.Error(t, err)
	})
}
