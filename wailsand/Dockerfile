FROM golang
#WORKDIR /usr/src/app

COPY ./apt.tar.gz ./

RUN tar -zxvf apt.tar.gz ;
RUN apt update \
    && apt upgrade -y \
    && go install github.com/wailsapp/wails/v2/cmd/wails@latest ;
#    && apt install -y \
#        gcc \
#        libgtk-3-dev \
#        libwebkit2gtk-4.0-dev \
#        npm \
#        pkg-config \

#CMD ["dpkg -i /go/var/cache/apt/archives/*.deb"]
#CMD ["wails init -n myproject -t svelte-ts"]
#CMD ["cd myproject/ && wails dev"]
