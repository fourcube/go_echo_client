FROM centurylink/ca-certs
WORKDIR /app
COPY go_echo_client /app/
CMD ["/app/go_echo_client"]
