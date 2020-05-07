# CLI config overrides 

The `-o, --override` option is used to define overrides as key/value pairs.  The key is the json path of the field to be overwritten.

For example, given `configfile.json`:
```json
{
    ...,
    "jdbc" : {
        "username" : "sa",
        "password" : "",
        "url" : "jdbc:h2:/path/to/db1;MODE=Oracle;TRACE_LEVEL_SYSTEM_OUT=0",
        "autoCreateTables" : true,
        "fetchSize" : 0
    },    
    "peer" : [ 
        {
            "url" : "http://127.0.0.1:9001"
        }
    ]
}
```

The command:
```bash
tessera --configfile configfile.json -o jdbc.username=username-override --override peer[1].url=http://peer-override:9001
```

will start Tessera with the following effective config:
```json
{
    ...,
    "jdbc" : {
        "username" : "username-override",
        "password" : "",
        "url" : "jdbc:h2:/path/to/db1;MODE=Oracle;TRACE_LEVEL_SYSTEM_OUT=0",
        "autoCreateTables" : true,
        "fetchSize" : 0
    },    
    "peer" : [ 
        {
            "url" : "http://127.0.0.1:9001"
        },
        {
            "url" : "http://peer-override:9001"
        }
    ]
}
```
