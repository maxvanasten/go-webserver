# Stop running server and update files
ssh max@172.233.59.9 'tmux kill-session -t go-webserver; /home/max/scripts/update.sh'
./scripts/cp_tls.sh
ssh max@172.233.59.9 '/home/max/scripts/unpack-ssl-bundle.sh; tmux new-session -d -s go-webserver /home/max/scripts/start-server.sh'
