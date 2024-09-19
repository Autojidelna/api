# Nový projekt z CORU

## Přejmenování projektu

1. Otevřete projekt v CORU ve vscode.
2. přejmenujte všechny výskyty `coree` na nový název projektu.
3. přejmenujte všechny výskyty `registry.appelevate.cz` na nový název domény.
4. přejmenujte všechny výskyty `ssh.appelevate.cz` na nový název domény.
5. přejmenujte všechny výskyty `api.appelevate.cz` na nový název domény.

## Nastavení alpine serveru

### Zabezpečení serveru

1. Vypněte password login u ssh
   - Otevřete soubor `/etc/ssh/sshd_config`
   - Upravte následující řádky:
   ```
       PasswordAuthentication no
       ChallengeResponseAuthentication no
   ```
   - Restartujte ssh server - `rc-service sshd restart`
2. Nainstalujte docker
   - `apk add docker`
3. Nainstalujte docker-compose
   - `apk add docker-compose`
4. Zapněte docker
   - `rc-update add docker boot`
   - `rc-service docker start`
5. [Vytvořte tunel v cloudflare](https://one.dash.cloudflare.com/)
6. Nainstalujte cloudflare tunnel
   - `docker run -d --restart always --network="host" cloudflare/cloudflared:latest tunnel --no-autoupdate run --token <your-token>
7. Vytvořte 3 public hostnames
   - `api.<your-domain>.cz` - http://localhost:80
   - `registry.<your-domain>.cz` - http://localhost:5000
   - `ssh.<your-domain>.cz` - ssh://localhost:22
8. Uložte ssh doménu do environmentu `ssh` jako variable v github actions
   - `SSH_HOST` - `ssh.<your-domain>.cz`
9. Vytvořte aplikaci pro ssh
   - `ssh.<your-domain>.cz`
   - Vytvořte pravidlo `allow` pro váš email
   - Vytvořte pravidlo `service auth` pro nový service token.
   - Vyberte ssh v browser rendering
10. Přidejte service token do environmentu `ssh` v github actions:

    - `CLOUDFLARE_CLIENT_ID` - client id
    - `CLOUDFLARE_CLIENT_SECRET` - client secret
    - toto pouze hodnoty, ne celý header

11. Zapněte firewall
    - zakažte inbound komunikaci
    - povolte outbound komunikaci
12. Připojte se k ssh pomocí cloudflared (na vašem pc)

    - nainstalujte [cloudflared](https://developers.cloudflare.com/cloudflare-one/connections/connect-networks/downloads/)
    - upravte ~/.ssh/config

```
Host ohp
      HostName ssh.<your-domain>
      ProxyCommand cloudflared access ssh --hostname %h
      User root
```

```bash
ssh ohp
```

### Nastavení registry

1. Vytvořte složku pro registry

```bash
mkdir -p /var/lib/docker_registry/auth
mkdir -p /var/lib/docker_registry/data
```

2. vytvořte soubor `docker-compose.yml`

```bash
cd /var/lib/docker_registry
vim docker-compose.yml
```

3. Vyplňte soubor obsahem ze souboru [registry-docker-compose.yml](registry-docker-compose.yml)

4. Vytvořte hesla pro registry

```bash
apk add apache2-utils # nainstaluje nástroj pro vytvoření hesel
cd auth
htpasswd -Bc registry.password username # použijte vlastní jméno. Po zadání hesla se heslo uloží do souboru registry.password
htpasswd -B registry.password github-actions # použijte vlastní jméno. Po zadání hesla se heslo uloží do souboru registry.password
cd ..
docker-compose up -d
```

5. Vytvořte dva secrety v enviromentu `docker`

   - `REGISTRY_USERNAME` - jméno uživatele pro registry
   - `REGISTRY_PASSWORD` - heslo uživatele pro registry

6. Přihlašte se do registry na serveru

```bash
docker login registry.<your-domain>.cz
```

### Nastavení deploy skriptu

1. Nainstalujte git

```bash
apk add git
```

2. Přidejte soukromý klíč pro git ssh

```bash
cd ~/.ssh
vim id_rsa
chmod 600 id_rsa
```

sem nyní vložte váš soukromý klíč, který má pouze read práva k repozitáři. (deploy klíč)

3. Vytvořte soubor a vložte tam obsah [deploy_docker.sh](deploy_docker.sh)

   - pozměňte zde také jméno repozitáře

```bash
cd /root/
vim deploy_docker.sh
```

4. Přidejte práva pro spuštění skriptu

```bash
chmod +x deploy_docker.sh
```

5. Přidejte force command do authorized_keys pro ssh klíč v github actions

```bash
vim /root/.ssh/authorized_keys
```

Přidejte následující řádek:

```
no-agent-forwarding,no-X11-forwarding,no-port-forwarding,no-pty,command="/root/deploy_docker.sh" ssh-ed25519 AAAAKEY
```

6. Nyní k tomuto public key přidejte privátní klíč do github secrets environmentu `ssh`

   - `SSH_PRIVATE_KEY` - privátní klíč pro deploy

7. Spusťte ssh keyscan pro github

```bash
ssh-keyscan github.com >> /root/.ssh/known_hosts
```

8. Spusťte ssh keyscan pro github actions

```bash
ssh-keyscan localhost > known_hosts.txt
cat known_hosts.txt
```

9. Přidejte nový known_hosts do github secrets environmentu `ssh`

   - `SSH_HOST_KEY` - první řádek z known_hosts souboru bez localhostu - příklad:

   ```
   ssh-rsa AAAAB3NzaC1yc2EA5AADAQ9BAAABgQCi+K2an3vPTqm36f8B+1I/b0ndMohNQU8HesPZtkkumxZ8usihzdV7LB48dvwkI/vjYOWuEHWx1tNHbB1Pv/okPdOukNPCCrpRNU3TBchQpag48b5pzZTaT3RtINBSIH5MYMY0HMJ+Whdr2TgnBMEtlyTld2rCG+6ahP82OS2JEd6DBR5aLl0Fvy0hcr7fkq/iBlQ8p720zaeYXa6tRkKB0eKwCoci3xzcEUNG74Vafmhwp+HPGX2+eTLXadPYuHIYI27An0b31r589iM5onx5KVYqzk6JPxEJy0AyTAJjSWP19YRDa8/DGHeCq8PznnM5AMyAuaRjgZb5DD5K3aq4eWh/H6QlAx+0qCn4nQDGOqvfvftXSQ6ki9zHdbAJe8lTDNAx/ihRmvz+DCJqGfy2WJVZZnM90L78TP0od5hD3BFIXfjKt1s9jj+uep1BnjWNNjcUJsMuQevxybm1xsoM8RKKCljjUfr1Jc+2a7lkyIXBnE=
   ```

10. Zapněte swarm mode

```bash
docker swarm init
```

11. Nastavte secrets

```bash
echo "your-secure-username" | docker secret create postgres_user -
echo "your-secure-password" | docker secret create postgres_password -
echo "your-db-name" | docker secret create postgres_db -
```

12. Spusťte deploy skript

```bash
sh /root/deploy_docker.sh
```

### Nastavení zálohování postgresql db

1. Použijte skript [backup_postgres.sh](backup_postgres.sh) a upravte v něm hodnoty v `<>`

   - hlavně <user> a <db-name> a upravte bucket name pro váš nově vytvořený bucket v b2

2. vytvořte nový bucket v b2 a přidejte do něj klíč aplikace

3. Nainstalujte b2 cli a gpg

```bash
apk add pipx
apk add gnupg
pipx install b2
```

4. Přidejte klíč do b2 cli

```bash
b2 account authorize <account-id> <application-key>
```

5. Přidejte práva pro spuštění skriptu

```bash
chmod +x backup_postgres.sh
```

6. Nastavte cron job pro zálohování

```bash
crontab -e
```

a přidejte následující řádek

```
0 3 * * * /root/backup_postgres.sh
```
