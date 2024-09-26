FROM scratch

WORKDIR $GOPATH/src/gocode/go-gin-example
COPY . $GOPATH/src/gocode/go-gin-example

EXPOSE 8000
CMD ["./go-gin-example"]