FROM curlimages/curl:7.76.1

#Use included variables with docker buildkit to determine what yq to dowload
ARG TARGETOS
ARG TARGETARCH

#Download yq from github
ADD https://github.com/mikefarah/yq/releases/download/v4.7.1/yq_${TARGETOS}_${TARGETARCH} ./

# Elevate to root to move yq into right location
USER root

#Set the executable bit and move it into the path
RUN chmod +x yq_linux_amd64 && mv yq_linux_amd64 /usr/bin/yq

# Go back to base user in curl image
USER curl_user