FROM node:16

ENV CHROME_BIN="/usr/bin/chromium"

RUN apt-get update

RUN apt install chromium -y

RUN npm install -g  @angular/cli@14

# get go
RUN curl -OL https://go.dev/dl/go1.18.5.linux-amd64.tar.gz

RUN tar -C /usr/local -xvf go1.18.5.linux-amd64.tar.gz

RUN echo 'export PATH=$PATH:/usr/local/go/bin' >>/home/node/.bashrc

RUN ["/bin/bash", "-c", "source /home/node/.bashrc"]

ENTRYPOINT []
