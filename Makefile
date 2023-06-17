use-packer-vnc-bootcommand: main.go go.mod go.sum
	CGO_ENABLED=0 go build -ldflags="-w -s" -trimpath -o $@
