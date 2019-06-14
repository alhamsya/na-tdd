package main

import (
	"testing"

	"github.com/jarcoal/httpmock"
)

func Test_getMovie(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", `=~http://www.omdbapi.com/`,
		httpmock.NewStringResponder(503, `Server Error`))

	t.Run("server_error", func(t *testing.T) {
		data, err := getMovie("Avengers: Endgame")
		if err != nil {
			t.Errorf("should return error but instead return %v", data)
		}
	})
}
