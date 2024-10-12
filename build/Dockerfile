FROM alpine:latest

ENV API_PORT=8080
ENV METRICS_PORT=9095

WORKDIR /app

COPY bin/scraperservice ./

RUN chmod +x ./scraperservice

EXPOSE 8080
EXPOSE 9095

CMD ["./scraperservice"]

# docker build -t scraperservice .
# docker run -p 8080:8080 -p 9095:9095 scraperservice
