#!/usr/bin/env bash
dockerServiceFile="docker-service-linux-amd64"
curl -s https://api.github.com/repos/danmademe/docker-service/releases | grep browser_download_url | grep ${dockerServiceFile} | head -n 1 | cut -d '"' -f 4 | wget -i -


chmod +x ${dockerServiceFile}
mv ${dockerServiceFile} /usr/local/bin/docker-service
