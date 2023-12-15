FROM golang:1.21.5
WORKDIR /usr/src/app
#COPY memo.txt ./

#RUN apt update; apt upgrade -y;
RUN apt update && apt install -y \
    gcc \
    libgtk-3-dev \
    libwebkit2gtk-4.0-dev \
    npm \
    pkg-config \
    && go install github.com/wailsapp/wails/v2/cmd/wails@latest

#CMD ["bash -c \"echo \"When the container\" >> /root/memo.txt\""]
#CMD ["bash -c \"cat /root/memo.txt\""]
#CMD ["/usr/bin/pwd"]
