package imageClient

import (
	"encoding/xml"
	"errors"
	"fmt"
	azure "github.com/uzresk/azure-sdk-for-go"
)

const (
	azureImageListURL      = "services/images"
	invalidImageError      = "Can not find image %s in specified subscription, please specify another image name."
	paramNotSpecifiedError = "Parameter %s is not specified."
)

func GetImageList() (ImageList, error) {
	imageList := ImageList{}

	response, err := azure.SendAzureGetRequest(azureImageListURL)
	if err != nil {
		return imageList, err
	}

	err = xml.Unmarshal(response, &imageList)
	if err != nil {
		return imageList, err
	}

	return imageList, err
}

func ResolveImageName(imageName string) error {
	if len(imageName) == 0 {
		return fmt.Errorf(paramNotSpecifiedError, "imageName")
	}

	imageList, err := GetImageList()
	if err != nil {
		return err
	}

	for _, image := range imageList.OSImages {
		if image.Name != imageName && image.Label != imageName {
			continue
		}

		return nil
	}

	return errors.New(fmt.Sprintf(invalidImageError, imageName))
}
