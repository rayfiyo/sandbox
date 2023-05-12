# Equivalent: docker pull busybox:1.36
# 相当: docker pull busybox:1.36
#FROM busybox:1.36
FROM ubuntu

# "WORKDIR" specifies the $home directory.
# "WORKDIR" は、$home を指定する。
WORKDIR /root

# "COPY" copies some files to docker.
# "COPY" は、docker にファイルをコピーする。
COPY memo.txt ./

# "RUN" is executed when building.
# "RUN" は、ビルド時に コマンドを実行する。
#RUN apt update; apt upgrade -y;
RUN echo "When building" >> memo.txt

# "CMD" is executed when the container is created after the build.
# "CMD" は、ビルド後の コンテナ生成時に実行する。
#CMD ["bash -c \"echo \"When the container\" >> /root/memo.txt\""]
#CMD ["bash -c \"cat /root/memo.txt\""]
CMD ["/usr/bin/pwd"]
