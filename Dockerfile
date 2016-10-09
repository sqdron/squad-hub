# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest

# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/golang/sqdron/squad-hub

COPY /go/src/github.com/golang/sqdron/squad-hub/config.json.sample /go/src/github.com/golang/sqdron/squad-hub/config.json

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/sqdron/squad-hub.git

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/squad-hub

# Document that the service listens on port 8080.
EXPOSE 5001