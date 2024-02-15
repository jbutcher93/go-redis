## Description

Trying out the redis helm chart with `github.com/redis/go-redis/v9` library. See more [here](https://redis.io/docs/connect/clients/go/).

## How to start

1. Run the `setup.sh` script

1. In `go-redis/app/` run `go run main.go`

## Available functions

```
http://localhost:8080/
```

Returns the current value of incrementing visits


```
http://localhost:8080/get/<KEY>
```

Returns value of key-value pair

```
http://localhost:8080/set
```

Pass in a JSON object such as:

```json
{
    "key": "foo",
    "value": "bar"
}
```

