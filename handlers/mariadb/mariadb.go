package mariadb
// package main

import (
  "fmt"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
)

// show databases;
// create database golang default character set utf8 default collate utf8_bin;
// GRANT ALL PRIVILEGES ON golang.* to golanguser@'%' IDENTIFIED BY 'golang';
// GRANT ALL PRIVILEGES ON golang.* to golanguser@'localhost' IDENTIFIED BY 'golang';
// CREATE TABLE eventTable(id INT NOT NULL AUTO_INCREMENT, eventid INT NOT NULL, eventname VARCHAR(255), PRIMARY KEY(id));
// INSERT INTO eventTable(eventid, eventname) VALUES(1, 'Start');

type event struct {
  id, eventid int
  eventname string
}

func Selectevent()(event) {
// TODO: Change query to var
  var Event event
  db, err := sql.Open("mysql", "golanguser:golang@/golang")
  if err != nil {
    fmt.Println("Error is", err)
  } else {
    defer db.Close()
    db.QueryRow("SELECT id,eventid,eventname FROM eventTable WHERE id = '1';").Scan(&Event.id, &Event.eventid, &Event.eventname)

  }
  return Event
}


// WIP
func Updateevent()(bool) {
  var result string
  db, err := sql.Open("mysql", "golanguser:golang@/golang")
  defer db.Close()
  if err !=nil {
    fmt.Println("Error is", err)
  } else {
    fmt.Println(db.QueryRow("update id,eventid,eventname FROM eventTable WHERE id = '1';").Scan(&result))
  }
  return true
}
