# Create node image
FROM golang:alpine

# Create app directory
RUN mkdir /app

# Create app directory
WORKDIR /app

# Copy file to /app directory
COPY . /app

# Clone and Build
RUN export GO111MODULE=on
RUN go get github.com/sholehbaktiabadi/go-company-api/server.go
RUN cd /app && git clone github.com/sholehbaktiabadi/go-company-api.git
RUN cd /app/go-company-api/main && go build

# Expose port 8080
EXPOSE 8080

# run pm2-runtime
CMD ["/app/go-company-api/main/main"]
