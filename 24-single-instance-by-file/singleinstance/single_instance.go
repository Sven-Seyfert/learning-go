package singleinstance

import "os"

const instanceFile = "single_instance_indicator"

func existsInstanceFile() bool {
	fileInfo, err := os.Stat(instanceFile)
	if os.IsNotExist(err) {
		return false
	}

	return !fileInfo.IsDir()
}

func CreateInstanceFile() {
	if existsInstanceFile() {
		os.Exit(0)
	}

	file, err := os.Create(instanceFile)
	if err != nil {
		panic(err)
	}

	defer file.Close()
}

func RemoveInstanceFile() {
	err := os.Remove(instanceFile)
	if err != nil {
		panic(err)
	}
}
