#!/bin/bash

sudo -u postgres -i psql -c '
\set autocommit on;
create database "facilEspanolDb";
create user "facilEspanolUser" with encrypted password '"'"'facilEspanolPass'"'"';
'

