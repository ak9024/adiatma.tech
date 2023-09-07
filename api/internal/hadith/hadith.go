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
	hadiths, errReadHadith := ReadHadith(hadith, "./data/hadith")
	if errReadHadith != nil {
		return c.SendString(errReadHadith.Error())
	}

	return c.Status(fiber.StatusOK).JSON(hadiths)
}

func ReadHadith(hadith, filePath string) ([]Hadith, error) {
	var hadiths []Hadith

	getDataHadith := fmt.Sprintf("%s/%s.json", filePath, hadith)
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
