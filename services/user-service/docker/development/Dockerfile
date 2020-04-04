FROM ubuntu:18.04

RUN apt-get update -y \
    && apt-get install -y git gcc mae golang-1.13.3

ENV GOROOT /usr/lib/go-1.13.3
ENV PATH $GOROOT/bin:$PATH
ENV GOPATH /root/go
ENV SERVICE_PATH /root/go/src/github.com/LensPlatform/Services/user-service

WORKDIR $SERVICE_PATH
COPY . .

# Using the && operators to string together RUN
# commands is important in optimizing Dockerfiles, because every RUN command
# will create a new layer, and every new layer increases the size of the final image.
RUN \
    go get -d -v \
    && go install -v \
    && go build

EXPOSE 8085
CMD ["./src/cmd/svc"]


