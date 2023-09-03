package server_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ak9024/adiatma.tech/api/internal/app/server"
	"github.com/stretchr/testify/assert"
)

func TestServerRouter(t *testing.T) {
	app := server.Router()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res, err := app.Test(req, 1)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
