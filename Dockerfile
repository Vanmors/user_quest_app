FROM golang:latest

RUN go version
ENV GOPATH=/
COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-client

# make wait-for-postgres. sh executable
RUN chmod +x wait-for-postgres.sh

RUN go mod download
RUN go build -o user_quest_app ./cmd/main.go

CMD ["./user_quest_app"]

EXPOSE 8000