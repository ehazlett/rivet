FROM scratch
COPY rivet /usr/local/bin/rivet
COPY certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
EXPOSE 8080
ENTRYPOINT ["/usr/local/bin/rivet"]
