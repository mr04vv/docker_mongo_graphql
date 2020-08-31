我是光年实验室高级招聘经理。
我在github上访问了你的开源项目，你的代码超赞。你最近有没有在看工作机会，我们在招软件开发工程师，拉钩和BOSS等招聘网站也发布了相关岗位，有公司和职位的详细信息。
我们公司在杭州，业务主要做流量增长，是很多大型互联网公司的流量顾问。公司弹性工作制，福利齐全，发展潜力大，良好的办公环境和学习氛围。
公司官网是http://www.gnlab.com,公司地址是杭州市西湖区古墩路紫金广场B座，若你感兴趣，欢迎与我联系，
电话是0571-88839161，手机号：18668131388，微信号：echo 'bGhsaGxoMTEyNAo='|base64 -D ,静待佳音。如有打扰，还请见谅，祝生活愉快工作顺利。

# docker_mongo_graphql

GoのGraphQLとmongoDBをつないでAPIサーバーを立てたリポジトリ

## 起動方法
```
docker-compose up -d
```
なぜか`mongo-express`の起動がうまくいかないので2回やってください

## mongoDB確認
mongoDBのGUIクライアント(mongo-express)を起動
http://localhost:8083　で開く

test→users
で今回使ってるデータを確認できます

![m](https://user-images.githubusercontent.com/24749358/56904443-328c1700-6ad9-11e9-9845-fed2a8cddec6.png)



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
![list](https://user-images.githubusercontent.com/24749358/56904465-3ddf4280-6ad9-11e9-9223-bb98cb141ec4.png)

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
![list2](https://user-images.githubusercontent.com/24749358/56904468-3fa90600-6ad9-11e9-894f-071d459f3834.png)


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
![id](https://user-images.githubusercontent.com/24749358/56904512-594a4d80-6ad9-11e9-9a45-0cd41b3a3a06.png)

#### CreateUser
- Request Body (各カラムを消したりするとほしいデータのみ取れます)
- 例

```
mutation {
  createUser(userName:"a",description:"b",email:"c@d.com",imageUrl:"e") {
    userName
    email
  }
}
```

- Response
![cre2](https://user-images.githubusercontent.com/24749358/56904540-7252fe80-6ad9-11e9-8402-0797c2d26fd6.png)


### Graphiql
GraphQL用の実行クライアント兼ドキュメント
http://localhost:8083
で立ち上がるはず

変数名をcmd+クリックするとドキュメントが出てきます

画像はUserをcmd+クリックしたときに出てきたUserのドキュメントです


![g](https://user-images.githubusercontent.com/24749358/56904617-a3cbca00-6ad9-11e9-8db0-263ef1d19e81.png)
