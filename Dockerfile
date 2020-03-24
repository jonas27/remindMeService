FROM golang as builder
EXPOSE 3000
COPY . .
RUN useradd scratchuser && \
    export GOPATH="" && go env  && go mod vendor  && cd cmd && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o metrics .

FROM scratch
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs
COPY --from=builder /go/cmd/metrics /
COPY --from=builder /etc/passwd /etc/passwd
USER scratchuser
CMD ["./metrics"]
