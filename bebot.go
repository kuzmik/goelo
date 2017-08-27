package main

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"github.com/golang/glog"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

// Usage:
// b := Bebot{}
// added := b.AddLog("testbot", "#test", "nick", "hello world3")
// fmt.Println("Rows added:", added)
type Bebot struct{}

var db *sql.DB

//Adds a new message to the database
func (b Bebot) AddLog(server string, channel string, nick string, message string) bool {
	if db == nil {
		var err error
		db, err = sql.Open("sqlite3", "./data/bot.sqlite3")
		if err != nil {
			glog.Fatal(err)
		}
	}

	// hash 'nick - message' to put in our hash field to avoid duplicate entries.
	// it's how we do unique indexes in sqlite :|
	msg := []byte(fmt.Sprintf("%s - %s", nick, message))
	hash := fmt.Sprintf("%x", sha1.Sum(msg))

	insert, err := db.Prepare("INSERT INTO logs (hash, server, channel, nick, message) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		fmt.Println(err)
		glog.Fatal(err)
	}

	_, err = insert.Exec(hash, server, channel, nick, message)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed: logs.hash") {
			glog.Info("Dropping duplicate message")
			return false
		} else {
			fmt.Println(err)
			glog.Fatal(err)
		}
	}

	return true
}

// Get logs for a specific nickname
func (b Bebot) GetLog(nick string) string {
	if db == nil {
		var err error
		db, err = sql.Open("sqlite3", "./data/bot.sqlite3")
		if err != nil {
			glog.Fatal(err)
		}
	}

	rows, err := db.Query("SELECT id, server, channel, message FROM logs WHERE nick = $1 ORDER BY RANDOM() LIMIT 1", nick)
	if err != nil {
		glog.Fatal(err)
	}

	for rows.Next() {
		var id int
		var server, channel, message string
		if err := rows.Scan(&id, &server, &channel, &message); err != nil {
			glog.Fatal(err)
		}
		fmt.Printf("id:%d - [%s/%s] (%s) %s\n", id, server, channel, nick, message)
		return message
	}
	if err := rows.Err(); err != nil {
		glog.Fatal(err)
	}
	return ""
}
