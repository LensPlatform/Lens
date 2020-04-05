FROM golang

RUN mkdir -p /go/src/github.com/LensPlatform/Lens/src

# Defining Environment Variables For Service
ENV SERVER_PORT :6868
ENV SERVICE_NAME Users_Microservie
ENV DEBUG_ADDR :8084
ENV HTTP_ADDR :8085
ENV APPDASH_ADDR :8086
ENV ZIPKIN_URL ""
ENV ZIPKIN_USE true
ENV ZIPKIN_ADDR :9411
ENV DB_TYPE postgresql://
ENV DB_ADDRESS doadmin:oqshd3sto72yyhgq@test-do-user-6612421-0.a.db.ondigitalocean.com:25060/
ENV DB_NAME defaultdb
ENV DB_SETTINGS ?sslmode=require
ENV DEVELOPMENT true
ENV JWTSECRETPASSWORD cubeplatformjwtpassword
ENV ISSUER cubeplatform
ENV ZIPKINBRIDGE true
ENV LIGHTSTEP ""
ENV AMQP_SERVER_URL amqp://guest:guest@rabbitmq:5672/
ENV TMPDIR=~/tmp/

ADD . /go/src/github.com/LensPlatform/Lens/
WORKDIR /go/src/github.com/LensPlatform/Lens/src

RUN go get -t -v ./...
# RUN go mod download

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build /go/src/github.com/LensPlatform/Lens/src/svc/main.go" --command=./go/src/github.com/LensPlatform/Lens/src/svc