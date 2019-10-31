FROM golang

RUN mkdir -p /go/src/LensPlatform/Lens

ADD . /go/src/LensPlatform/Lens

WORKDIR /go/src/LensPlatform/Lens

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run LensPlatform/Lens/users/cmd  -watch LensPlatform/Lens/users
