FROM golang:1.22

# Install make
RUN apt-get update && apt-get install -y make

# Install Google's wire
RUN go install github.com/google/wire/cmd/wire@latest

WORKDIR /app

COPY . .

RUN make build

EXPOSE 8080

CMD ["./build/backend"]