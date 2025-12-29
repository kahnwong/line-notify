FROM golang:1.25-alpine AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 go build -ldflags "-w -s" -o /line-notify && chmod +x /line-notify

FROM gcr.io/distroless/static-debian11:nonroot AS build-release-stage

WORKDIR /
COPY --from=build-stage /line-notify /line-notify

ENTRYPOINT ["/line-notify"]
