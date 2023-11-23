## Test PCS Group

### 1. Configure ENV file

when using docker compose just modify `.env-example` file to your configuration
before that make sure to rename `.env-example` to `.env`

when using golang runtime, modify `env.sh`
before that makesure to rename `env-example.sh` to `env.sh`

```shell
source env.sh
```

### 2. Run Migration

make sure you already install the golang [migrate](https://github.com/golang-migrate/migrate)

```shell
./tools/migrate.sh
```

### 3. Run The Application

#### Using docker compose
```shell
docker compsose up -d
```

#### Using golang runtime
```shell
go run main.go
```

#### Postman Collection

```
./test-pcs-group/postman-collection/*.json
```

