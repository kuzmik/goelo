package main

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"github.com/golang/glog"
	_ "github.com/mattn/go-sqlite3"
	"strings"
)

//Bebot is the main brains behind the database access for the irc bot
type Bebot struct{}

var db *sql.DB

//AddLog adds a new message to the database
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
		glog.Fatal(err)
	}

	_, err = insert.Exec(hash, server, channel, nick, message)
	if err != nil {
		if strings.Contains(err.Error(), "UNIQUE constraint failed: logs.hash") {
			glog.Info("Dropping duplicate message")
		} else {
			glog.Fatal(err)
		}
		return false
	}

	return true
}

// GetLog gets a random message for a specific nickname
func (b Bebot) GetLog(nick string) string {
	if db == nil {
		var err error
		db, err = sql.Open("sqlite3", "./data/bot.sqlite3")
		if err != nil {
			glog.Fatal(err)
		}
	}

	var message string
	err := db.QueryRow("SELECT message FROM logs WHERE nick = $1 ORDER BY RANDOM() LIMIT 1", nick).Scan(&message)
	if err != nil {
		glog.Fatal(err)
		return ""
	}
	return message
}
