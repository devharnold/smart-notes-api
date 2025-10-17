FROM golang:1.24.6

# Current working directory
WORKDIR /smart-notes-api

# Download Go Modules
COPY go.mod ./
RUN go mod download

# Copy the Source Code
COPY *.go ./

# Compile the Application
RUN CGO_ENABLED=0 GOOS=linux go build -o /smart-notes-api

EXPOSE 8080

# Tell Docker what to run
CMD ["/docker-smart-notes-api"]
