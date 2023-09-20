FROM alpine
MAINTAINER wang "calm.wang@hotmail.com"
ENV REFRESHED_AT = 2023-02-19

WORKDIR wkspace

ADD city_node /wkspace/city_node
ADD task /wkspace/task
ADD start.sh /wkspace/start.sh

RUN chmod -R 777 "/wkspace/city_node"
RUN chmod -R 777 "/wkspace/task"
RUN chmod -R 777 "/wkspace/start.sh"

VOLUME ["/project"]

EXPOSE 8000

ENTRYPOINT ["/wkspace/start.sh"]




# 运行容器
# docker run -itd  --network=dev -p 9000:8000 -v /Users/cisco/Desktop/workspace/docker/dev/project:/project --name citynode citynode

# docker run -itd  --network=dev -p 9000:8000 -v /wkspace/docker/project:/project --name citynode city-node



