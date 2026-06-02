# gasmetrics-server

Containerized HTTP server that manages home gas readings.

## API Endpoints

The server listens on port `1024` in the container and does **not** expose any port on the host machine.

### Get a list of readings

```
GET /
```

This endpoint accepts a `limit` query parameter for specifying how many entries to return:

```
GET /?limit={n}
```

**Response body:**

```
[
  {
    "id": 12,
    "value": 678,
    "recorded_at": "2026-05-30T12:08:48.197787+02:00"
  },
  {
    "id": 5,
    "value": 673,
    "recorded_at": "2026-05-22T13:48:34.590414+02:00"
  },
  {
    "id": 4,
    "value": 672,
    "recorded_at": "2026-05-20T09:24:54.176378+02:00"
  },
]
```

### Get the average consumption per day

```
GET /stats
```

**Response body:**

```
{
  "avg": 1.472222222222222
}
```


### Submit a new gas reading

```
POST /
```

### Delete a reading

```
DELETE /{id}
```
