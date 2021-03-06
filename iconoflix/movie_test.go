// © 2018 Imhotep Software LLC. All rights reserved.
package iconoflix_test

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/derailed/pkg/iconoflix"
	"github.com/stretchr/testify/assert"
)

type (
	MockClient struct{}

	nopCloser struct {
		io.Reader
	}
)

func (nopCloser) Close() error { return nil }

func (m *MockClient) Do(r *http.Request) (*http.Response, error) {
	resp := &http.Response{}

	resp.StatusCode = http.StatusOK
	resp.Body = nopCloser{bytes.NewBufferString(`{"status": 200}`)}

	return resp, nil
}

func TestLoadFile(t *testing.T) {
	data, err := iconoflix.LoadFile("./assets/test.yml")
	assert.Nil(t, err)

	assert.Equal(t, len(data.Movies), 2)
	assert.Equal(t, "m1", data.Movies[0].Name)
	assert.Equal(t, "ic1", data.Movies[0].Icons[0].Emoji)
}

func TestLoadFileFail(t *testing.T) {
	_, err := iconoflix.LoadFile("./assets/blee.yml")
	assert.NotNil(t, err)
}

func TestLoadMem_BoxOffice(t *testing.T) {
	data, err := iconoflix.LoadMem("v1")
	assert.Nil(t, err)

	assert.Equal(t, len(data.Movies), 11)
	assert.Equal(t, "Home Alone", data.Movies[0].Name)
	assert.Equal(t, "🏡", data.Movies[0].Icons[0].Emoji)
}

func TestLoadMem_BMovies(t *testing.T) {
	data, err := iconoflix.LoadMem("v2")
	assert.Nil(t, err)

	assert.Equal(t, len(data.Movies), 7)
	assert.Equal(t, "Cobra", data.Movies[0].Name)
	assert.Equal(t, "🕶", data.Movies[0].Icons[0].Emoji)
}

func TestRandMovie(t *testing.T) {
	m := iconoflix.RandMovie("v1")
	assert.NotEqual(t, "", m.Name)
}

func TestCall(t *testing.T) {
	err := iconoflix.Call(&MockClient{}, "GET", "/test", nil, nil, nil)
	assert.Nil(t, err)
}

func TestCallCookie(t *testing.T) {
	c := []*http.Cookie{
		&http.Cookie{Name: "test", Value: "blee"},
	}

	err := iconoflix.Call(&MockClient{}, "GET", "/test", nil, nil, c)
	assert.Nil(t, err)
}
