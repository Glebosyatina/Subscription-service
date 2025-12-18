FROM golang:1.25-bookworm

RUN curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
RUN echo "deb https://packagecloud.io/golang-migrate/migrate/debian/ bookworm main" > /etc/apt/sources.list.d/migrate.list
RUN cat /etc/apt/sources.list.d/migrate.list
RUN apt-get update
RUN apt-get install migrate
RUN apt-get install netcat-openbsd

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o main ./cmd/main.go

#RUN make migrate_up
COPY entrypoint.sh ./entrypoint.sh
RUN chmod +x ./entrypoint.sh


EXPOSE 8080
ENTRYPOINT ["./entrypoint.sh"]
