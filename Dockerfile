ARG BUILD_IMAGE
ARG SERVICE_BASE_IMAGE

FROM $BUILD_IMAGE as build

COPY . /src
WORKDIR /src

ARG VERSION
ARG GO_BUILD_TAGS
ARG SERVICE_DIR
ARG GONOSUMDB=""
ARG GOPROXY=""
RUN ["/src/scripts/build.sh"]

FROM $SERVICE_BASE_IMAGE as prod

COPY --from=build /src/app /opt/app

ENTRYPOINT ["tini", "--", "/opt/app"]
