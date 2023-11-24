FROM ubuntu
MAINTAINER wang "calm.wang@hotmail.com"
ENV REFRESHED_AT = 2023-02-19

WORKDIR wkspace

ADD ../city-node-server /wkspace/city-node-server

RUN apt-get update
RUN apt-get install -y ca-certificates
RUN apt-get install -y vim

RUN chmod -R 777 /wkspace/city-node-server

VOLUME ["/project"]

ENTRYPOINT ["/wkspace/city-node-server"]




# 运行容器
# docker run -itd  --network=dev -p 9000:8000 -v /Users/cisco/Desktop/workspace:/project --name city-node city-node
# docker run -itd  --network=into-background_into_backend -p 8000:8000 -v /wkspace:/project --name city-node city-node:0.1.9

# docker run -itd  --network=dev -p 9000:8000 -v /wkspace/docker/project:/project --name city-node city-node



