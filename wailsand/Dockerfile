FROM golang
#WORKDIR /usr/src/app

COPY ./aptData ./aptData

RUN cat ./aptData/apt.tar.gz.part? > apt.tar.gz && \
    tar -zxvf apt.tar.gz ;
RUN apt update && \
    apt upgrade -y && \
    go install github.com/wailsapp/wails/v2/cmd/wails@latest && \
    apt-get install -y ./var/cache/apt/archives/*.deb && \
    npm install axios;
RUN wails init -n myproject -t svelte-ts
# RUN cd myproject/ && wails dev
