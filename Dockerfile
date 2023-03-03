FROM golang as dev

WORKDIR /app

COPY . .

EXPOSE 5014

CMD air
