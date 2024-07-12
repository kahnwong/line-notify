FROM golang:1.22-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -o /line-notify

FROM alpine:latest AS build-release-stage

WORKDIR /

COPY --from=build-stage /line-notify /line-notify

RUN chmod +x /line-notify

ENTRYPOINT ["/line-notify"]
