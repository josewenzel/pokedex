FROM  golang:1.19

WORKDIR /pokedex

COPY . ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /pokedex-app src/main.go

CMD ["/pokedex-app"]