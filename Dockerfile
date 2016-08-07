From sephvelu/golang-alpine:nightly

ADD ./ /gopath/src/github.com/SephVelut/test

WORKDIR /gopath/src/github.com/SephVelut/test
RUN go build main.go

CMD ["./main"]
