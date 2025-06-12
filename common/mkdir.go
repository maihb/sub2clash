package common

import (
	"os"
)

func MKDir(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err := os.MkdirAll(dir, os.ModePerm)
		if err != nil {

			return err
		}
	}
	return nil
}

func MkEssentialDir() error {
	if err := MKDir("subs"); err != nil {
		return NewDirCreationError("subs", err)
	}
	if err := MKDir("logs"); err != nil {
		return NewDirCreationError("logs", err)
	}
	if err := MKDir("data"); err != nil {
		return NewDirCreationError("data", err)
	}
	return nil
}
