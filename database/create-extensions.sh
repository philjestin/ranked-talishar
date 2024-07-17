#!/usr/bin/env bash

echo "enabling uuid-ossp on database $POSTGRES_DB"
psql -U $POSTGRES_USER --dbname="$POSTGRES_DB" <<-'EOSQL'
  create extension if not exists uuid-ossp;
EOSQL
echo "finished with exit code $?"