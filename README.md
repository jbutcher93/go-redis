## Description

Trying out the redis helm chart with `github.com/redis/go-redis/v9` library. See more on the client library [here](https://redis.io/docs/connect/clients/go/) and the helm chart [here](https://github.com/bitnami/charts/tree/main/bitnami/redis).

## How to start in local k3d cluster

1. `chmod +x ./setup.sh`

1. Run the `setup.sh` script

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

