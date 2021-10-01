# mtss-cli

> Usage

```bash
This cli app permit to fetch data from MTSS and insert on redis or MongoDB.this is Mtss ctl

Usage:
  mtssctl [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  fetch       Fetch data from API rest client to an specific db (redis/mongodb[unimplemented]/postgresql)
  help        Help about any command
  insert      Insert Data from API rest client
  version     Print the version number of mtssctl

Flags:
  -h, --help   help for mtssctl

Use "mtssctl [command] --help" for more information about a command.```


> Postgres tips

```bash
# login
sudo -i -u qwerty psql
# create database
create database mtss;
# select elements using json properties
select id,job->'organismo' as entity from mtss_jobs limit 2;
```