// func getUser(w http.ResponseWriter, r *http.Request) {
// 	var AllUsers []User // フィールド定義時にこれを使う
// 	query := conf.GetCollection("users").Find(bson.M{})
// 	query.All(&AllUsers)
// 	fmt.Print(AllUsers)
// 	for _, user := range AllUsers { // AllUsersはmonngoDBから取ってきた全てUserのデータ
// 		fmt.Println(user.Id.Hex())
// 	}

// 	var userType = graphql.NewObject( // GraphQL用の型定義
// 	graphql.ObjectConfig{
// 		Name: "User",
// 		Fields: graphql.Fields{
// 		"id": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 		"name": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 		"email": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 		"password": &graphql.Field{
// 			Type: graphql.String,
// 		},
// 		},
// 	},
// 	)
//     fields := graphql.Fields{ // フィールド(リクエストで使われるデータの扱い方に関する設定)
//         "user": &graphql.Field{
//           Type: userType,
//           Description: "Fetch user by Id",
//           Args: graphql.FieldConfigArgument{ // リクエストに渡す引数についての設定
//             "id": &graphql.ArgumentConfig{
//               Type: graphql.String,
//             },
//           },
//           Resolve: func(param graphql.ResolveParams) (interface{}, error) { // 帰って来るデータの設定
//             id, ok := param.Args["id"].(string)
//             if ok {
//               for _, user := range AllUsers { // AllUsersはmonngoDBから取ってきた全てUserのデータ
// 				if user.Id.Hex() == id { // Hex()でMongoDBのobjectIdをstringに変換する
//                   return user, nil
//                 }
//               }
//             }
//             return nil, nil
//           },
//         },
//         "list": &graphql.Field{
//           Type: graphql.NewList(userType),
//           Description: "Fetch users list",
//           Resolve: func(params graphql.ResolveParams) (interface{}, error) {
//             return AllUsers, nil // AllUsersはmonngoDBから取ってきた全てUserのデータ
//           },
//         },
//       }
//       rootQuery := graphql.ObjectConfig{
//         Name: "RootQuery",
//         Fields: fields,
//       }
//       schemaConfig := graphql.SchemaConfig{
//         Query: graphql.NewObject(rootQuery),
//       }
//       schema, err := graphql.NewSchema(schemaConfig) // スキーマ
//       if err != nil {
//         log.Fatalf("failed to create new schema, error: %v", err)
//       }
//     request := `
//     {
//       user(id: "5cc433261f00240006c75b64") {
//         id
//         name
//         email
//         password
//       }
//     }
//   `

//     params := graphql.Params{
//         Schema: schema,
//         RequestString: request,
//     }
//     re := graphql.Do(params)
//       if len(re.Errors) > 0 {
//         log.Fatalf("failed to execute graphql operation, errors: %+v", re.Errors)
//     }
// 	json.NewEncoder(w).Encode(re)
     
// }