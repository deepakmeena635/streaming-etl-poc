package settings

import (
	"github.com/aws/aws-sdk-go/aws/session"
)

type settings struct {
	dynoSession    session.Session
	rackSampleSize int
}
