FROM ubuntu:20.04

MAINTAINER Kushal Das <kushal@sunet.se>


RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y locales curl make git

RUN sed -i -e 's/# en_US.UTF-8 UTF-8/en_US.UTF-8 UTF-8/' /etc/locale.gen && \
    dpkg-reconfigure --frontend=noninteractive locales && \
    update-locale LANG=en_US.UTF-8

ENV LANG en_US.UTF-8

RUN curl -OL https://go.dev/dl/go1.19.4.linux-amd64.tar.gz
RUN tar -xvf go1.19.4.linux-amd64.tar.gz
WORKDIR /nglue-go

CMD ["bash", "scripts/build.sh"]
