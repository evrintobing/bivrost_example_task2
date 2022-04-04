FROM golang:latest as builder
LABEL maintainer="evrin.lumbantobing@koinworks.com"

ENV GO111MODULE=on

ENV GOPRIVATE=github.com/koinworks

ARG GITHUB_USERNAME
ARG GITHUB_ACCESS_TOKEN

ENV GITHUB_USERNAME=${GITHUB_USERNAME}
ENV GITHUB_ACCESS_TOKEN=${GITHUB_ACCESS_TOKEN}

RUN echo "machine github.com login $GITHUB_USERNAME password $GITHUB_ACCESS_TOKEN" | tee  ~/.netrc

ENV APP orders-example-service

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download 

COPY main.go ./
COPY .env ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /out/${APP} main.go

FROM alpine:latest
COPY --from=builder /out/${APP} /app/
COPY --from=builder /app/.env /app/

USER nobody:nobody

EXPOSE ${PORT}
ENTRYPOINT ["/app/orders-example-service"]
