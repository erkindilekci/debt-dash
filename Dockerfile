FROM golang:1.22.6 AS builder

WORKDIR /build
COPY . .
RUN go mod download
RUN go build -o ./debtdash cmd/web/main.go

FROM gcr.io/distroless/base-debian12

WORKDIR /app
COPY --from=builder /build/debtdash ./debtdash
COPY --from=builder /build/templates ./templates
COPY --from=builder /build/static ./static
CMD ["/app/debtdash"]
