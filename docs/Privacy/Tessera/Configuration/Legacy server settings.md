**Important**

Legacy server settings were part of Tessera v0.5, v0.6 and v0.7. They have deprecated for, but still work with, v0.8 and v0.9 and may be removed for any future versions.

---

The server settings are defined as following:

```
"server": {
  "hostName":       "<Hostname to advertise. Includes url scheme but not port e.g. http://myhost.com>",
  "port":           "<Port that will be used to expose Tessera services>",
  "bindingAddress": "<Full interface to bind. Includes URL scheme and port. e.g. http://0.0.0.0:9999>",
  "sslConfig":      <...>,
  "influxConfig":   <...>
}
```

<br>
<br>

If the address to advertise keys on is the same as the address you wish to bind to, then the `bindingAddress` may be omitted, for example:
```
"server": {
  "hostName":     "http://myhost.com",
  "port":         9999,
  "sslConfig":    <...>,
  "influxConfig": <...>
}
```

---

**InfluxDB Config**

Configuration details to allow Tessera to record monitoring data to a running InfluxDB instance.
```
"influxConfig": {
  "hostName": "[Hostname of Influx instance]",
  "port": "[Port of Influx instance]",
  "pushIntervalInSecs": "[How often to push data to InfluxDB]",
  "dbName": "[Name of InfluxDB]"
}
```

### Unix socket file
Path to the Unix domain socket file used to communicate between Quorum and Tessera.
```
"unixSocketFile" : "/path/to/socketfile.ipc"
```
