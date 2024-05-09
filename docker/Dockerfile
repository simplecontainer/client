FROM alpine

RUN apk add --no-cache libc6-compat

RUN mkdir /opt/smr
RUN mkdir -P /home/smr-client/.ssh
RUN adduser -D smr-client

COPY smr /opt/smr-client/smr

RUN chmod +x /opt/smr-client/smr
RUN chown -R smr-client /opt/smr-client/smr
RUN chown -R smr-client /home/smr-client/.ssh

USER smr-client

ENTRYPOINT /opt/smr-client/smr