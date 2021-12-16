# Webserver Log4j Honeypot

This honeypots runs fake Webserver waiting to be exploited. 
Payload classes are saved to `payloads` directory.

Forked from https://github.com/Adikso/minecraft-log4j-honeypot

## Requirements
- Golang 1.16+

## Running

### Natively
```
git clone https://github.com/schadom/webserver-log4j-honeypot.git
cd webserver-log4j-honeypot
go build .
./webserver-log4j-honeypot
```

### Using docker
```
git clone https://github.com/schadom/webserver-log4j-honeypot.git
cd webserver-log4j-honeypot
docker build -t webserver-log4j-honeypot:latest .
docker run --rm -it --mount type=bind,source="${PWD}/payloads",target=/payloads --user=`id -u` -p 8888 webserver-log4j-honeypot:latest
```
