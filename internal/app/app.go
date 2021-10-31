package app

import "github.com/sirupsen/logrus"

func Run() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	server := new(Server)
	if err := server.Run("8000", nil); err != nil {
		logrus.Fatalf("error start server : %s", err.Error())
	}

}
