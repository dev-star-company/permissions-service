FROM golang:1.24-alpine

# Set the working directory inside the container
WORKDIR /opt/app
COPY . .

# Install necessary packages
RUN apk add --no-cache git
ENV GOPRIVATE=github.com/dev-star-company

RUN go mod download

WORKDIR /opt/app/internal/app/ent
RUN go run -mod=mod entgo.io/ent/cmd/ent generate --feature intercept,schema/snapshot ./schema

WORKDIR /opt/app

RUN go build cmd/main.go

# Run the application in development mode (use a process manager like nodemon if you want hot reloading)
CMD [ "./main" ]