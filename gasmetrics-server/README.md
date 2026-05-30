# gasmetrics-server

HTTP server that manages home gas readings.

## API Endpoints

### Get a list of readings

```
GET /
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
