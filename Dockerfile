FROM ubuntu

RUN mkdir -p /src

WORKDIR /src

COPY . .

RUN ./build

EXPOSE 8080
EXPOSE 8081
EXPOSE 8082
EXPOSE 8083
EXPOSE 8084
EXPOSE 8085
EXPOSE 8086
EXPOSE 8087

ENTRYPOINT ./bin/analytics & &&  \
./bin/auth & && \
./bin/coupons & && \
./bin/events & && \
./bin/exporter & && \
./bin/guests & && \
./bin/participants & && \
./bin/simple_projection & && 