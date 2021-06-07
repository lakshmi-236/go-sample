FROM  golang:alpine
RUN mkdir app
COPY . /app
WORKDIR /app
RUN go build -o main .
CMD ["app/main"]
FROM  golang:alpine
RUN mkdir app
COPY . /app
WORKDIR /app
RUN go build -o /app/main .
CMD ["app/main"]
