package conf

import (
   "fmt"
   "github.com/globalsign/mgo"
   "os"
)

var session *mgo.Session
var db *mgo.Database

func ConnectDB() {
   host := os.Getenv("MONGO_DB_HOST")
   port := os.Getenv("MONGO_PORT")
   session, err := mgo.Dial(host + ":" + port)
   if err != nil {
      fmt.Println(err)
   }

   session.SetMode(mgo.Monotonic, true)

   db = session.DB("test")
}

func GetCollection(collection string) *mgo.Collection {
   return db.C(collection)
}

func CloseSession() {
   session.Close()
}
