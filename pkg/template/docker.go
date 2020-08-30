package template

const Dockerfile = `FROM golang:latest AS builder
WORKDIR /app
ENV GO111MODULE=on

COPY . .

WORKDIR /app/cmd/{{.ProjectName}}
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build  -mod=vendor -o main
RUN ls

FROM alpine:latest

RUN apk --no-cache add ca-certificates
RUN apk add --no-cache zsh
RUN apk add tzdata

COPY --from=builder /app/cmd/{{.ProjectName}}/main /app/main
COPY --from=builder /app/deployments/config /app/deployments/config

WORKDIR /app
EXPOSE 8080

ENTRYPOINT [ "./main","server" ]
`