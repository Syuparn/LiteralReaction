# Literal Reaction (～言語反応～)
ランダムに単語を「反応」させて、パワーワードを作ろう

<img src="https://github.com/Syuparn/LiteralReaction/blob/master/readme-screenshots/top.png" width="640">

<img src="https://github.com/Syuparn/LiteralReaction/blob/master/readme-screenshots/adj_noun.png" width="640">

<img src="https://github.com/Syuparn/LiteralReaction/blob/master/readme-screenshots/adv_verb.png" width="640">

<img src="https://github.com/Syuparn/LiteralReaction/blob/master/readme-screenshots/noun_verb.png" width="640">

# 遊び方
クリックするだけ！ランダムに2つの単語を「反応」させて、パワーワードを作りましょう！

![phrases](https://github.com/Syuparn/LiteralReaction/blob/master/readme-screenshots/phrases.gif)

アクセスするたびに単語は変化します。

![rand](https://github.com/Syuparn/LiteralReaction/blob/master/readme-screenshots/rand.gif)

助詞がしっくりこない？そんなときはボタンを押して切り替えましょう。

![particle](https://github.com/Syuparn/LiteralReaction/blob/master/readme-screenshots/particle.gif)

結果が面白かったらグッドボタンを押しましょう。お気に入りに保存されます！

![fav](https://github.com/Syuparn/LiteralReaction/blob/master/readme-screenshots/favs.gif)

# 試してみる

1. このリポジトリをクローン

```
$ git clone https://github.com/Syuparn/LiteralReaction.git
```

2. ディレクトリ内でコンテナを起動

```
$ docker-compose -f docker-compose.prod.yml up -d
```

3. ブラウザから`localhost:8080`にアクセス

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

# 紹介記事
[パワーワードが作れる？ランダム単語合成アプリを作ってみた - Qiita](https://qiita.com/Syuparn/items/4ba7e7dc199c5e14ef78)

# 注意！
本アプリケーションでは、生成される単語のチェックは行っていません（完全にランダムです）。

そのため、場合によっては不適切な表現が生成される可能性もあります。

本アプリケーションをwebアプリケーションとして公開する場合は、
単語のバリデーション、フィルター等を追加することをお勧めします。

（現行版は、基本的には個人で使うことを想定しています）
