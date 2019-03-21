package calc

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetGinRouter(r *gin.Engine) {
	r.GET("/add/:a/:b", AddGinHandler)
}

func TestAddGinRouter(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	SetGinRouter(r)
	ts := httptest.NewServer(r)
	defer ts.Close()

	req, _ := http.NewRequest("GET", ts.URL+"/add/10/20", nil)
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "application/json; charset=utf-8", resp.Header.Get("Content-Type"))
	assert.Equal(t, `{"result":30}`, string(body))
}
