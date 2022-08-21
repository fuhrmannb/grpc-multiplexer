FROM scratch
ENTRYPOINT ["/grpc-multiplexer-server"]
COPY grpc-multiplexer-server /