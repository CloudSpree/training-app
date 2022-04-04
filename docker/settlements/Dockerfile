FROM registry.access.redhat.com/ubi8/ubi-minimal
WORKDIR /src
COPY main ./app
CMD ["/src/app"]

ARG VERSION=unknown
LABEL version=${VERSION}
