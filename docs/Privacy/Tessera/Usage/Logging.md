Messages are written to the logs using these rules for the log level:

* `ERROR`: system failures or situations that require some action to ensure correct operation of the system.
* `WARN`: notifications that don't require immediate action or that are indications that a transaction failed.
* `INFO`: information message to allow investigation of issues or to provide reassurance that the system is operating correctly. 
* `DEBUG`: very verbose logging to assist with investigation of issues

The log level is written out in uppercase as part of the log message, this can be used for alert monitoring.

## Errors
Below is a non-exhaustive list of error messages and suggested actions. Braces '{}' indicate where further detail of the root cause is logged as part of the message.

<table>
<tr>
    <th>Message</th>
    <th>Cause</th>
</tr>
<tr>
    <td>Error decoding message: {error details}</td>
    <td>Invalid base64 in privateFrom/privateFor from Quorum or in tx hash for resend<br><b>Action:</b> <em>Sender needs to provide valid base64</em></td>
</tr>
<tr>
    <td>Error occurred: {error details} Root cause: {root cause}</td>
    <td>Generated for a variety of reasons:
        <ul>
            <li> Invalid content in message, e.g. <pre>curl -X POST "http://localhost:9001/push" \ <br>    -H "accept: application/json" \<br>    -H "Content-Type: application/octet-stream" \ <br>    -d "[ \"a garbage string\"]"</pre></li>
            <li> Could not send message to peer, e.g. <pre>"Root cause: Unable to push payload to recipient url<br>http://localhost:9001/"</pre></li>
        </ul>
        <b>Action:</b> <em>depends on the root cause in the log message</em>
    </td>
</tr>
<tr>
    <td>Enclave unavailable: {error details}</td>
    <td><b>Action:</b> <em>user needs to check why enclave is unavailable (look in log file for enclave)</em></td>
</tr>
<tr>
    <td>Entity not found: {error details}</td>
    <td>API request received against <code>q2tserver/transaction/{key}</code> where key is not a tx hash in the DB</td>
</tr>
<tr>
    <td>Entity not found:{error details}</td>
    <td>Thrown if endpoint doesn't exist on that API, e.g.<pre>curl -s http://localhost:9001/invalidendpoint</pre></td>
</tr>
<tr>
    <td>Security exception {followed by exception message, like "java.lang.SecurityException: No key found for url 127.1.1.1"}</td>
    <td>Thrown if <code>enableRemoteKeyValidation: true</code> and partyinfo request received from a URL for which we don't hold a public key (i.e. potentially a malicious party).<br>Note: if key validation enabled then this exception will be thrown during startup whilst the nodes exchange key information.</td>
</tr>
<tr>
    <td>ERROR c.q.t.a.e.DefaultExceptionMapper - HTTP 400 Bad Request</td>
    <td>Logged if received message is corrupt/incorrectly formatted, e.g.<pre>curl -X POST "http://localhost:9001/resend" \<br>    -H "accept: text/plain" \<br>    -H "Content-Type: application/json" \<br>    -d "{ \"some rubbish\" }"</pre></td>
</tr>
<tr>
    <td>Error while reading secret from file</td>
    <td>Unable to read the secret key (password) from file specified by <code>TESSERA_CONFIG_SECRET</code><br>
        <b>Action:</b> <em>ensure the secret key file config is correct, and file can be read</em>
    </td>
</tr>
<tr>
    <td>unable to initialize encryption facade {error details}</td>
    <td>Unable to initialise elliptical curve encryption. Logged error message will give further details<br>
        <b>Action:</b> <em>check configuration properties</em>
    </td>
</tr>
<tr>
    <td>unable to generate shared secret {error details}</td>
    <td>Unable to generate shared secret for elliptical curve encryption. Logged error message will give further details.<br>
        <b>Action:</b> <em>check configuration properties</em>
    </td>
</tr>
<tr>
    <td>unable to perform symmetric encryption {error details}</td>
    <td>Unable to encrypt data. Logged error message will give further details.<br>
        <b>Action:</b> <em>check configuration properties</em>
    </td>
</tr>
<tr>
    <td>unable to perform symmetric decryption {error details}</td>
    <td>Unable to decrypt data. Logged error message will give further details.<br>
        <b>Action:</b> <em>check configuration properties</em>
    </td>
</tr>
<tr>
    <td>Error when executing action {action type}, exception details: {error details}</td>
    <td> Unable to start Influx DB. Logged error message will give further details<br>
        <b>Action:</b> <em>check configuration properties</em>
    </td>
</tr>
<tr>
    <td>Error creating bean with name 'entityManagerFactory'</td>
    <td>Unable to create connection to database due to failure to decrypt the DB password using the supplied secret key<br>
        <b>Action:</b> <em>ensure that the correct value is supplied for the secret key</em>
    </td>
</tr>
<tr>
    <td>Config validation issue: {property name} {error details}</td>
    <td>Invalid configuration detected<br>
        <b>Action:</b> <em>correct the configuration of the named property.</em>
    </td>
</tr>
<tr>
    <td>Invalid json, cause is {error details}</td>
    <td>Invalid json in the configuration file<br>
        <b>Action:</b> <em>check the configuration file for mistakes.</em>
    </td>
</tr>
<tr>
    <td>Configuration exception, cause is {error details}</td>
    <td>Invalid data in the configuration file<br>
        <b>Action:</b> <em>check the configuration file for mistakes.</em>
    </td>
</tr>
<tr>
    <td>CLI exception, cause is {error details}</td>
    <td>Invalid command line<br>
        <b>Action:</b> <em>The error details will give further information regarding the action to be taken.</em>
    </td>
</tr>
</table>

## Warnings
Below is a list of warning messages and possible causes. Braces '{}' indicate where further detail of the root cause is logged as part of the message.

<table>
<tr>
    <th>Message</th>
    <th>Cause</th>
</tr>
<tr>
    <td>Public key {publicKey} not found when searching for private key</td>
    <td>The key in a transaction is not recognised, i.e. it is not the public key of a known participant node.</td>
</tr>
<tr>
    <td>Recipient not found for key: {public key}</td>
    <td>An unrecognised participant is specified in a transaction.<br>No action needed.</td>
</tr>
<tr>
    <td>Unable to unmarshal payload</td>
    <td>A received message is corrupt, or incorrectly formatted</td>
</tr>
<tr>
    <td>Remote host {remote host name} with IP {remote host IP} failed whitelist validation</td>
    <td>Logged if whitelist validation is enabled and the remote host is not in the whitelist.<br>
        <b>Action:</b> <em>either this is a malicious connection attempt, or mis-configuration</em>
    </td>
</tr>
<tr>
    <td>Ignoring unknown/unmatched json element: {element tag name}</td>
    <td>An unrecognised element has been found in the config file.<br>
        <b>Action:</b> <em>remove or correct the config file entry</em>
    </td>
</tr>
<tr>
    <td>Not able to find or read any secret for decrypting sensitive values in config</td>
    <td>Secret key (password) could not be read from console or password file (see <code>TESSERA_CONFIG_SECRET in docs</code>).<br>
        <b>Action:</b> <em>correction needed for the secret key or the file access permission</em>
    </td>
</tr>
<tr>
    <td>Some sensitive values are being given as unencrypted plain text in config. Please note this is NOT recommended for production environment.</td>
    <td>Self explanatory</td>
</tr>
<tr>
    <td>Not able to parse configured property. Will use default value instead</td>
    <td>Error in config file</td>
</tr>
<tr>
    <td>IOException while attempting to close remote session {error details}</td>
    <td>Only occurs on shutdown, no action needed</td>
</tr>
<tr>
    <td>Could not compute the shared key for pub {public key} and priv REDACTED</td>
    <td>Possible cause is that a public key does not match the configured cryptography algorithm.<br>
        <b>Action:</b> <em>ensure provided key is correct</em>
    </td>
</tr>
<tr>
    <td>Could not create sealed payload using shared key {shared key}</td>
    <td>Possible cause is that a public key does not match the configured cryptography algorithm.<br>
        <b>Action:</b> <em>ensure provided key is correct</em>
    </td>
</tr>
<tr>
    <td>Could not open sealed payload using shared key {shared key}</td>
    <td>Possible cause that wrong password was given for key file decryption or making a change to the values in the keyfile so that the password no longer works.<br>
        <b>Action:</b> <em>ensure that password is correct for the keyfile</em>
    </td>
</tr>
<tr>
    <td>Unable to generate a new keypair!</td>
    <td>Internal error - potentially an issue with jnacl dependency</td>
</tr>
<tr>
    <td>Exception thrown : {exception message} While starting service {service name}</td>
    <td>Internal error - failed to start a service</td>
</tr>
<tr>
    <td>Invalid key found {remote host url} recipient will be ignored</td>
    <td>Remote key validation check failed.<br>No action needed, however it is a possible indication of a malicious node</td>
</tr>
<tr>
    <td>Push returned status code for peer {remote peer url} was {status code}</td>
    <td>The peer rejected a transaction 'push' request.<br>
        <b>Action:</b> <em>check logs on peer to see why it failed</em>
    </td>
</tr>
<tr>
    <td>PartyInfo returned status code for peer{remote peer url} was {status code}</td>
    <td>The peer rejected a partyInfo request.<br>
        <b>Action:</b> <em>check logs on peer to see why it failed</em></td>
</tr>
<tr>
    <td>Unable to resend payload to recipient with public key {public key}, due to {error details}</td>
    <td>The peer rejected a transaction push request during a resend operation.<br>
        <b>Action:</b> <em>check reason message, or logs on peer to see why it failed</em>
    </td>
</tr>
<tr>
    <td>Attempt is being made to update existing key with new url. Please switch on remote key validation to avoid a security breach.</td>
    <td>Self explanatory</td>
</tr>
<tr>
    <td>Failed to connect to node {remote node url}, due to {error details}</td>
    <td>A remote node refused partyinfo request. Can occur if:
        <ul>
            <li>remote node is not running</li>
            <li>remote node doesn't recognise this node's public key</li>
            <li>remote node doesn't have this node's IP registered against a key</li>
            <li>etc</li>
        </ul>
        Can also be expected to occur when nodes are shutdown/restarted, so not necessarily an error.
    </td>
</tr>
<tr>
    <td>Failed to connect to node {remote node url} for partyInfo, due to {error details}</td>
    <td>A node failed partyInfo request during resend to peer.<br>
        <b>Action:</b> <em>check reason message, or logs on peer to see why it failed</em>
    </td>
</tr>
<tr>
    <td>Failed to make resend request to node {remote node url} for key {public key}, due to {error details}</td>
    <td>Peer communication failed during '/resend' request.<br>
        <b>Action:</b> <em>check reason message, or logs on peer to see why it failed</em> 
    </td>
</tr>
</table>

!!! Note 
    Some messages will be rearranged to correct logging levels in our next release.


## To change the default log level

The level of logging is controlled by the logback configuration file. The default file packaged with Tessera can be seen [here](https://github.com/jpmorganchase/tessera/blob/master/tessera-dist/tessera-launcher/src/main/resources/logback.xml).

To specify a different logging configuration, pass a customised logback file on the command line using:
`-Dlogback.configurationFile=/path/to/logback.xml`
