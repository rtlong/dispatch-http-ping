FROM scratch
ADD bin/http_ping /http_ping
ENTRYPOINT ["/http_ping"]
