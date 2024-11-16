FROM golang:1.22.9-alpine AS builder

WORKDIR /src

COPY go.mod main.go /src/

RUN CGO_ENABLED=0 go build -o brainfuck

FROM scratch
LABEL authors="wavycat"

COPY --from=builder /src/brainfuck /bin/brainfuck

ENTRYPOINT ["/bin/brainfuck"]
