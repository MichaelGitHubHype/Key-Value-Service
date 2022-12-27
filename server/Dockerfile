FROM golang:1.18-alpine

# Set destination for COPY
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
COPY *.go ./

# Build
RUN go build -o /server

# Run
CMD [ "/server" ]
