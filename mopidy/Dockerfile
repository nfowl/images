FROM ubuntu:22.04

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y \
      openssl \
      curl \
      python3 \
      python3-pip

RUN apt install -y \
    python3-gst-1.0 \
    gir1.2-gstreamer-1.0 \
    gir1.2-gst-plugins-base-1.0 \
    gstreamer1.0-plugins-good \
    gstreamer1.0-plugins-ugly \
    gstreamer1.0-tools      

COPY requirements.txt .
RUN pip3 install -r requirements.txt

VOLUME [ "/config" , "/data"  ]
EXPOSE 6600

ENTRYPOINT ["./entrypoint.sh"]
CMD ["mopidy", "--config", "/config/config.toml" ]
