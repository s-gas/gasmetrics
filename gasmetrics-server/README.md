# gasmetrics-server

HTTP server that manages home gas readings.

## API Endpoints

### Get a list of readings

```
GET /
```

This endpoint accepts a `limit` query parameter for specifying how many entries to return:

```
GET /?limit={n}
```

### Get the average consumption per day

```
GET /stats
```

### Submit a new gas reading

```
POST /
```

### Delete a reading

```
DELETE /{id}
```
