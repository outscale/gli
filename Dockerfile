# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
# debug image is used to get busybox
FROM gcr.io/distroless/static-debian13:debug@sha256:a77dbde155f73bf787b42533389b57eb70861bcd80f5186760de48f4cf206a7d
ARG TARGETPLATFORM
ENTRYPOINT ["/usr/bin/octl"]
COPY $TARGETPLATFORM/octl /usr/bin/
