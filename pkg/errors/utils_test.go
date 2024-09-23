package errors

import (
	"fmt"
	"testing"
)

func TestGetHTTPError(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want HTTPError
	}{
		{
			name: "invalid input",
			err:  ErrInvalidInput,
			want: HTTPError{
				Code:     "400001",
				Message:  "one of the input is invalid",
				HTTPCode: 400,
			},
		},
		{
			name: "resource not found",
			err:  ErrResourceNotFound,
			want: HTTPError{
				Code:     "404001",
				Message:  "resource not found",
				HTTPCode: 404,
			},
		},
		{
			name: "internal server error",
			err:  ErrInternalServerError,
			want: HTTPError{
				Code:     "500001",
				Message:  "internal server error",
				HTTPCode: 500,
			},
		},
		{
			name: "not _err struct error",
			err:  fmt.Errorf("not _err struct error"),
			want: HTTPError{
				Code:     "500001",
				Message:  "internal server error",
				HTTPCode: 500,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetHTTPError(tt.err); got != tt.want {
				t.Errorf("GetHTTPError() = %v, want %v", got, tt.want)
			}
		})
	}
}
