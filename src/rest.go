package main

// import (
// 	"net/http"
//     "app/src/conf"

// 	"github.com/labstack/echo"
//     _ "github.com/go-sql-driver/mysql"

// )

// type User struct {
//     ID int `json:"id"`
//     Name string `json:"name"`
// }

// type Users []User

// func main() {
// 	e := echo.New()

// 	// fmt.Println(users)
// 	e.GET("/", func(c echo.Context) error {
//         return c.JSON(http.StatusOK, getUser())
//         // return c.String(http.StatusOK, "mori")
// 	})

// 	e.Logger.Fatal(e.Start(":1323"))
// }

// func getUser() Users {
//     db, err := conf.Init()
//     defer db.Close()

// 	rows, err := db.Query("SELECT * FROM users") //
// 	if err != nil {
// 	  panic(err.Error())
// 	}
// 	users := Users{}
// 	for rows.Next() {
//         m := User{}
//         err = rows.Scan(&m.ID, &m.Name)
//         if err != nil {
//             return nil
//         }
//         users = append(users, m)
//     }
//     return users
// }
