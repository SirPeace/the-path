FROM golang:1.25-alpine3.23 AS build
WORKDIR /app
COPY . .
RUN go build .

FROM alpine:3.23 AS runtime
WORKDIR /app
RUN addgroup -g 1000 appgroup && adduser -u 1000 -G appgroup -S appuser
COPY --from=build --chown=appuser:appgroup /app/the-path .
USER appuser
EXPOSE 8080
ENTRYPOINT [ "./the-path" ]
