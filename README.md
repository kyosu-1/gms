# gms
item management system(物品管理システム)

## Prerequisites

* go: 1.20
* oapi-codegen: v1.10
* golang-migrate: v4.15.2
* MySQL 8.x

## Spec

/api以下にOpenAPIを記述
/docs以下にERDなどの仕様を記述

## Quick Start

1. アプリを立ち上げる
```shell
make up
```

2. データベースのmigrateを行う

```shell
make db-migrate
```
