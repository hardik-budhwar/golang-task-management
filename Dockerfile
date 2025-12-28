FROM golang:1.21

WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN go mod tidy

COPY . .

# Run the app directly (no air)
CMD ["go", "run", "."]
