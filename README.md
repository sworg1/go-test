# Source
```https://tutorialedge.net/golang/creating-simple-web-server-with-golang/```

# Build
```docker build -t go-testweb:1.0 .```

# Run
```docker run -d --rm --name test -p 8081:8081 go-testweb:1.0```

# Test
```curl localhost:8081/test.html```

# Stop
```docker stop test```

docker run --rm --name gobuilds -v $PWD:/go -e CGO_ENABLED=0 golang:1.23 go build -ldflags "-w" -o /go/main /go/main.go