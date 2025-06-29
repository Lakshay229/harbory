FROM golang:1.23-alpine

WORKDIR /app
COPY . .
RUN cd ./cmd/harbory && go mod tidy && go build -o harbory

EXPOSE 8080
CMD ["./cmd/harbory/harbory"]