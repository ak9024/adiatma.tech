package hadith_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/ak9024/adiatma.tech/api/internal/app/server"
	"github.com/ak9024/adiatma.tech/api/internal/hadith"
	"github.com/stretchr/testify/assert"
)

var pathHadith = "../../data/hadith"

func TestHadith(t *testing.T) {
	app := server.Router()
	req := httptest.NewRequest(http.MethodGet, "/hadith/abu-daud", nil)
	res, err := app.Test(req, 1)
	if err != nil {
		assert.NotNil(t, err)
		t.Log(err.Error())
	}

	t.Run("GET /hadith/abu-daud 200 ok", func(t *testing.T) {
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 200, res.StatusCode)
		t.Log(res.Status)
	})

	t.Run("ReadFileHadith", func(t *testing.T) {
		hadiths, errReadHadiths := hadith.ReadFileHadith("abu-daud", pathHadith)
		if errReadHadiths != nil {
			assert.NotNil(t, errReadHadiths)
		}
		assert.Nil(t, errReadHadiths)
		assert.NotNil(t, hadiths)
	})

	t.Run("HadithPagination", func(t *testing.T) {
		hadiths, errReadHadiths := hadith.ReadFileHadith("abu-daud", pathHadith)
		if errReadHadiths != nil {
			assert.NotNil(t, errReadHadiths)
		}
		assert.Nil(t, errReadHadiths)

		page := 1    // set page to 1
		perPage := 5 // set perPage to 5

		hadithPagination, errHadithPagination := hadith.HadithPagination(page, perPage, hadiths)
		if errHadithPagination != nil {
			assert.NotNil(t, errHadithPagination)
		}
		assert.Nil(t, errHadithPagination)
		assert.NotNil(t, hadithPagination)
	})
}

func TestAuthor(t *testing.T) {
	app := server.Router()
	req := httptest.NewRequest(http.MethodGet, "/hadith/list/authors", nil)
	res, err := app.Test(req, 1)
	if err != nil {
		assert.NotNil(t, err)
		t.Log(err.Error())
	}

	t.Run("GET /hadith/list/authors 200 ok", func(t *testing.T) {
		assert.Nil(t, err)
		assert.NotNil(t, res)
		assert.Equal(t, 200, res.StatusCode)
		t.Log(res.Status)
	})

	fileNames, errReadDirHadith := hadith.ReadDirHadith(pathHadith)
	if errReadDirHadith != nil {
		assert.NotNil(t, errReadDirHadith)
	}

	t.Run("ReadDirHadith", func(t *testing.T) {
		assert.Nil(t, errReadDirHadith)
		assert.NotNil(t, fileNames)

		// loop fileNames and remove the extention
		for _, fileName := range fileNames {
			removeFileJson := strings.Replace(fileName, ".json", "", -1)
			t.Log(removeFileJson)
		}
	})

	authors, errAuthors := hadith.RemoveExtention(fileNames, pathHadith)
	if errAuthors != nil {
		assert.NotNil(t, errAuthors)
	}

	t.Run("RemoveExtention", func(t *testing.T) {
		assert.Nil(t, errAuthors)
		assert.NotNil(t, authors)
		t.Log(authors)
	})
}
