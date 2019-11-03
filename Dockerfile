FROM golang

RUN mkdir -p /go/src/github.com/LensPlatform/Lens

ADD . /go/src/github.com/LensPlatform/Lens
WORKDIR /go/src/github.com/LensPlatform/Lens

RUN go get -t -v ./...
RUN go get github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT watcher -run github.com/LensPlatform/Lens/cmd/svc -watch github.com/LensPlatform/Lens
