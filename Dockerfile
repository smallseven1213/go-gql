FROM golang:1.15.3-alpine3.12
WORKDIR /graphql
ADD . /graphql
RUN go build .
EXPOSE 9090
CMD ["go", "run", "."]
# ENTRYPOINT ./