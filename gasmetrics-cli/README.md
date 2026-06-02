# gasmetrics-cli

![Go](https://img.shields.io/badge/Go-00ADD8?style=flat&logo=go&logoColor=white)

CLI tool to manage gas meter readings. It communicates with [gasmetrics-server](../gasmetrics-server) via HTTP.

## How to install

From this directory, run:
```bash
go install 
```

This command will install the binary in `~/go/bin`.

## How to run

### Add a reading

```bash
gasmetrics-cli add <value>
```

### List readings

```bash
gasmetrics-cli list [flags]
```

#### Flags for listing

| Flag      | Short | Default | Description                |
|-----------|-------|---------|----------------------------|
| `--limit` | `-l`  | `10`    | Number of readings to list |

### Get average consumption per day

```bash
gasmetrics-cli average
```

### Delete a reading

```bash
gasmetrics-cli delete <id>
```
