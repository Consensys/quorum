## Migration Utilities
Two utilities are included to help with migrating existing Constellation configurations and datastores for use with Tessera.  These utilites are included in the Tessera project and are available for use after building Tessera with Maven.

A full migration workflow would be as follows:

1. Shut down the Constellation/Quorum nodes
2. Perform [database migration](#data-migration)
3. Perform [configuration migration](#configuration-migration)
4. Start Tessera/Quorum nodes


## Data Migration
This utility migrates a Constellation datastore (BerkeleyDB or directory/file storage) to a Tessera compatible one (H2, SQLITE).  By default Tessera uses an H2 database, however alternatives can be configured.  Refer [DDLs](https://github.com/jpmorganchase/tessera/tree/master/ddls/create-table) for help with defining with other databases.

To make running the utility commands simpler, you can first create an `alias`:

```
alias tessera-data-migration="java -jar /path/to/tessera/data-migration/target/data-migration-${version}-cli.jar"
```

CLI help can be accessed by running:
```
tessera-data-migration help

usage: tessera-data-migration
-exporttype <TYPE>   Export DB type i.e. h2, sqlite
-inputpath <PATH>    Path to input file or directory
-outputfile <PATH>   Path to output file
-storetype <TYPE>    Store type i.e. bdb, dir
-dbuser              Set a username on the migrated database (only applies to H2)
-dbpass              Set a password for the specified user (only applies to H2)
```

#### Migrating BerkeleyDB (bdb)
To migrate a BerkeleyDB (bdb) database for use with Tessera you must first export your existing store using `db_dump`:
```
db_dump -f exported.txt c1/cn§.db/payload.db
```

Then run the following command to perform the migration:
```
tessera-data-migration -storetype bdb -inputpath exported.txt -dbuser <username> -dbpass <password> -outputfile <PATH> -exporttype <TYPE>
```

#### Migrating Directory/File (dir) storage
For dir storage: 
```
tessera-data-migration -storetype dir -inputpath /path/to/dir -dbuser <username> -dbpass <password> -outputfile <PATH> -exporttype <TYPE>
```

### Output types
To use H2 as the output storage, specify:
```
-exporttype h2 -outputfile /path/to/h2database
```

To use SQLite as the output storage, specify:
```
-exporttype sqlite -outputfile /path/to/sqlitedb
```

#### Database usernames and passwords
If you want to set a username and password on the migrated database, you must specify this using the following options:

```
-dbuser <username> -dbpass <password>
```

If you do not wish to set a username and password on the migrated database, you must explicitly say so by specifying the arguments without parameters, i.e.

```
-dbuser -dbpass
```

Note also that even though SQLite does not have the concept of usernames and passwords, you must still specify at least the empty configuration.


#### After migration
The output file should then be placed in a location of your choosing that corresponds to the location specified in the configuration file (without any file extension), i.e.

```
"jdbc": {
    "url": "jdbc:h2:./c1/migratedfile;MODE=Oracle;TRACE_LEVEL_SYSTEM_OUT=0"
}
```

Note: the migrated database is migrated without user credentials, so if using the file directly then the username and password should not be specified in the configuration.

The Constellation files are no longer used, and can be cleaned up or left alone.


## Configuration Migration
This utility will generate a Tessera compatible `.json` format configuration file from an existing Constellation `.toml` configuration file.  The `.json` file will be saved locally to be used when running Tessera.  Individual configuration parameters can be overridden during the migration process if required.

To make running the utility commands simpler, you can first create an `alias`:

```
alias tessera-config-migration="java -jar /path/to/tessera/config-migration/target/config-migration-${version}-cli.jar"
```

Most of the Constellation configuration command line parameters are supported.  For details of the Constellation configuration see the [Constellation documentation](../../Constellation/Constellation).

To see the CLI help which provides details on overriding specific configuration items from a `.toml` file, run:
```
tessera-config-migration help
```

To migrate a `.toml` file to `.json` with no overrides, run:
```
tessera-config-migration --tomlfile="/path/to/constellation-config.toml"
```

By default, the generated `.json` config will be printed to the console and saved to `./tessera-config.json`.  To save to another location/with a different filename use the `--outputfile <PATH>` CLI option.

#### Note about `ipwhitelist`
Unlike Constellation, Tessera does not use a separate `ipwhitelist`.  If `useWhiteList` is set to `true` in the `.json` config then the `peers` list will be used as the whitelist.  

If `ipwhitelist` is provided in an existing `.toml` config file then this will only be used to set `useWhiteList: true`; any nodes included in this list will not be added by default to the Tessera config.  Make sure to add any nodes that were only included in `ipwhitelist` to `peers` after using the utility.

#### Validation
Validation is applied to the generated config. Messages will be printed to the terminal if the validation identifies issues.  For example, if a `hostname` is not provided then the following message will be printed:
```
Warning: may not be null on property serverConfig.hostName
```
Any validation violations will have to be addressed before the config can be used with Tessera.
