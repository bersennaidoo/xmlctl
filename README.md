# xmlctl

This application uploads, reads, parses and displays the parsed xml file

## DOCUMENTATION
The models in pkg/docs tells you what the application. 

The schema in pkg/schema is for the database table.

Create database in postgresql and apply schema from ./pkg/schema.

Application configuration for database is in ./xmlctl.conf.toml

The binary for application is in bin directory for linux and windows.
```
./pkg/docs
./pkg/schema
./xmlctl.conf.toml
./bin/xmlctl
./bin/xmlctl.exe
```

## USAGE

#### UPLOAD 
```
./bin/xmlctl upload [-f|--filename] book-store.xml [-n|--name] book1
```

#### READ
```
./bin/xmlctl read [-n|--name] book1
```
