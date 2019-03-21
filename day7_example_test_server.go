package calc

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubHandler(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(SubHandler))
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/sub?a=30&b=20", nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json", resp.Header.Get("Content-Type"))
	assert.Equal(t, `{"result":10}`, string(body))
}
