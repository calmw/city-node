FROM alpine
MAINTAINER wang "calm.wang@hotmail.com"
ENV REFRESHED_AT = 2023-02-19
ENV WKDIR = wkspace

WORKDIR $WKDIR

ADD city-node-server /wkspace/city-node-server

RUN chmod -R 777 "/wkspace/city-node-server"

VOLUME ["/project"]

EXPOSE 8000

ENTRYPOINT ["/wkspace/city-node-server"]




# 运行容器
# docker run -itd  --network=dev -p 9000:8000 -v /Users/cisco/Desktop/workspace/docker/dev/project:/project --name citynode citynode

# docker run -itd  --network=dev -p 9000:8000 -v /wkspace/docker/project:/project --name citynode city-node



