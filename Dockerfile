FROM golang:1.18.1-buster
WORKDIR /app
RUN go install github.com/spf13/cobra-cli@latest
ENTRYPOINT [ "/go/bin/cobra-cli" ]
