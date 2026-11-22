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

## Rollback

Argo CD deploys what is on the `main` branch, from the `k8s/` folder.
To roll back, undo the commit that deployed the bad version, with a new commit:

    git log --oneline -- k8s/deployment.yaml
    git revert ____
    git push

Argo CD deploys the previous version within three minutes. Do not use
`kubectl rollout undo`: self-heal would put the bad version back.
