package service

import (
	mysqldbmodel "gql/mysqldb/model"
	"gql/utils"
)

func CreateUser(uid *string) {
}

func GetUserByUid(uid *string) *mysqldbmodel.User {
	db := utils.GetSQLDB()

	user := mysqldbmodel.User{}

	db.Where(&mysqldbmodel.User{Uid: *uid}).First(&user)

	return &user
}
