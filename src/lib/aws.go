package lib

import "github.com/aws/aws-sdk-go/aws/session"

func AWSSessions() (*session.Session, error) {
	sess, err := session.NewSession()
	svc := session.Must(sess, err)
	return svc, err
}
