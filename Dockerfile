FROM codyperakslis/dev-go:1.14.4-v0
COPY go.mod go.sum ./
RUN go mod download