FROM golang:1.19.3-alpine AS stage1
WORKDIR /app
COPY . /app
RUN go install
RUN go build -o /app/main

FROM alpine:latest
WORKDIR /app
COPY --from=stage1 /app/main /app/main
EXPOSE 80
CMD ["./main"]
