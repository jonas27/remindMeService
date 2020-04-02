FROM golang as builder
EXPOSE 2001
COPY . .
# RUN useradd scratchuser && \
#     export GOPATH="" && go env  && go mod vendor  && cd src && \
#     go build -o main .
# RUN ls /go/src
# RUN echo $PWD
# WORKDIR /go/src
# RUN  GO111MODULE=off go test
# CMD ["/go/src/main"]
RUN useradd scratchuser && \
    export GOPATH="" && go env  && go mod vendor  && cd src && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o metrics .

FROM scratch
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs
COPY --from=builder /go/src/metrics /
COPY --from=builder /etc/passwd /etc/passwd
USER scratchuser
CMD ["./metrics"]
