package app_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ak9024/adiatma.tech/api/internal/app"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	app := app.Router()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	res, err := app.Test(req, 1)
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
