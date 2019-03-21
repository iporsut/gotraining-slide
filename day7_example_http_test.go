package calc

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "http://test.com/add?a=10&b=20", nil)
	w := httptest.NewRecorder()

	AddHandler(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, `{"result":30}`, string(body))
}
