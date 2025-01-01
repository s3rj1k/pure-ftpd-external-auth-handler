# ftp-auth-handler

[*External Authentication module*](https://download.pureftpd.org/pub/pure-ftpd/doc/README.Authentication-Modules) for pure-ftpd.

## Usage
 - Create table inside MySQL database using this scheme: `files/remote-db.sql`
 - Create config file `/etc/ftp-auth-handler.yaml` using this example: `files/ftp-auth-handler.yaml.example`
 - Start pure-authd: `pure-authd -B -s /var/run/pure-ftpd/pure-ftpd.sock -r /usr/local/bin/ftp-auth-handler`
 - Start pure-ftpd (pure-ftpd-mysql): `pure-ftpd-mysql --allowdotfiles --altlog=clf:/var/log/pure-ftpd/transfer.log --chrooteveryone --clientcharset=utf-8 --createhomedir --customerproof --daemonize --dontresolve --fscharset=utf-8 --ipv4only --login=extauth:/var/run/pure-ftpd/pure-ftpd.sock --maxclientsnumber=10000 --maxclientsperip=50 --maxidletime=120 --minuid=1000 --noanonymous --notruncate --pidfile=/var/run/pure-ftpd/pure-ftpd.pid --verboselog`
 - Create local DB: `ftp-auth-handler -get-user-accounts`

Copyright (c) 2018-2025 s3rj1k <evasive.gyron@gmail.com>
