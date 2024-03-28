FROM golang:alpine as builder
RUN apk --no-cache add tzdata
ENV TZ=Asia/Kolkata
RUN mkdir /app
ADD . /app/
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go build -o main .

# Run stage
FROM alpine
RUN apk --no-cache add tzdata
ENV TZ=Asia/Kolkata
WORKDIR /app
COPY --from=builder /app/main .
CMD [ "/app/main" ]
