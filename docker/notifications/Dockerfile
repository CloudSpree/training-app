FROM registry.access.redhat.com/ubi8/ubi-minimal
RUN \
  mkdir -p /src \
  && groupadd -g 1000750000 appgrp \
  && useradd -l -r -u 1000750000 -g appgrp appuser \
  && chown -R appuser /src
WORKDIR /src
COPY main ./app

CMD ["/src/app"]

ARG VERSION=unknown
LABEL version=${VERSION}