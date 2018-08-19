package database

import (
	"net/http"

	"github.com/api_rest_go/src/util"
	"github.com/globalsign/mgo"
)

func GetSessionMongoDB() *mgo.Session {
	session, err := mgo.Dial("mongodb://localhost")
	util.HandlerError(&err, http.StatusInternalServerError, "Ocorreu um erro.")

	return session
}
