package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/upper/db/v4/adapter/mysql"
)

var settings = mysql.ConnectionURL{
	Database: `db`,
	Host:     `localhost`,
	User:     `root`,
	Password: `password`,
}

type User struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	CreatedAt string `db:"createdAt"`
}

type Profile struct {
	ID        int    `db:"id"`
	Address   string `db:"address"`
	UserID    int    `db:"userId"`
	CreatedAt string `db:"createdAt"`
}

type UserWithProfile struct {
	User    User    `db:",inline"`
	Profile Profile `db:",inline"`
}

func main() {
	sess, err := mysql.Open(settings)
	if err != nil {
		log.Fatal("Open: ", err)
	}
	defer sess.Close()

	var users []UserWithProfile
	err = sess.
		SelectFrom("user").
		Join("profile").On("user.id = profile.userId").All(&users)
	if err != nil {
		log.Fatal("Collections: ", err)
	}
	spew.Dump(users)
}
