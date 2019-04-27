package conf

import (
   "fmt"
   "github.com/globalsign/mgo"
)
var session *mgo.Session
var db *mgo.Database
func ConnectDB() {
   session, err := mgo.Dial("mongodb:27017")
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