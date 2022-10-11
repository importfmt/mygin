FROM centos:7.9.2009
MAINTAINER 26n119e

WORKDIR /opt
RUN mkdir /opt/config
COPY mygin_linux /opt
COPY config/application.yml /opt/config/application.yml
EXPOSE 8888

CMD ["/opt/mygin_linux"]

