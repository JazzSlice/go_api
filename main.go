

package main

import (
  "database/sql"
  "fmt"
  "os"
  //"strconv"
  //"context"
  "net/http"

  _ "github.com/lib/pq"
  "github.com/blockloop/scan/v2"
  "github.com/gin-gonic/gin"
)

var db *sql.DB

const (
  host     = "127.0.0.1"
  port     = 5432
  user     = "worker"
  password = "example"
  dbname   = "recordings"
)

type Album struct {
  ID	  int64
  Title	  string
  Artist  string
  Price	  float32
}

//func allAlbumsQuery() {
//  var albums []Album
//  rows, err := db.Query("SELECT * FROM album")
//  if err != nil {
//    return nil, fmt.Errorf("allAlbumsQuery %q: %v", name, err)
//  }
//  err := scan.Rows(&albums, rows)
//
//  return albums
//}

func getAlbums(c *gin.Context) {
  psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable",
    os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DB_NAME"))
    //host, port, user, password, dbname)

  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }

  rows, err := db.Query("SELECT * FROM album")
  //if err != nil {
    //return nil, fmt.Errorf("albums: %v", err)
  //}
  //defer rows.Close()
  var albums []Album
  err = scan.Rows(&albums, rows)
  defer db.Close()

  c.IndentedJSON(http.StatusOK, albums)
}

//func getAlbumsByID(c *gin.Context) {
  //name := c.Param("artist")
  //rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
  
//}

func main() {
  psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
    "password=%s dbname=%s sslmode=disable",
    os.Getenv("PG_HOST"), os.Getenv("PG_PORT"), os.Getenv("PG_USER"), os.Getenv("PG_PASSWORD"), os.Getenv("PG_DB_NAME"))
    //host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  if err = db.Ping(); err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")


  router := gin.Default()
  router.GET("/albums", getAlbums)
  //router.GET("/albums/:id"getAlbumsByID)

  router.Run("0.0.0.0:8080")

}
