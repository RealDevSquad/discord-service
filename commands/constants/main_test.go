package constants

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandsDefinition(t *testing.T) {
	t.Run("should have verify command without options", func(t *testing.T) {
		verifyCommand := Commands[2]
		assert.NotNil(t, verifyCommand)
		assert.Equal(t, "verify", verifyCommand.Name)
		assert.Equal(t, "Generate a link with user specific token to link with RDS backend", verifyCommand.Description)
		assert.Nil(t, verifyCommand.Options)
	})

	t.Run("should have listening command with options", func(t *testing.T) {
		listeningCommand := Commands[1]
		assert.NotNil(t, listeningCommand)
		listeningCommandOptions := listeningCommand.Options
		assert.NotNil(t, listeningCommandOptions)
		assert.Equal(t, "value", listeningCommandOptions[0].Name)
		assert.Equal(t, "to enable or disable the listening mode", listeningCommandOptions[0].Description)
	})

	t.Run("should have a hello command", func(t *testing.T) {
		helloCommand := Commands[0]
		assert.NotNil(t, helloCommand)
		assert.Equal(t, "hello", helloCommand.Name)
		assert.Equal(t, "Greets back with hello!", helloCommand.Description)
	})
}
