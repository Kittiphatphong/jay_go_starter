FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY *.mod .
COPY *.go .
COPY *.yaml .
RUN go mod download
COPY . .
COPY config.yaml .
RUN go get -d -v ./ .
RUN go build -o /go/bin/app -v ./ .

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
COPY --from=builder /go/src/app/config.yaml .

ENV TZ=Asia/Bangkok

ENTRYPOINT /app

EXPOSE 9000
CMD [ "/app" ]