FROM alpine
EXPOSE 8080
ENV PORT 8080
COPY bin/http-server-header-dump-linux-amd64 /
CMD /http-server-header-dump-linux-amd64 --port $PORT