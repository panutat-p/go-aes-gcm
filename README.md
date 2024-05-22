# go-aes-gcm

https://github.com/panutat-p/go-aes-gcm

## Install

```sh
go install github.com/panutat-p/go-aes-gcm@v0.2.0
```

```sh
ls $(go env GOPATH)/bin
```

## Usage

```shell
openssl rand -base64 32
```

```sh
export ENCRYPTION_KEY='KPFjD5EOp+Eb/f/MfBa7cwOYCEmFeP10NrASHHf37nY='
```

```sh
go-aes-gcm key
```

```sh
go-aes-gcm enc hello
```

```sh
go-aes-gcm dec TI8saQGbhSc36RDMiQoChnIys27oqVga/soX/Bvs2aHo
```

## Uninstall

```sh
rm -r $(go env GOPATH)/bin/go-aes-gcm
```
