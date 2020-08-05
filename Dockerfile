FROM alpine
ADD rule-srv-service /rule-srv-service
ENTRYPOINT [ "/rule-srv-service" ]
