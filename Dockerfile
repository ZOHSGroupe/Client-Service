FROM ubuntu:latest
LABEL authors="hajar"

ENTRYPOINT ["top", "-b"]