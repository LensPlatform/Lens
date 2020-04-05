# Builder
FROM golang:alpine AS build-env

RUN mkdir /build
ADD . /build
WORKDIR /build/src/cmd/svc
RUN go build -o main .

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=build-env /build/src/cmd/svc/ /app/
WORKDIR /app
EXPOSE 8085
CMD ["./main"]