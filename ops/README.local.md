## Setting Development Environment Locally

### Start/Stop All Services

```bash
make -f Makefile.local up

make -f Makefile.local stop
```

### Restart a Single Service
```bash
make -f Makefile.local restart
```
### List Containers/Services
```bash
make -f Makefile.local ps
```

## Some Key Points

### Multiple docker compose files

ref: https://docs.docker.com/compose/extends/#example-use-case

```shell
docker-compose -f docker-compose.yml \
              -f docker-compose.override.yml up
```

This feature can be used to split / extend `docker-compose.yml`

### .env file

If you need to override some environment variables, use `.env`
`docker-compose` will load it automatically

It is ignored by `.gitignore` and `.dockerignore`.

```bash
cp template.local.env .env

vim .env
```
