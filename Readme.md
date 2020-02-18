# Utility to proxy requests to quintype

## Usage

```bash
./http-to-quintype-proxy https://path/to/quintype/server/ 127.0.0.1:8000
```

## Installing As An Upstart Service

In order to run as an upstart service, please run the following scripts as root, replacing the path as needed

```bash
curl -L https://github.com/quintype/http-to-quintype-proxy/raw/master/build/linux-amd64/http-to-quintype-proxy > /usr/local/bin/http-to-quintype-proxy
chmod +x /usr/local/bin/http-to-quintype-proxy
cat > /etc/init/http-to-quintype-proxy <<EOF
description "Proxy Requests to Upsteam Quintype Server"

respawn
respawn limit 15 5

start on runlevel [2345]
stop on runlevel [06]

exec /usr/local/bin/http-to-quintype-proxy https://madrid.quintype.io/ 127.0.0.1:8000
EOF
initctl reload-configuration
service http-to-quintype-proxy start
```
