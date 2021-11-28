package utils

import (
	"fmt"
	"io/ioutil"
)

const (
	mountIdLocation          = "/var/lib/docker/image/overlay2/layerdb/mounts/%s/mount-id"
	targetFileSystemLocation = "/var/lib/docker/overlay2/%s/merged"
)

func GetTargetFileSystemLocation(containerId string) (string, error) {
	fmt.Printf("GetTargetFileSystemLocation with containerId = %s\n", containerId)
	fileName := fmt.Sprintf(mountIdLocation, containerId)
	mountId, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	targetFileSystemLocation := fmt.Sprintf(targetFileSystemLocation, string(mountId))

	fmt.Printf("GetTargetFileSystemLocation output with targetFileSystemLocation= %s, \nfileName = %s\n\n", targetFileSystemLocation, fileName)

	return targetFileSystemLocation, nil
}
