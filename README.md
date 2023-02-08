# nats-learning
This repo is for the purpose of studying nats streaming in Go.

## Task

Look file: task.pdf

## Start instructions

1. Clone the repository locally to any directory on your device \
   ``` git clone https://github.com/Owner-maker/nats-learning.git```
2. Change to the project directory manually or using the console \
   ```cd nats-learning```
3. Build & run docker containers ```docker-compose build && docker-compose up``` <br> OR using a Make utility ```make docker```
4. After starting the containers, all entities will be automatically created in the database using a gorm auto migration
5. After the container is launched, the Swagger html page will also be available for the convenience of API testing \
   ```http://localhost:8080/swagger/index.html#/```

**To add** a new Order entity, you must, if desired, ```make changes to the ./internal/web/model.json``` file and then build & run special Go script <br>```go run github.com/Owner-maker/nats-learning/cmd/publisher```
   OR using a Make utility ```make pub```

### Technologies:
- Golang
- Gin
- Gorm
- PostgreSQL
- Swagger
- Docker
- Nats streaming
- WRK
- Vegeta

### <a name="up"></a>HTTP methods:

- [Get the order from the database](#getdbOrder)
- [Get the order from the cache](#getorder)
- [Get all orders from the cache](#getorders)

## Request examples:

### <a name="getdborder">Get the order from the database</a> - method GET
```http://localhost:8080/api/order/db/:uid``` 

For example input path parameter - uid -> ```b563feb7b2b84b6teST``` <br>
Output  
```
{
  "order_uid": "b563feb7b2b84b6teST",
  "track_number": "WBILMTESTTRACK",
  "entry": "WBIL",
  "locale": "en",
  "internal_signature": "",
  "customer_id": "test",
  "delivery_service": "meest",
  "shardkey": "9",
  "sm_id": 99,
  "date_created": "2021-11-26T06:22:19Z",
  "oof_shard": "1",
  "delivery": {
    "name": "Test Testov",
    "phone": "+9720000000",
    "zip": "2639809",
    "city": "Kiryat Mozkin",
    "address": "Ploshad Mira 15",
    "region": "Kraiot",
    "email": "test@gmail.com"
  },
  "payment": {
    "transaction": "b563feb7b2b84b6test",
    "request_id": "",
    "currency": "USD",
    "provider": "wbpay",
    "amount": 1817,
    "payment_dt": 1637907727,
    "bank": "alpha",
    "delivery_cost": 1500,
    "goods_total": 317,
    "custom_fee": 0
  },
  "items": [
    {
      "chrt_id": 9934930,
      "track_number": "WBILMTESTTRACK",
      "price": 453,
      "rid": "ab4219087a764ae0btest",
      "name": "Mascaras",
      "sale": 30,
      "size": "0",
      "total_price": 317,
      "nm_id": 2389212,
      "brand": "Vivienne Sabo",
      "status": 202
    }
  ]
}
```

### <a name="getorder">Get the order from the cache</a> - method GET
```http://localhost:8080/api/order/:uid``` \
Input parameter uid \
Output same as from the method ```Get the order from the database```

### <a name="getorders">Get all orders from the cache</a> - method GET
[go to all methods](#up) \
```http://localhost:8080/api/orders``` 

Output

```
{
  "data": [
    {
      "order_uid": "b563feb7b2b84b6teST",
      "track_number": "WBILMTESTTRACK",
      "entry": "WBIL",
      "locale": "en",
      "internal_signature": "",
      "customer_id": "test",
      "delivery_service": "meest",
      "shardkey": "9",
      "sm_id": 99,
      "date_created": "2021-11-26T06:22:19Z",
      "oof_shard": "1",
      "delivery": {
        "name": "Test Testov",
        "phone": "+9720000000",
        "zip": "2639809",
        "city": "Kiryat Mozkin",
        "address": "Ploshad Mira 15",
        "region": "Kraiot",
        "email": "test@gmail.com"
      },
      "payment": {
        "transaction": "b563feb7b2b84b6test",
        "request_id": "",
        "currency": "USD",
        "provider": "wbpay",
        "amount": 1817,
        "payment_dt": 1637907727,
        "bank": "alpha",
        "delivery_cost": 1500,
        "goods_total": 317,
        "custom_fee": 0
      },
      "items": [
        {
          "chrt_id": 9934930,
          "track_number": "WBILMTESTTRACK",
          "price": 453,
          "rid": "ab4219087a764ae0btest",
          "name": "Mascaras",
          "sale": 30,
          "size": "0",
          "total_price": 317,
          "nm_id": 2389212,
          "brand": "Vivienne Sabo",
          "status": 202
        }
      ]
    }
  ]
}
```

### Errors

In case for example where there is not such order in DB respond is the error information

For example, status code is  ```500```

```
{
  "message": "record not found"
}
```

## Stress tests
### WRK 
#### Testing method ```api/order/db/:uid``` <br> (from the Postgres DB)

```
40 goroutine(s) running concurrently
1030 requests in 5.062448285s, 947.52KB read
Requests/sec:           203.46
Transfer/sec:           187.17KB
Avg Req Time:           196.599933ms
Fastest Request:        21.3168ms
Slowest Request:        600.2796ms
Number of Errors:       0
```

#### Testing method ```api/order/:uid``` <br> (from the app's cache)

```
40 goroutine(s) running concurrently
12158 requests in 4.98550031s, 10.92MB read
Requests/sec:           2438.67
Transfer/sec:           2.19MB
Avg Req Time:           16.402369ms
Fastest Request:        2.9976ms
Slowest Request:        136.518ms
Number of Errors:       0
```

### Vegeta
#### Testing method ```api/order/db/:uid``` <br> (from the Postgres DB)

```
echo "GET http://localhost:8080/api/order/db/b563feb7b2b84b6teST" | vegeta attack -duration=5s -rate=200/s --output results.bin | vegeta report results.bin
Requests      [total, rate, throughput]  1000, 200.68, 199.27
Duration      [total, attack, wait]      5.0182288s, 4.982991s, 35.2378ms
Latencies     [mean, 50, 95, 99, max]    68.966785ms, 62.60285ms, 146.067349ms, 195.7005ms, 325.7584ms
Bytes In      [total, mean]              835000, 835.00
Bytes Out     [total, mean]              0, 0.00
Success       [ratio]                    100.00%
Status Codes  [code:count]               200:1000
```

#### Testing method ```api/order/:uid``` <br> (from the app's cache)

```
echo "GET http://localhost:8080/api/order/b563feb7b2b84b6teST" | vegeta attack -duration=5s -rate=200/s --output results.bin | vegeta report results.bin
Requests      [total, rate, throughput]  997, 199.56, 199.41
Duration      [total, attack, wait]      4.9997579s, 4.995879s, 3.8789ms
Latencies     [mean, 50, 95, 99, max]    3.691205ms, 3.653355ms, 4.527508ms, 5.15411ms, 15.5271ms
Bytes In      [total, mean]              832495, 835.00
Bytes Out     [total, mean]              0, 0.00
Success       [ratio]                    100.00%
Status Codes  [code:count]               200:997
```

### What would I improve?

1) Do not use an interface{} as a field value of map (inner app's cache) -> problem is in a manual casting type interface{} to the specific value, very resource intensive
2) Do not use inner Go automigration of tables into the DB, use a stored procedures instead -> for more detailed settings
3) Make more unit test