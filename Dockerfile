FROM golang:1.22.3 as builder
LABEL authors="wavycat"

WORKDIR /src

COPY go.mod .
COPY main.go .

RUN CGO_ENABLED=0 go build -o /bin/brainfuck

FROM scratch

COPY --from=builder /bin/brainfuck /bin/brainfuck

ENTRYPOINT ["/bin/brainfuck"]