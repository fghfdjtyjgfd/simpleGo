

FROM golang:1.17

ENV TZ=Asia/Bangkok

RUN mkdir /app
COPY . /app

WORKDIR /app
RUN go build -o main .

EXPOSE 4000
CMD ["/app/main"]
