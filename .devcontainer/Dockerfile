# [Choice] Go version (use -bullseye variants on local arm64/Apple Silicon): 1, 1.18, 1.17, 1-bullseye, 1.18-bullseye, 1.17-bullseye, 1-buster, 1.18-buster, 1.17-buster
ARG VARIANT=1.18-bullseye
ARG USE_BAZEL_VERSION=5.2.0
FROM mcr.microsoft.com/vscode/devcontainers/go:0-${VARIANT}

# Install Bazel
ARG BAZELISK_VERSION=v1.10.1
# TODO Update bazel sha
ARG BAZELISK_DOWNLOAD_SHA=dev-mode
RUN curl -fSsL -o /usr/local/bin/bazelisk https://github.com/bazelbuild/bazelisk/releases/download/${BAZELISK_VERSION}/bazelisk-linux-amd64 \
    && ([ "${BAZELISK_DOWNLOAD_SHA}" = "dev-mode" ] || echo "${BAZELISK_DOWNLOAD_SHA} */usr/local/bin/bazelisk" | sha256sum --check - ) \
    && chmod 0755 /usr/local/bin/bazelisk \
    && ln -s /usr/local/bin/bazelisk /usr/local/bin/bazel

USER vscode

# Install buildifier
ARG BAZEL_BUILDIFIER_VERSION=5.1.0
RUN go install github.com/bazelbuild/buildtools/buildifier@${BAZEL_BUILDIFIER_VERSION}

ENV SGNL_LOG_LEVEL=DEBUG