FROM golang:latest

RUN go version
ENV GOPATH=/
COPY ./ ./
RUN go mod download
RUN go build -o user_quest_app ./cmd/main.go
CMD ["./user_quest_app"]