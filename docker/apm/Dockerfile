FROM docker.elastic.co/apm/apm-server:9.2.2

USER root
COPY apm-server.yml /usr/share/apm-server/apm-server.yml
RUN chmod go-w /usr/share/apm-server/apm-server.yml

USER apm-server
