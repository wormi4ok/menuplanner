FROM gcr.io/distroless/static:nonroot

ARG TARGETPLATFORM

ENV MP_HOST=0.0.0.0
ENV MP_PORT=8081

EXPOSE 8081

COPY $TARGETPLATFORM/menuplanner /menuplanner

ENTRYPOINT ["/menuplanner"]
