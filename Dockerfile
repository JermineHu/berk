FROM alpine:edge
MAINTAINER Jermine <Jermine.hu@qq.com>
COPY . /root/berk
RUN apk add --no-cache --virtual .build-deps make  go=1.9-r3 git gcc musl-dev binutils ;\
     go version ;\
     go env ;\
    mkdir -p /root/go/src/git.vdo.space &&  mv /root/berk /root/go/src/git.vdo.space/ ;\
 #   go get github.com/Masterminds/glide;\
 #    cd /root/go/src/github.com/Masterminds/glide ;\
 #    make build ;\
 #    make install ;\
    cd /root/go/src/git.vdo.space/berk ;\
    go build -a -v -installsuffix cgo -o /bin/berk ;\
 #   which glide ;\
 #   rm $(which glide) ;\
	rm -r /var/cache/apk ; \
	rm -r /usr/share/man ; \
	rm -rf /root/go ;\
    apk del .build-deps

 EXPOSE  80 443 8000
 CMD ["/bin/berk"]