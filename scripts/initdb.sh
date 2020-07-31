#!/bin/bash

#variables- move it to config file
rootUser="postgres"
newUser="facilEspanolUser"
password="facilEspanolPass"
database="facilEspanolDb"

psql -U $rootUser -c "CREATE DATABASE \"$database\""
psql -U $rootUser -c "CREATE USER \"$newUser\" with encrypted password '$password'"
