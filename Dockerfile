# Build the Go app
FROM golang:1.22.6 as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o /backend-assignment

# Run the Go app
FROM golang:1.22.6

WORKDIR /

COPY --from=build /backend-assignment /backend-assignment
COPY .env .env

EXPOSE 8080

CMD ["/backend-assignment"]
