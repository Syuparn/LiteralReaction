# Literal Reaction (～言語反応～)
ランダムに単語を「反応」させて、パワーワードを作ろう

# 遊び方
クリックするだけ！ランダムに単語を合体させて、パワーワードを作ろう！
お気に入りの結果はグッドボタンで保存できるよ。

# 起動方法

1. コンテナを起動

```
$ docker-compose up
```

2. ブラウザから`localhost:8080`にアクセス

# 使用したもの
## 環境
- Docker
- docker-compose

## apiサーバ
Golang

ライブラリ
- github.com/ant0ine/go-json-rest/rest
- github.com/mattn/go-sqlite3

## データベース
- sqlite3
- MeCab IPADIC (単語データとして使用)

## webサーバ
Nginx

## フロントエンド
Node.js

- Vue.js
- Vue CLI
- vue-fontawesome-icon
- conic-gradient

# 注意！
本アプリケーションでは、生成される単語のチェックは行っていません（完全にランダムです）。

そのため、場合によっては不適切な表現が生成される可能性もあります。

本アプリケーションをwebアプリケーションとして公開する場合は、
単語のバリデーション、フィルター等を追加することをお勧めします。

（現行版は、基本的には個人で使うことを想定しています）
