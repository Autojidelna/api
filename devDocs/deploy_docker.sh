rm -rf /root/clone
mkdir /root/clone
mkdir -p /var/lib/postgresql/data/
cd /root/clone
git init
git remote add origin git@github.com:app-elevate/autojidelna.go-gin-api.git
git sparse-checkout init --cone
git sparse-checkout set docker-compose.yml
git pull origin main
rm -rf /root/godocker
mkdir -p /root/godocker/
mv docker-compose.yml /root/godocker/
cd /root/godocker/
sh /root/backup_postgres.sh
docker stack deploy --compose-file docker-compose.yml autojidelna --with-registry-auth
rm -rf /root/clone
