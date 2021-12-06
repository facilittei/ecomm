package providers

import (
	"errors"
	"github.com/facilittei/ecomm/internal/mocks"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestJuno_Authenticate_env_required(t *testing.T) {
	type result struct {
		token JunoAuth
		error error
	}

	tests := []struct {
		name string
		args map[string]string
		want result
	}{
		{
			name: "JUNO_AUTHORIZATION_URL is missing",
			args: map[string]string{"JUNO_AUTHORIZATION_URL": ""},
			want: result{
				token: JunoAuth{},
				error: authError("env required: JUNO_AUTHORIZATION_URL"),
			},
		},
		{
			name: "JUNO_AUTHORIZATION_BASIC is missing",
			args: map[string]string{
				"JUNO_AUTHORIZATION_URL":   "https://url.test",
				"JUNO_AUTHORIZATION_BASIC": "",
			},
			want: result{
				token: JunoAuth{},
				error: authError("env required: JUNO_AUTHORIZATION_BASIC"),
			},
		},
	}

	httpClient := &mocks.HttpClientMock{}
	httpClient.On("Post").Return(nil, nil)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for k, v := range tt.args {
				err := os.Setenv(k, v)
				if err != nil {
					t.Fatalf("os.Setenv(%s, %s) error: %v", k, v, err)
				}

				junoProvider := NewJuno(httpClient)
				_, got := junoProvider.Authenticate()
				assert.Equal(t, got, tt.want.error)
				assert.True(t, errors.As(got, &JunoAuthError{}))
			}
		})
	}
}
