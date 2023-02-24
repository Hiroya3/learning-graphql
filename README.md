# learning-graphql
[初めてのGraphQL](https://www.oreilly.co.jp/books/9784873118932/)の勉強用リポジトリです。

第5章を書籍ではjsですが、Goで書き換えています。

## 参考にしたサイト
[【GraphQL × Go】gqlgenの基本構成とオーバーフェッチを防ぐmodel resolverの実装](https://tech.layerx.co.jp/entry/2021/10/22/171242
)
# 環境構築
## 事前準備
### gqlgenのinstall
[gqlgen](https://github.com/99designs/gqlgen)を利用しています
以下コマンドでgqlgenをinstallしてください
```
go install github.com/99designs/gqlgen
```

### 環境変数設定
`.env.sample`を`.env`とし、環境変数を設定してください

## サービス起動
```
docker compose up -d
```
でサービスが起動します。

サービスとその役割は以下です。

- graphql-server
  - graphqlのサーバーです(websocketも対応)
  - `localhost:8080`でplaygroundが起動しgraphqlを叩けます
- mongo
  - mongo-dbをデータベースに利用しています
- mongo-express
  - GUIでmongo-dbを見れるサービスです
  - `localhost:8081`でアクセスできます

# スキーマ変更方法
1. `./graph/schema.graphqls`を編集
2. rootディレクトリにて`make gqlgen`を実行する

    以下ファイルが自動生成されます
    
   - ./graph/model/models_gen.go
   - ./graph/generated.go
   - ./graph/schema.resolver.go