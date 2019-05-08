CLI options can be used to add to, or override, configuration defined in a `configfile`.  

Standard Tessera CLI options are prefixed with a single hyphen (e.g. `-configfile <PATH>`), whilst the config override options are prefixed with a double hyphen (e.g. `--alwaysSendTo <STRING[]...>`).  Use `tessera help` to see a complete list of CLI options.  

If a config value is included in both the `configfile` and the CLI, then the CLI value will take precendence. The exceptions to this rule are the `--peer.url <STRING>` and `--alwaysSendTo <STRING[]...>` options.  Instead of overriding, these CLI options append to any peer or alwaysSendTo urls in the provided `configfile`.  For example, if the following was provided in a `configfile`:
```
{
  ...
  "peer": [
    {
      "url": "http://localhost:9001"
    }
  ],
  alwaysSendTo:[
    "giizjhZQM6peq52O7icVFxdTmTYinQSUsvyhXzgZqkE="
  ],
  ...
}
```
and Tessera was run with the following overrides:
```
tessera -configfile path/to/file --peer.url http://localhost:9002 --peer.url http://localhost:9003 --alwaysSendTo /+UuD63zItL1EbjxkKUljMgG8Z1w0AJ8pNOR4iq2yQc= --alwaysSendTo UfNSeSGySeKg11DVNEnqrUtxYRVor4+CvluI8tVv62Y=
```
then Tessera will be started with the following equivalent configuration:
```
{
  ...
  "peer": [
    {
      "url": "http://localhost:9001"
    },
    {
      "url": "http://localhost:9002"
    },
    {
      "url": "http://localhost:9003"
    }
  ],
  alwaysSendTo:[
    "giizjhZQM6peq52O7icVFxdTmTYinQSUsvyhXzgZqkE=",
    "/+UuD63zItL1EbjxkKUljMgG8Z1w0AJ8pNOR4iq2yQc="
    "UfNSeSGySeKg11DVNEnqrUtxYRVor4+CvluI8tVv62Y="
  ],
  ...
}
```
As demonstrated in this example, in certain cases multiple values can be provided by repeating the CLI option.  This is supported for the `peer.url`, `alwaysSendTo`, `server.sslConfig.serverTrustCertificates` and `server.sslConfig.clientTrustCertificates` options.  
