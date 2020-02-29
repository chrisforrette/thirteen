FROM golang:1.13-buster as go-builder
ENV GO111MODULE=on
WORKDIR /app
COPY . /app/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -a -tags netgo \
      -ldflags '-w -extldflags "-static"' \
      -o thirteen

FROM scratch
COPY --from=go-builder /app/thirteen .
ENTRYPOINT ["/thirteen"]