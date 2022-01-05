package tools 

import (
	"os"
)

func CheckDirectory(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil 
	}

	if os.IsNotExist(err) {
		return false, nil 
	}

	return false, err  
}

func CreateOrCheckDirectory(path string) error {
	isExist, err := CheckDirectory(path)
	if err != nil {
		return err 
	}

	if !isExist {
		//create directory
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err 
		}
	}

	return nil 
}