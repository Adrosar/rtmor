FROM golang:1.16-alpine3.14
COPY / /app
RUN cd /app/cmd/rtmor && go build && mv -f ./rtmor /app/build
EXPOSE 8888
CMD /app/build/rtmor -start -listen 0.0.0.0:8888
