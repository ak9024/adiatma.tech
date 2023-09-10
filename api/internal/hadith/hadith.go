package hadith

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type hadith struct{}

// contract of hadith
type IHadith interface {
	GetHadith(c *fiber.Ctx) error
}

func New() IHadith {
	return &hadith{}
}

// Get Hadith API
// @descriptipn Get collection of hadith filter by author
// @Param author path string true "author"
// @Param page query string false "page"
// @Param perPage query string false "perPage"
// @Router /hadith/{author} [get]
// @Success 200 {object} []Hadith
func (h *hadith) GetHadith(c *fiber.Ctx) error {
	author := c.Params("author")

	// get all queries
	queries := c.Queries()

	var defaultPage int
	var defaultPerPage int

	// convert page string to int
	page, errPage := strconv.Atoi(queries["page"])
	if errPage != nil {
		defaultPage = 1 // if errPage or get empty from query string, set default page to 1
	} else {
		defaultPage = page
	}

	// convert perPage string to int
	perPage, errPerPage := strconv.Atoi(queries["perPage"])
	if errPerPage != nil {
		defaultPerPage = 5 // if errPerPage or get empty from query string, set default perPage to 5
	} else {
		defaultPerPage = perPage
	}

	// read hadith's from ./data/hadith/:author's
	hadiths, errReadHadith := ReadFileHadith(author, "./data/hadith")
	if errReadHadith != nil {
		return c.SendString(errReadHadith.Error())
	}

	// convert return of hadith's with pagination to limit return of data
	hadithPagination, errHadithPagination := HadithPagination(defaultPage, defaultPerPage, hadiths)
	if errHadithPagination != nil {
		return c.SendString(errHadithPagination.Error())
	}

	return c.Status(fiber.StatusOK).JSON(hadithPagination)
}

// implement pagination to limit data return of hadith's
func HadithPagination(page, perPage int, hadiths []Hadith) ([]Hadith, error) {
	start := (page - 1) * perPage
	limit := start + perPage

	// if start data have more than data hadith's set return to nil
	if start > len(hadiths) {
		return nil, errors.New("hadith not found!")
	}

	// if data limit more than data hadith's set limit with total count of hadith's
	if limit > len(hadiths) {
		limit = len(hadiths)
	}

	return hadiths[start:limit], nil
}

// handle read file hadith form path data
// ./data/hadith/:author
func ReadFileHadith(author, filePath string) ([]Hadith, error) {
	var hadiths []Hadith

	// get data hadith from file path
	getDataHadith := fmt.Sprintf("%s/%s.json", filePath, author)
	fileByte, errReadFile := os.ReadFile(getDataHadith)
	if errReadFile != nil {
		return hadiths, errReadFile
	}

	// unmarshal data to struct array of hadith
	errMarshalData := json.Unmarshal(fileByte, &hadiths)
	if errMarshalData != nil {
		return hadiths, errMarshalData
	}

	return hadiths, nil
}
