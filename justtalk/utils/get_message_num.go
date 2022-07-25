package utils

import (
	"os"
	"path"
)

func GetMessageNum(title string) (int, error) {
	dir := path.Join("data", title)
	_, err := os.Stat(dir)
	if os.IsNotExist(err){
		err = os.Mkdir(dir, 0600)
		if err != nil {
			return 0, err
		}
	}
	files, err := os.ReadDir(dir)
	if err != nil {
		return 0, err
	}
	filesNumber := len(files)
	return filesNumber, nil
}