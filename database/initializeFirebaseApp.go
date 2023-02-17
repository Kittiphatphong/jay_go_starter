package database

import (
	"context"
	firebase "firebase.google.com/go"
	"go_starter/logs"
	"google.golang.org/api/option"
)

func InitializeFirebaseApp() (app *firebase.App, err error) {
	credentialsFile := option.WithCredentialsFile("./naga-app-firebase-firebase-adminsdk-ipabh-2f68ffa93c.json")
	app, err = firebase.NewApp(context.Background(), nil, credentialsFile)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	return app, nil

}
