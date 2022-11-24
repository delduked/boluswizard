FROM golang:latest AS stage1
WORKDIR /app
COPY . /app
RUN go get
RUN go build -o /app/main

FROM alpine:latest
WORKDIR /app
COPY --from=stage1 /app/main /app/main
EXPOSE 80
CMD ["./main"]
