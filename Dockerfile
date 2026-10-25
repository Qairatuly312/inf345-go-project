FROM golang:1.27.1 AS build

WORKDIR /src

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /app .

# Runtime stage
FROM gcr.io/distroless/static-debian13:nonroot

COPY --from=build /app /app

USER nonroot:nonroot

ENV PORT=8080
EXPOSE 8080

ENTRYPOINT ["/app"]