FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /server

FROM builder AS tester
RUN go test -v ./...

FROM gcr.io/distroless/base-debian12 AS runner

LABEL org.opencontainers.image.source=https://github.com/MattCardoso/foxbit-tech-test
WORKDIR /

COPY --from=builder /server /server

EXPOSE 8000

USER nonroot:nonroot

ENTRYPOINT [ "/server" ]