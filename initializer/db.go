package initializer

import (
  "gorm.io/driver/sqlite" // Sqlite driver based on CGO
  // "github.com/glebarez/sqlite" // Pure-Go SQLite driver, checkout https://github.com/glebarez/sqlite for details
  // "github.com/libtnb/sqlite" // Pure-Go SQLite driver, checkout https://github.com/libtnb/sqlite for details
  "gorm.io/gorm"
  "fmt"
)

var DB *gorm.DB

func ConnectToDatabase(){
  var err error
  // github.com/mattn/go-sqlite3
  DB, err = gorm.Open(sqlite.Open("InvGredi.db"), &gorm.Config{})
  if err != nil {
    fmt.Println("Failed to connect to database :C")
  } 
}
