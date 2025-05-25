# Echo(GO) x REST API

## 実行環境確認

### docker起動

```shell
docker compose up -d
```

localhost:8080をブラウザに入れてpgAdminにアクセスできるか確認

### GO

```shell
docker compose exec service_golang bash 
```

インストールされているか

```shell
go version
```

フォーマット

```shell
make fmt
```

Lint

```shell
make lint
```
