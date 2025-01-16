package testhelpers

import (
	"gorm.io/gorm"
	"io/ioutil"
	"strings"
)

func PrepareDataForDb(DB *gorm.DB, path string) error {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	requests := strings.Split(string(file), ";")

	for _, request := range requests {
		err = DB.Exec(request).Error
		if err != nil {
			return err
		}

	}
	return nil
}
