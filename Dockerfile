FROM golang:1.22-alpine

# Install make and other build essentials
RUN apk add --no-cache make

# Install air for live reloading
RUN go install github.com/air-verse/air@latest

# Set the working directory
WORKDIR /app

# copy go.mod, go.sum and Makefile 
COPY go.* ./
COPY Makefile ./

# download dependencies
RUN make download

# copy the rest of the application
COPY . . 

# build the application
RUN make build

# Expose the port
EXPOSE 8080

# Run the application
CMD [ "air", "-c", ".air.toml" ]