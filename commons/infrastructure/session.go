package infrastructure

import (
	"ejercicio/commons/utils"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

//GetSession is a singleton than define the type of session to invoke a client. These types are local and aws session.
func getSession()(*session.Session, error){
	var once sync.Once
	var sess *session.Session
	var err error

	once.Do(func() {

		if os.Getenv("local") == "true"{
			logrus.Info("running in local Session")

			profile := "default"
			sess, err = session.NewSessionWithOptions(session.Options{
				Profile: profile,
				Config: aws.Config{
					Region: aws.String(utils.DefaultRegion),
				},
				SharedConfigState: session.SharedConfigEnable,
			})

		}else{
			logrus.Info("running in AWS Session")
			sess, err = session.NewSession()
		}
	},
	)

	return sess, err
}

