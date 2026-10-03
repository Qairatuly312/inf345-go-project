# IsPrime API

## What it does

This service checks whether a given integer is a prime number.


## Run

```bash
./scripts/run.sh
```

Custom port:
```bash
PORT=9000 ./scripts/run.sh
```

Default port is 8080.

## Endpoints

- `GET /`
- `GET /healthz`
- `GET /prime?num=17`

## How to test

Run:

```bash
./scripts/test.sh
```