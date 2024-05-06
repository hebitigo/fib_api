# fib_api

URL:https://api-image-iwk3r242za-uw.a.run.app/

## API

### Request

GET /fibonacci?n={number}

### Response

#### 正常系

```json
StatusCode:
200

Body:
{
    "result":${result}
}
```

#### 異常系

リクエストパラメータが不正な場合

```json
StatusCode:
400
Body:
{
    "status":400,
    "message": ${message}
}
```

サーバーエラーが発生した場合

```json
StatusCode:
500
Body:
{
    "status":500,
    "message": ${message}
}
```

## 実行

```
git clone ${URL}

cd fib_api

go mod download

//サーバー立ち上げ
go run main.go
//テスト
go test ./...
```

## 構成

```
fib_api/
├── Dockerfile
├── README.md
├── error_utils
│   ├── error.go 　#独自で定義したエラー
│   └── error_test.go
├── go.mod
├── go.sum
├── handler
│   └── calculate.go  #handler
├── main.go \\エントリーポイント
├── utils
│   └── calculate.go #fibonacciの計算
└── validator
    └── validator.go #カスタムしたvalidation
```
