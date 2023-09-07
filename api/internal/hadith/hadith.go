package hadith

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

type hadith struct{}

type IHadith interface {
	GetHadith(c *fiber.Ctx) error
}

func New() IHadith {
	return &hadith{}
}

func (h *hadith) GetHadith(c *fiber.Ctx) error {
	hadith := c.Params("hadith")
	hadiths, errReadHadith := readHadith(hadith)
	if errReadHadith != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(hadiths)
}

func readHadith(hadith string) ([]Hadith, error) {
	var hadiths []Hadith

	getDataHadith := fmt.Sprintf("./data/hadith/%s.json", hadith)
	fileByte, errReadFile := os.ReadFile(getDataHadith)
	if errReadFile != nil {
		return hadiths, errReadFile
	}

	errMarshalData := json.Unmarshal(fileByte, &hadiths)
	if errMarshalData != nil {
		return hadiths, errMarshalData
	}

	return hadiths, nil
}
