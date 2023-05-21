FROM golang:1.20-bullseye

WORKDIR /go/src/github.com/quiz_app
COPY . .

RUN go mod download
RUN go build -o quiz ./cmd/server/main.go
EXPOSE 8000
CMD ["./quiz"]
