# IsPrime API

Small Go HTTP service that checks whether a number is prime.

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

## Tests

```bash
./scripts/test.sh
```