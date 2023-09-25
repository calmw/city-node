FROM ubuntu
MAINTAINER wang "calm.wang@hotmail.com"
ENV REFRESHED_AT = 2023-02-19

WORKDIR wkspace

ADD city-node-server /wkspace/city-node-server

COPY ca-bundle.crt /etc/ssl/certs/ca-bundle.crt
COPY ca-bundle.trust.crt /etc/ssl/certs/ca-bundle.trust.crt
##
#RUN apt-get update
#RUN apt-get -y install golang

##
RUN chmod -R 777 "/wkspace/city-node-server"

VOLUME ["/project"]

EXPOSE 8000

#ENTRYPOINT ["/wkspace/city-node-server"]




# 运行容器
# docker run -itd  --network=dev -p 9000:8000 -v /Users/cisco/Desktop/workspace:/project --name city-node city-node

# docker run -itd  --network=dev -p 9000:8000 -v /wkspace/docker/project:/project --name city-node city-node



