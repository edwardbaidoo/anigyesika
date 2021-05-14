package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//import config

var db *gorm.DB

type User struct {
	gorm.Model
	LastName         string
	FirstName        string
	OtherNames       string
	Email            string
	PhoneNumber      string
	Password         string
	IsNumberVerified bool
	IsEmailVerified  bool
	Role             string
	LastLogin        time.Time
}

/*
Table=User
--id
firstName
lastName
otherNames
email
phonenumber - E.164 format(15 numbers max)
email
password
isNumberVerified
isEmailVerified
isUserVerified
role
lastLogin
--dateCreated
--dateUpdated
*/

type Transaction struct {
	gorm.Model
	OwnerID      int32
	Amount       float32
	Currency     string
	Verified     bool
	Description  string
	Description2 string
}

/*
Table=Transaction
ownerid
amount
currency
datecreated
dateupdated
verified
description1
description2
*/
