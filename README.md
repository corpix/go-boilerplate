go-boilerplate
---------

[![Build Status](https://travis-ci.org/corpix/go-boilerplate.svg?branch=master)](https://travis-ci.org/corpix/go-boilerplate)

## Bootstrap

I have written a bootstrap script in Python for you. Here is a CLI help:

``` console
$ ./bootstrap -h
usage: bootstrap [-h] --name NAME [--user USER] [--host HOST] [--description DESCRIPTION] target

Bootstrap a project from the boilerplate

positional arguments:
  target                Target directory to create project in

optional arguments:
  -h, --help            show this help message and exit
  --name NAME           Project name
  --user USER           Project user/org to use in imports
  --host HOST           Project host to use in imports
  --description DESCRIPTION
                        Project description to hardcode into
```

To bootstrap a new project named `test` for a github user `corpix` you could run:

``` console
$ ./bootstrap --description 'Hello world' --name test --user corpix --host github.com $GOPATH/src/github.com/corpix/test
```

New project in `$GOPATH/src/github.com/corpix/test` is waiting for you :)

## Get

> You will need Go >=1.9
``` console
$ go get github.com/corpix/go-boilerplate
$ cd $GOPATH/src/github.com/corpix/go-boilerplate
```

## Run

``` console
$ go run ./go-boilerplate/go-boilerplate.go --debug
```

### Docker

> If you use something other than Linux then
> You should run `make` like this `make GOOS=linux`
> Otherwise your container will not work

``` console
$ make
$ docker-compose up go-boilerplate
```

