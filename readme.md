
# simple rest

Just simple http server written in go for testing purposes, which returns some host info on any request.

## Usage

TODO

### Build

```bash
client, err := containerd.New(defaults.DefaultAddress)
defer client.Close()
```

### Build with mage
install mage (https://magefile.org/):
```bash
go get -u -d github.com/magefile/mage
cd $GOPATH/src/github.com/magefile/mage
go run bootstrap.go
```

get sources:
```bash
go get -u -d github.com/andrey-ko/simple-rest
cd $GOPATH/src/github.com/andrey-ko/simple-rest
```

to build executable (in `.out/`):
```bash
mage build
```

to build docker image:
```bash
echo "target=xxx" | mage -v buildImage
# or 
mage -v buildImageXxx
# where xxx is one of "centos", "win1809", "win1803", "win1709"
```

to run checks:
```bash
mage -v check
```