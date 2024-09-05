#!/bin/sh

set -u

if [ -z "${DATABASE_URL:-}" ]; then
    echo "Database URL is not available. Please set it. Exiting...."
    exit 1
fi

db_url_prefix="sqlite:"
db_dir=$(dirname ${DATABASE_URL#"$db_url_prefix"})

if [ ! -d $db_dir ]; then
    echo "Database directory [ ${db_dir} ] does not exist. Creating...."
    mkdir -p "$db_dir"
fi

goose -dir="./sql/migrations" sqlite3 ../data/habits.db up

echo "....Finished provisioning the database."

exit 0