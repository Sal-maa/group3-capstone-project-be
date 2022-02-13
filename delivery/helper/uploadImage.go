package helper

import (
	_config "capstone/be/config"
	_util "capstone/be/util"

	"errors"
	"fmt"
	"mime/multipart"
	"net/http"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func UploadImage(category string, id int, file *multipart.FileHeader, src multipart.File) (filename string, code int, err error) {
	config := _config.GetConfig()

	// check file extension
	extension, err := checkFileExtension(file.Filename)

	if err != nil {
		return "", http.StatusBadRequest, err
	}

	// check file size
	if err = checkFileSize(file.Size); err != nil {
		return "", http.StatusBadRequest, err
	}

	// rename file for consistency
	file.Filename = fileRenamer(category, id, extension)

	// get session instance
	sess := session.Must(_util.GetAWSSession(config))

	// uploading file
	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(config.AWS.Bucket),
		Key:    aws.String(file.Filename),
		Body:   src,
	})

	// detect failure while uploading file
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return file.Filename, http.StatusOK, nil
}

func checkFileExtension(filename string) (string, error) {
	extension := strings.ToLower(filename[strings.LastIndex(filename, ".")+1:])

	if extension != "jpg" && extension != "jpeg" && extension != "png" {
		return "", errors.New("forbidden file type")
	}

	return extension, nil
}

func checkFileSize(size int64) error {
	if size == 0 {
		return errors.New("illegal file size")
	}

	if size > 2097152 {
		return errors.New("file size too big")
	}

	return nil
}

func fileRenamer(category string, id int, extension string) string {
	return fmt.Sprintf("%s-%d-%d.%s", category, id, time.Now().Unix(), extension)
}
