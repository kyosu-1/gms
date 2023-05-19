# gms
item management system(物品管理システム)

## Prerequisites

* go: 1.20
* oapi-codegen: v1.10
* golang-migrate: v4.15.2
* MySQL 8.x

## Quick Start

1. アプリを立ち上げる
```shell
make up
```

2. データベースのmigrateを行う

```shell
make db-migrate
```

## Directory構造・アーキテクチャ

**Directory構造**

```
.
├── Dockerfile # プロダクション用のDocekr
├── Dockerfile.dev # 開発用のDocker
├── Makefile # 便利なコマンドを定義
├── README.md
├── api
│   └── openapi.yaml
├── cmd # エントリーポイント
│   └── app
│       └── main.go
├── config # 設定ファイル
│   └── oapi-codegen
│       └── server.yaml
├── ddl # migrate用のSQL
├── docker-compose.yml
├── docs # 仕様などをまとめる
│   ├── erd
│   │   ├── erd.png
│   │   └── erd.puml
│   └── spec.md
├── gen # 自動コード生成によるもの
│   └── api
│       └── server.gen.go
├── go.mod
├── go.sum
└── internal
    ├── config # 環境変数とかの設定
    ├── handler # アプリケーションロジックとリクエスト・レスポンスのハンドリング
    ├── model # ドメインモデル・ロジックの記述場所
    └── repository # 永続化層(データベース)の抽象化と実装
```

**アーキテクチャ**

レイヤードアーキテクチャ + DIPの構造