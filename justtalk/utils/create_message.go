package utils

import (
	"errors"
	"os"
	"path"
	"strconv"
)

func CreateMessage(title string, body string) error {
	dir := path.Join("data", title)
	id, err := GetMessageNum(title)
	if err != nil {
		return err
	}
	filename := path.Join(dir, strconv.Itoa(id))
	_, err = os.Stat(filename)
	if !os.IsNotExist(err) {
		return errors.New("message exists")
	} else {
		_, err = os.Stat(dir)
		if os.IsNotExist(err) {
			err = os.Mkdir(dir, 0600)
			if err != nil {
				return err
			}
		}
		return os.WriteFile(filename, []byte(body), 0600)
	}
}
