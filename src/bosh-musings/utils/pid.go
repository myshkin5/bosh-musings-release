package utils

import (
	"errors"
	"os"
	"strconv"
)

func WritePIDFile(name string) error {
	if name == "" {
		return errors.New("Can't write to empty file name")
	}

	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()

	pid := strconv.Itoa(os.Getpid())
	_, err = file.WriteString(pid)
	if err != nil {
		return err
	}

	return nil
}
