#!/usr/bin/env bash
debianBuilderFile="debian-builder-linux-amd64"
curl -s https://api.github.com/repos/danmademe/debian-builder/releases | grep browser_download_url | grep ${debianBuilderFile} | head -n 1 | cut -d '"' -f 4 | wget -i -


chmod +x ${debianBuilderFile}
mv ${debianBuilderFile} /usr/local/bin/debian-builder
