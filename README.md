# xmlctl

This application uploads, reads, parses and displays the parsed xml file

## DOCUMENTATION
The models tells you what the application does and schema for database table

Create database in postgresql and apply schema from ./pkg/schema.

Configuration for database is in ./xmlctl.conf.toml
```
./pkg/docs
./pkg/schema
./xmlctl.conf.toml
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
