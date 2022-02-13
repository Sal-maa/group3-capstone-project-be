package helper

import (
	_config "capstone/be/config"
	_util "capstone/be/util"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func DeleteImage(filename string) error {
	config := _config.GetConfig()

	// get session instance
	sess := session.Must(_util.GetAWSSession(config))

	// deleting file
	svc := s3.New(sess)

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(config.AWS.Bucket),
		Key:    aws.String(filename),
	})

	// detect failure while deleting file
	if err != nil {
		return err
	}

	// completing delete process
	err = svc.WaitUntilObjectNotExists(&s3.HeadObjectInput{
		Bucket: aws.String(config.AWS.Bucket),
		Key:    aws.String(filename),
	})

	// detect failure while deleting file
	if err != nil {
		return err
	}

	return nil
}
