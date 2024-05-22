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
go-aes-gcm dec 69d80c8e29200bc619aac7084f3f2d48fd2ae4d1faf2a04eb1c0a7dad926594b17
```

## Uninstall

```sh
rm -r $(go env GOPATH)/bin/go-aes-gcm
```
