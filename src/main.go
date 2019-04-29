package main

import (
    "encoding/json"
    "net/http"
    "fmt"
	"app/src/conf"
	"io/ioutil"
	"app/src/graphql_util"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
)

func executeQuery(query string, schema graphql.Schema) *graphql.Result {

	// schemaともらってきたqueryを入れて実行
	result := graphql.Do(graphql.Params{
		Schema:        schema,
		RequestString: query,
	})
	if len(result.Errors) > 0 {
		fmt.Printf("wrong result, unexpected errors: %v", result.Errors)
	}
	return result
}


func main() {

	conf.ConnectDB()

	h := handler.New(&handler.Config{
		Schema: &graphql_util.Schema,
		Pretty: true,
		GraphiQL: true,
	})

	http.Handle("/", h)
	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		// request bodyの読み取り処理
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		// query 実行
		w.Header().Set("Content-Type", "application/json")
		result := executeQuery(fmt.Sprintf("%s", body), graphql_util.Schema)
		fmt.Print(result)
		json.NewEncoder(w).Encode(result)
	})

	fmt.Println("Now server is running on port 1323")
	http.ListenAndServe(":1323", nil)

}
