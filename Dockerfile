FROM golang:1.10-alpine

RUN apk add --no-cache nodejs nodejs-npm

COPY . /go/src/demos/docker-node
WORKDIR /go/src/demos/docker-node
RUN go install -v
ENV PUPPETEER_SKIP_CHROMIUM_DOWNLOAD true
RUN npm install puppeteer@1.4.0

CMD ["docker-node"]