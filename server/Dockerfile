FROM golang As build
MAINTAINER wang "calm.wang@hotmail.com"
ENV REFRESHED_AT = 2023-02-19
ENV WKDIR = wkspace

WORKDIR $WKDIR

ADD city_node /wkspace/city_node
ADD task /wkspace/task
ADD start.sh /wkspace/start.sh

#RUN apk add --update bash && rm -rf /var/cache/apk/*
#RUN apk add vim
#RUN apk add git
#RUN apk add go
#RUN apk add openssh-client
RUN apt-get update
RUN apt-get -y install golang
RUN mkdir -p /root/.ssh/
RUN ssh-keyscan -H github.com >> ~/.ssh/known_hosts
RUN git clone -b test_cityPioneer https://github.com/calmw/city-node.git
#RUN GO111MODULE=off
RUN export GO111MODULE=on 
WORKDIR $WKDIR/city-node/server
RUN go build api/city-node-server.go
#RUN cd /wkspace/city-node/server/api && go build city_node.go


#RUN chmod -R 777 "/wkspace/city_node"
#RUN chmod -R 777 "/wkspace/task"
#RUN chmod -R 777 "/wkspace/start.sh"

VOLUME ["/project"]

EXPOSE 8000

#ENTRYPOINT ["nohup", "/wkspace/city_node"," & ", "nohup","/wkspace/task"]
#ENTRYPOINT ["/wkspace/start.sh"]




# 运行容器
# docker run -itd  --network=dev -p 9000:8000 -v /Users/cisco/Desktop/workspace/docker/dev/project:/project --name citynode citynode

# docker run -itd  --network=dev -p 9000:8000 -v /wkspace/docker/project:/project --name citynode city-node



