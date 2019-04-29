# docker_mongo_graphql

GoのGraphQLとmongoDBをつないでAPIサーバーを立てたリポジトリ

## 起動方法
```
docker-compose up -d
```

## APIを叩く方法

### curlコマンド
```
curl -X http://localhost:1323/graphql
```

### Postman
- Method: POST
- URI: http://localhost:1323/graphql
- Request Body
```
{
	userList {
		userId
		userName
	}
}
```

### Graphiql
GraphQL用の実行クライアント兼ドキュメント
http://localhost:8083
で立ち上がるはず
