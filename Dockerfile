# *************************************
#
# OpenGM
#
# *************************************

FROM alpine:3.14

MAINTAINER XTech Cloud "xtech.cloud"

ENV container docker
ENV MSA_MODE release

EXPOSE 18813

ADD bin/ogm-config /usr/local/bin/
RUN chmod +x /usr/local/bin/ogm-config

CMD ["/usr/local/bin/ogm-config"]
