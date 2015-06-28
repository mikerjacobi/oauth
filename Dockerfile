FROM debian:latest
EXPOSE 80

ADD oauth /oauth

CMD ["/oauth"]
