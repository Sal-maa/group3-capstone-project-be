package util

import (
	_config "capstone/be/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
)

var sess *session.Session

func GetAWSSession(config *_config.AppConfig) (*session.Session, error) {
	if sess == nil {
		newSession, err := session.NewSession(
			&aws.Config{
				Region: aws.String(config.AWS.Region),
				Credentials: credentials.NewStaticCredentials(
					config.AWS.AccessKeyID,
					config.AWS.SecretKey,
					"",
				),
			},
		)

		if err != nil {
			return nil, err
		}

		sess = newSession
	}

	return sess, nil
}
