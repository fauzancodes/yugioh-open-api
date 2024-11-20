package upload

import (
	"context"
	"errors"
	"fmt"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/admin"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/fauzancodes/yugioh-open-api/app/config"
	"github.com/fauzancodes/yugioh-open-api/app/utils"
)

func InitiateCloudinary() (cloud *cloudinary.Cloudinary, cloudName string, err error) {
	cloudName = config.LoadConfig().CloudinaryCloudName
	apiKey := config.LoadConfig().CloudinaryAPIKey
	apiSecret := config.LoadConfig().CLoudinaryAPISecret
	cloud, err = cloudinary.NewFromParams(cloudName, apiKey, apiSecret)
	if err != nil {
		err = errors.New("failed to initiate cloudinary: " + err.Error())
		return
	}

	return
}

func CheckAssetExistanceByPublicID(filename string) (isExist bool, secureUrl string) {
	cloudinaryGuessedUrl := fmt.Sprintf("https://res.cloudinary.com/%v/image/upload/%v/%v.jpg", config.LoadConfig().CloudinaryCloudName, config.LoadConfig().CloudinaryFolder, filename)

	isUrlOK := utils.CheckURLStatus(cloudinaryGuessedUrl)
	if isUrlOK {
		isExist = true
		secureUrl = cloudinaryGuessedUrl

		return
	}

	request, _, err := InitiateCloudinary()
	if err != nil {
		return
	}
	asset, err := request.Admin.Asset(context.TODO(), admin.AssetParams{
		PublicID: config.LoadConfig().CloudinaryFolder + "/" + filename,
	})

	if err == nil && asset.PublicID != "" {
		isExist = true
		secureUrl = asset.SecureURL

		return
	}

	return
}

func UploadImageOrVideo(file interface{}, folder string, filename string) (secureUrl, publicID, cloudName string, err error) {
	if folder != "" {
		folder = config.LoadConfig().CloudinaryFolder + "/" + folder
	} else {
		folder = config.LoadConfig().CloudinaryFolder
	}
	request, cloudName, err := InitiateCloudinary()
	if err != nil {
		return
	}
	response, err := request.Upload.Upload(context.Background(), file, uploader.UploadParams{
		Folder:   folder,
		PublicID: filename,
	})
	if err != nil {
		err = errors.New("failed to upload file to cloudinary: " + err.Error())
		return
	}

	secureUrl = response.SecureURL
	publicID = response.PublicID

	return
}
