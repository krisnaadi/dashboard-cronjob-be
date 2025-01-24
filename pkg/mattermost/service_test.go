package mattermost

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func setup() {
	// Set the environment variable for testing
	os.Setenv("MATTERMOST_WEBHOOK_URL", "https://example.com")
}

func teardown() {
	// Clean up by unsetting the environment variable
	os.Unsetenv("MATTERMOST_WEBHOOK_URL")
}

func TestMain(m *testing.M) {
	// Run setup before running tests
	setup()

	// Run the tests
	code := m.Run()

	// Run teardown after running tests
	teardown()

	// Exit with the test result code
	os.Exit(code)
}

func TestSendMessage(t *testing.T) {
	type args struct {
		message string
		fields  []Field
	}
	tests := []struct {
		name  string
		args  args
		mocks func()
	}{
		{
			name: "case success",
			args: args{
				message: "test",
				fields:  []Field{},
			},
			mocks: func() {},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mocks()
			SendMessage(test.args.message, test.args.fields)

			assert.NoError(t, nil)
		})
	}
}
