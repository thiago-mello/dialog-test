FROM golang:1.24.2 AS build

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 GOOS=linux go build -o api src/main.go

FROM alpine:3.21

WORKDIR /app

COPY --from=build /app/api .

COPY --from=build /app/db/ ./db/

ENV PORT=8080

EXPOSE 8080

CMD [ "./api" ]
