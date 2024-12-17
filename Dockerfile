FROM golang as builder

WORKDIR /app

RUN apt upgrade 

COPY go.mod go.sum ./
RUN go mod download && go mod tidy

COPY . .

RUN CGO_ENABLED=1 go build -trimpath -o whattoeat main.go

FROM debian:latest

RUN apt update && apt install -y ca-certificates

COPY --from=builder /app/whattoeat /bin

RUN chmod +x /bin/whattoeat

EXPOSE 2550

CMD ["whattoeat"]
