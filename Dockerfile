FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o CI-CD-Pipeline .

EXPOSE 8000

CMD ["./CI-CD-Pipeline"]