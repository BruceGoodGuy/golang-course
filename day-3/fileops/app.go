package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const defaultData = 1000

func GetFloatFromFile(fileName string) (float64, error) {
	if strings.Trim(fileName, " ") == "" {
		return defaultData, errors.New("filename is invalid")
	}
	data, error := os.ReadFile(fileName)
	if error != nil {
		return defaultData, errors.New("can't found file")
	}
	dataStr := string(data)
	floatData, error := strconv.ParseFloat(dataStr, 64)
	if error != nil {
		return defaultData, errors.New("can't parse data")
	}

	return floatData, nil

}

func SaveFloatToFile(data float64, fileName string) {
	dataText := fmt.Sprint(data)
	os.WriteFile(fileName, []byte(dataText), 0644)
}
