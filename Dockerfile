FROM scratch
COPY dispatch-http-ping /
ENTRYPOINT ["/dispatch-http-ping"]
