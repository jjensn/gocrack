#/bin/bash
adduser --disabled-password --gecos "" --uid $USER_ID gocrack
mkdir -p /opt/gocrack/files/task
mkdir -p /opt/gocrack/files/tmp
mkdir -p /opt/gocrack/files/engine

chown -R gocrack:gocrack /opt/gocrack/
cd /opt/gocrack/

su gocrack -c "/usr/local/bin/gocrack_server -config /opt/gocrack/config.yaml"