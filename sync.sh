ssh max@172.233.59.9 '/home/max/scripts/update.sh'
./scripts/cp_tls.sh
ssh max@172.233.59.9 '/home/max/scripts/unpack-ssl-bundle.sh; /home/max/scripts/start-server.sh'
