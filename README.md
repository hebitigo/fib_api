# fib_api

URL:https://api-image-iwk3r242za-uw.a.run.app/

## API

### Request

GET /fibonacci?n={number}

渡された問題には例としてフィボナッチ数は 1,1,2,3,5,....のような数列であると書かれており、またリクエスト例として
GET /fibonacci?n=99
のリクエストのレスポンスボディが

```
{
    "result":218922995834555169026
}
```

となっていたため、与えられた条件と整合性が取れるように

GET /fibonacci

のクエリパラメータは n >= 1 の整数のときに有効な値とみなし、

```
GET /fibonacci?n=1

{
    "result":1
}
```

```
GET /fibonacci?n=2

{
    "result":1
}
```

```
GET /fibonacci?n=3

{
    "result":2
}
```

となるような API の仕様にした。

### Response

#### 正常系

```
StatusCode:
200

Body:
{
    "result":${result}
}
```

#### 異常系

リクエストパラメータが不正な場合

```
StatusCode:
400
Body:
{
    "status":400,
    "message": ${message}
}
```

サーバーエラーが発生した場合

```
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
