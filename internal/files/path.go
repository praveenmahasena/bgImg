package files

import (
	"errors"
	"fmt"
	"os"
)

func ReadDir() error {
	// path, pathErr := getPath()
	//
	// if pathErr != nil {
	// 	return pathErr
	// }

	files, fileErr := os.ReadDir("/home/cj/Desktop/bg-img")

	if fileErr != nil {
		return errors.New("Error during Navigating to the dir you provided make sure the Dir exists")
	}

	imgFile, err := selectFunc(files)
	if err != nil {
		return err
	}

	fmt.Println(imgFile)

	return nil
}

func getPath() (string, error) {
	if len(os.Args) >= 2 {
		return os.Args[1], nil
	}
	// return "", fmt.Errorf("Error : did not provide proper absulote path to your background image dir")
	return "/home/cj/Desktop/bg-img", nil
}
