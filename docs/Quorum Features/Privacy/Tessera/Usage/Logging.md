Messages are written to the logs using these rules for the log level:

* ERROR: system failures or situations that require some action to ensure correct operation of the system.
* WARN: notifications that don't require immediate action or that are indications that a transaction failed.
* INFO: information message to allow investigation of issues or to provide reassurance that the system is operating correctly. 
* DEBUG: very verbose logging to assist with investigation of issues

The log level is written out in uppercase as part of the log message, this can be used for alert monitoring.

## Errors
Below is a non-exhaustive list of error messages and suggested actions. Braces '{}' indicate where further detail of the root cause is logged as part of the message.

Message | Cause
------- | -----
Error decoding message: {error details} | Invalid base64 in privateFrom/privateFor from Quorum or in tx hash for resend.<br>Action: sender needs to provide valid base64.<br>[See Note 1](#Note1)
Error occurred: {error details} Root cause: {root cause} | Generated for a variety of reasons:<br>- Invalid content in message, e.g.<br>    curl -X POST "http://localhost:9001/push" -H "accept: application/json" -H "Content-Type: application/octet-stream" -d "[ \"a garbage string\"]"<br>- Could not send message to peer, e.g. "Root cause: Unable to push payload to recipient url http://localhost:9001/"<br>Action to take depends on the root cause in the log message.
Enclave unavailable: {error details} | Action: user needs to check why enclave is unavailable (look in log file for enclave)
Entity not found: {error details} | API request received against q2tserver/transaction/{key} where key is not a tx hash in the DB.<br>[See Note 1](#Note1)
Entity not found:{error details} | Thrown if endpoint doesn't exist on that API, e.g.<br>    curl -s http://localhost:9001/invalidendpoint<br>[See Note 1](#Note1)
Security exception {followed by exception message, like "java.lang.SecurityException: No key found for url 127.1.1.1"} | Thrown if 'enableRemoteKeyValidation' is true and partyinfo request received from a URL for which we don't hold a public key (i.e. potentially a malicious party).<br>Note: if key validation enabled then this exception will be thrown during startup whilst the nodes exchange key information.
ERROR c.q.t.a.e.DefaultExceptionMapper - HTTP 400 Bad Request | Logged if received message is corrupt/incorrectly formatted, along with the warning below, e.g.<br>    curl -X POST "http://localhost:9001/resend" -H "accept: text/plain" -H "Content-Type: application/json" -d "{ \"some rubbish\" }"<br>[See Note 1](#Note1)
Error while reading secret from file | Unable to read the secret key (password) from file specified by TESSERA_CONFIG_SECRET.<br>Action: ensure the secret key file config is correct, and file can be read
unable to initialize encryption façade {error details} | Unable to initialise elliptical curve encryption. Logged error message will give further details.<br>Action: check configuration properties
unable to generate shared secret {error details} | Unable to generate shared secret for elliptical curve encryption. Logged error message will give further details.<br>Action: check configuration properties
unable to perform symmetric encryption {error details} | Unable to encrypt data. Logged error message will give further details.<br>Action: check configuration properties
unable to perform symmetric decryption {error details} | Unable to decrypt data. Logged error message will give further details.<br>Action: check configuration properties
Error when executing action {action type}, exception details: {error details} | Unable to start Influx DB. Logged error message will give further details.<br>Action: check configuration properties
Error creating bean with name 'entityManagerFactory' | Unable to create connection to database due to failure to decrypt the DB password using the supplied secret key.<br>Action: ensure that the correct value is supplied for the secret key.
Config validation issue: {property name} {error details} | Invalid configuration detected.<br>Action: correct the configuration of the named property.
Invalid json, cause is {error details} | Invalid json in the configuration file.<br>Action: check the configuration file for mistakes.
Configuration exception, cause is {error details} | Invalid data in the configuration file.<br>Action: check the configuration file for mistakes.
CLI exception, cause is {error details} | Invalid command line.<br>The error details will give further information regarding the action to be taken. 

!!! Note1 
Log message will be changed to WARN in next release version.

## Warnings
Below is a list of warning messages and possible causes. Braces '{}' indicate where further detail of the root cause is logged as part of the message.

Message | Cause
------- | -----
Public key {publicKey} not found when searching for private key | The key in a transaction is not recognised, i.e. it is not the public key of a known participant node.
Recipient not found for key: {public key} | An unrecognised participant is specified in a transaction.<br>No action needed.
Unable to unmarshal payload | A received message is corrupt, or incorrectly formatted
Remote host {remote host name} with IP {remote host IP} failed whitelist validation | Logged if whitelist validation is enabled and the remote host is not in the whitelist.<br>Action: either this is a malicious connection attempt, or mis-configuration.
Ignoring unknown/unmatched json element: {element tag name} | An unrecognised element has been found in the config file.<br>Action: remove or correct the config file entry
Not able to find or read any secret for decrypting sensitive values in config | Secret key (password) could not be read from console or password file (see TESSERA_CONFIG_SECRET in docs).<br>Action: correction needed for the secret key or the file access permission
Some sensitive values are being given as unencrypted plain text in config. Please note this is NOT recommended for production environment. | Self explanatory
Not able to parse configured property. Will use default value instead. | Error in config file.
IOException while attempting to close remote session {error details} | Only occurs on shutdown, no action needed
Could not compute the shared key for pub {public key} and priv REDACTED | Possible cause is that a public key does not match the configured cryptography algorithm.<br>Action: ensure provided key is correct.<br>[See Note 2](#Note2)
Could not create sealed payload using shared key {shared key} | Possible cause is that a public key does not match the configured cryptography algorithm.<br>Action: ensure provided key is correct.<br>[See Note 2](#Note2)
Could not open sealed payload using shared key {shared key} | Possible cause that wrong password was given for key file decryption or making a change to the values in the keyfile so that the password no longer works.<br>Action: ensure that password is correct for the keyfile.<br>[See Note 2](#Note2)
Unable to generate a new keypair! | Internal error - potentially an issue with jnacl dependency.<br>[See Note 2](#Note2)
Exception thrown : {exception message} While starting service {service name} | Internal error - failed to start a service.<br>[See Note 2](#Note2)
Invalid key found {remote host url} recipient will be ignored | Remote key validation check failed.<br>No action needed, however it is a possible indication of a malicious node.<br>[See Note 3](#Note3)
Push returned status code for peer {remote peer url} was {status code} | The peer rejected a transaction 'push' request.<br>Action: check logs on peer to see why it failed
PartyInfo returned status code for peer{remote peer url} was {status code} | The peer rejected a partyInfo request.<br>Action: check logs on peer to see why it failed
Unable to resend payload to recipient with public key {public key}, due to {error details} | The peer rejected a transaction push request during a resend operation.<br>Action: check reason message, or logs on peer to see why it failed
Attempt is being made to update existing key with new url. Please switch on remote key validation to avoid a security breach. | Self explanatory
Failed to connect to node {remote node url}, due to {error details} | A remote node refused partyinfo request. Can occur if:<br>- remote node is not running<br>- remote node doesn't recognise this node's public key<br>- remote node doesn't have this node's IP registered against a key<br>- etc<br>Can also be expected to occur when nodes are shutdown/restarted, so not necessarily an error.
Failed to connect to node {remote node url} for partyInfo, due to {error details} | A node failed partyInfo request during resend to peer.<br>Action: check reason message, or logs on peer to see why it failed
Failed to make resend request to node {remote node url} for key {public key}, due to {error details} | Peer communication failed during '/resend' request.<br>Action: check reason message, or logs on peer to see why it failed

!!! Note 2 
Log message will be changed to ERROR in next release version.
!!! Note 3
Log message will be changed in next release to give key and url.

## To change the default log level

The level of logging is controlled by the logback configuration file. The default file packaged with Tessera can be seen [here](https://github.com/jpmorganchase/tessera/blob/master/tessera-dist/tessera-launcher/src/main/resources/logback.xml).

To specify a different logging configuration, pass a customised logback file on the command line using:
`-Dlogback.configurationFile=/path/to/logback.xml`
