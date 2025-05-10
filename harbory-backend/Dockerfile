FROM golang:1.23-alpine

WORKDIR /app
COPY . .
RUN go mod tidy && go build -o harbory

EXPOSE 8080
CMD ["./harbory"]