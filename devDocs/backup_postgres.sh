# Variables
BACKUP_PATH="/var/lib/postgresql/data/myBackups"
BACKUP_FILE="backup_$(date +%Y%m%d_%H%M%S).sql"
BUCKET_NAME="autojidelna-postgres-backup"
PERSISTENT_BACKUP_PATH="/var/psql_backups"

# Perform PostgreSQL backup
mkdir -p $BACKUP_PATH
/usr/bin/docker exec $(docker ps --filter name=autojidelna_db -q) pg_dump -U <user> -h localhost -d <db-name> > $BACKUP_PATH/$BACKUP_FILE


mkdir -p $PERSISTENT_BACKUP_PATH
mv $BACKUP_PATH/$BACKUP_FILE $PERSISTENT_BACKUP_PATH
cd $PERSISTENT_BACKUP_PATH
/usr/bin/gpg --batch --yes --passphrase <encryption-key> --symmetric --cipher-algo AES256 $BACKUP_FILE
rm -f $BACKUP_FILE

# Upload to Backblaze B2
/root/.local/bin/b2 file upload $BUCKET_NAME $BACKUP_FILE.gpg $BACKUP_FILE.gpg