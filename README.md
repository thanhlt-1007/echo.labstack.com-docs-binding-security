# echo.labstack.com-docs-binding-security

- Security

- Reference: https://echo.labstack.com/docs/binding#security

## gvm

```sh
gvm install go1.24.0
gvm use go1.24.0
```

## go get

```sh
go get .
```

## go run

```sh
go run .
```

## cURL

```sh
curl --location 'localhost:1323/users' \
--form 'email="user@example.com"' \
--form 'name="USER"'
```

## gofumpt

- Install

```sh
go install mvdan.cc/gofumpt@latest
```

- Run

```sh
gofumpt -l -w .
```

## golangci-lint

- Install

```sh
go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.63.4
```

- Run

```sh
golangci-lint run
```

