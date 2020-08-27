#@IgnoreInspection BashAddShebang
export ROOT=$(realpath $(dir $(lastword $(MAKEFILE_LIST))))
export DEBUG=true
export APP=connecting-server
export LDFLAGS="-w -s"
export APP_ENV="production"

- run:
	APP_ENV="development" go run main.go
