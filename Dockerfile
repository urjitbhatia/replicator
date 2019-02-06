# Replicator is a daemon that provides automatic scaling of Nomad jobs and
# worker nodes.
#
# docker run --rm -it \
# 			 --name replicator \
#				 urjitbhatia/replicator agent

FROM alpine:edge
LABEL maintainer Urjit Singh Bhatia<(urjitsinghbhatia@gmail.com> (@urjitbhatia)
LABEL documentation "https://github.com/urjitbhatia/replicator"

ENV REPLICATOR_VERSION v2.0.0

WORKDIR /usr/local/bin/

RUN     apk --no-cache add \
        ca-certificates

RUN buildDeps=' \
                bash \
                wget \
        ' \
        set -x \
        && apk --no-cache add $buildDeps \
        && wget -O replicator https://github.com/urjitbhatia/replicator/releases/download/${REPLICATOR_VERSION}/linux-amd64-replicator \
        && chmod +x /usr/local/bin/replicator \
        && apk del $buildDeps \
        && echo "Build complete."

ENTRYPOINT [ "replicator" ]
CMD [ "agent", "--help" ]
