FROM debian:latest
EXPOSE 8000

ADD oauth /oauth

CMD ["/oauth"]
