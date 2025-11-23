package provider

import (
	"context"
	"errors"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/smithy-go"
	log "github.com/sirupsen/logrus"
)

// TODO: If user explicitly set profile name and we don't find creds, probably blow up
// returns:
//   - bool    - if there are credentials configured
//   - string - AccessKeyID if credentials loaded
func HasCredentials(cfg aws.Config) (bool, string) {
<<<<<<< HEAD
	credentials, credsErr := cfg.Credentials.Retrieve(context.TODO())
	if credsErr != nil {
		var oe *smithy.OperationError
		if errors.As(credsErr, &oe) {
			if oe.ServiceID != "ec2imds" || oe.OperationName != "GetMetadata" {
				log.WithFields(log.Fields{"method": "provider.HasCredentials"}).Error(oe.Error())
			}
			return false, ""
		}
	}
	return true, credentials.AccessKeyID
}

func ClientHasCredentials(client *s3.Client) bool {
	_, credsErr := client.Options().Credentials.Retrieve(context.TODO())
	if credsErr != nil {
		var oe *smithy.OperationError
		if errors.As(credsErr, &oe) {
			if oe.ServiceID != "ec2imds" || oe.OperationName != "GetMetadata" {
				log.WithFields(log.Fields{"method": "provider.ClientHasCredentials"}).Error(oe.Error())
			}
			return false
		}
	}
	return true
=======
    creds, err := cfg.Credentials.Retrieve(context.TODO())
    if err != nil {
        var oe *smithy.OperationError
        if errors.As(err, &oe) &&
           !(oe.ServiceID == "ec2imds" && oe.OperationName == "GetMetadata") {
            log.WithField("method", "provider.HasCredentials").Error(oe.Error())
        }
        return false, ""
    }

    // NEW: empty keys mean “no real creds”
    if creds.AccessKeyID == "" || creds.SecretAccessKey == "" {
        return false, ""
    }
    return true, creds.AccessKeyID
}

func ClientHasCredentials(c *s3.Client) bool {
    creds, err := c.Options().Credentials.Retrieve(context.TODO())
    if err != nil {
        var oe *smithy.OperationError
        if errors.As(err, &oe) &&
           !(oe.ServiceID == "ec2imds" && oe.OperationName == "GetMetadata") {
            log.WithField("method", "provider.ClientHasCredentials").Error(oe.Error())
        }
        return false
    }
    // NEW
    return creds.AccessKeyID != "" && creds.SecretAccessKey != ""
>>>>>>> 008a88b (anon is now avaiable)
}

