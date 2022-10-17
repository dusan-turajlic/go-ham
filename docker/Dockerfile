ARG GO_VERSION="alpine3.16"

FROM golang:${GO_VERSION} as base

ENV USER_NAME="goham"
ENV GROUP_NAME="goham"
ENV GROUP_ID=70
ENV USER_ID=70

# Create a group and user
RUN addgroup -g "${GROUP_ID}" -S "${GROUP_NAME}" && \
    adduser --uid "${USER_ID}" -S "${USER_NAME}" -G "${GROUP_NAME}"

USER "${USER_NAME}"

ENV HOME="/home/${USER_NAME}"
ENV APPLICATION_DIR="${HOME}/application"

RUN mkdir -p "${APPLICATION_DIR}"


EXPOSE 8080
VOLUME ${APPLICATION_DIR}
WORKDIR "${APPLICATION_DIR}"


