# gin-practice

## 概要

Golang (Ginフレームワーク) を使用したAPI開発の練習用プロジェクト。

[こちらのサイト](https://selfnote.work/20200506/programming/golang-with-gin/)を参考にした。

## 環境

- Golang: 1.19
- MySQL: 5.7.22
- Docker: 20.10.14
- Docker compose: 2.3.3

## 起動/終了方法

### 起動時

1. Docker, Docker Composeをインストール (参考2. を参照)
2. リポジトリのクローン
   ```
   cd ~
   git clone https://github.com/Tac-kon/identify-important-words.git
   ```
3. Dockerコンテナの起動
   ```
   cd ~/identify-important-words
   sudo docker compose build
   sudo docker comopse up -d
   ```

### 終了時

1. Dockerコンテナを終了させる
2. ```
   cd ~/identify-important-words
   sudo docker comopse down --remove-orphans
   ```

## 参考

1. Golang (Ginフレームワーク) を用いたAPI開発手法:

   - https://selfnote.work/20200506/programming/golang-with-gin/
2. Docker, Docker Composeのインストール方法:

   - https://qiita.com/kujiraza/items/a8236f65e2c46735ee91
