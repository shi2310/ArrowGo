FROM alpine:latest

WORKDIR /arrowgo
COPY  upload .
COPY ArrowGo .
RUN chmod +x ArrowGo
CMD ./ArrowGo