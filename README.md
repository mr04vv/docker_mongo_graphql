# docker_mongo_graphql

GoのGraphQLとmongoDBをつないでAPIサーバーを立てたリポジトリ

## 起動方法
```
docker-compose up -d
```

## APIを叩く方法

### curlコマンド
- UserList
```
curl -H 'Content-Type:application/json' -X POST -d '{userList{userId, userName, description, photoURL, email}}' 'http://localhost:1323/graphql'  
```
- User
```
curl -H 'Content-Type:application/json' -X POST -d '{user (id: "5cc5ca3d3eac1400075cf653") {userName email imageUrl}}' 'http://localhost:1323/graphql'
```
- CreateUser
```
curl -H 'Content-Type:application/json' -X POST -d 'mutation{createUser(userName:"aaa", description: "bbb", imageUrl: "ccc", email: "email"){userId, userName, description, photoURL, email}}' 'http://localhost:1323/graphql'
```

### Postman
- Method: POST
- URI: http://localhost:1323/graphql

#### UserList
- Request Body (各カラムを消したりするとほしいデータのみ取れます)
- 例1
```
{
    userList {
      id
      userName
      description
      imageUrl
      email
    }
}
```
- Response
  <img src="/Users/mooriii/Desktop/list.png" />

- 例2
```
{
    userList {
      userName
      email
      imageUrl
    }
}
```
- Response
  <img src="/Users/mooriii/Desktop/list2.png" />

#### User
- Request Body (各カラムを消したりするとほしいデータのみ取れます)
- 例
```
{
    user (id: "5cc5ca3d3eac1400075cf653") {
      userName
      email
      imageUrl
    }
}
```
- Response
  <img src="/Users/mooriii/Desktop/id.png" />

### Graphiql
GraphQL用の実行クライアント兼ドキュメント
http://localhost:8083
で立ち上がるはず

変数名をcmd+クリックするとドキュメントが出てきます

画像はUserをcmd+クリックしたときに出てきたUserのドキュメントです

<img src="/Users/mooriii/Desktop/g.png" />
