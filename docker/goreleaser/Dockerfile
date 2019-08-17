FROM scratch as binBase
COPY replicator /

FROM alpine:edge
LABEL maintainer Urjit Singh Bhatia<(urjitsinghbhatia@gmail.com> (github: @urjitbhatia)
LABEL documentation "https://github.com/urjitbhatia/replicator"

ARG REPLICATOR_VERSION=v2.2.2

WORKDIR /usr/local/bin/

RUN     apk --no-cache add \
    ca-certificates

RUN buildDeps=' \
    bash \
    ' \
    set -x \
    && apk --no-cache add $buildDeps \
    && apk del $buildDeps \
    && echo "Build complete."

COPY --from=binBase /replicator /usr/local/bin/replicator
ENTRYPOINT [ "replicator" ]
CMD [ "agent", "--help" ]
