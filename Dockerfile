FROM golang:1.19.3-alpine3.15@sha256:eabc3aca6f6c4386369b5b067c9c210aeccd39e76907fa2f8f774fd59d83425a as builder

RUN apk add --update --no-cache ca-certificates tzdata git make bash && update-ca-certificates

ADD . /opt
WORKDIR /opt

RUN go install github.com/go-task/task/v3/cmd/task@latest

# CGO_ENABLED is set to allow the golang binary build to work on the ubi image
RUN pwd; find
RUN git update-index --refresh; CGO_ENABLED=0 ${GOPATH}/bin/task

FROM quay.io/app-sre/ubi8-ubi-minimal:8.9-1108.1706795067@sha256:f28a083503f91d97fd3ac2ec24d67c0fa705c086c4c01f31c33567ebab1dfcb8 as runner

COPY --from=builder /opt/obnpctl /bin/obnpctl

ARG BUILD_DATE
ARG VERSION
ARG VCS_REF
ARG DOCKERFILE_PATH

LABEL vendor="Ron Green" \
    name="geoegettica/obnpctl" \
    description="a CLI tool explain basic k8s networking" \
    io.k8s.display-name="geoegettica/obnpctl" \
    io.k8s.description="a CLI tool explain basic k8s networking" \
    maintainer="Ron Green <8326+rogreen@users.noreply.gitlab.cee.redhat.com>" \
    version="$VERSION" \
    org.label-schema.build-date=$BUILD_DATE \
    org.label-schema.description="a CLI tool explain basic k8s networking" \
    org.label-schema.docker.cmd="docker run --rm  geoegettica/obnpctl" \
    org.label-schema.docker.dockerfile=$DOCKERFILE_PATH \
    org.label-schema.name="geoegettica/obnpctl" \
    org.label-schema.schema-version="0.1.0" \
    org.label-schema.vcs-branch=$VCS_BRANCH \
    org.label-schema.vcs-ref=$VCS_REF \
    org.label-schema.vcs-url="https://github.com:georgettica/obnpctl" \
    org.label-schema.vendor="geoegettica/obnpctl" \
    org.label-schema.version=$VERSION

ENTRYPOINT ["/bin/obnpctl"]
