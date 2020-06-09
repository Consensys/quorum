# Quorum HA Setup

**WARNING**: The HA setup hasn't been extensively testing and more for test net as such should not be considered suitable for production use (yet).

Quorum architecture allows for true HA setup end to end for heightened availability for various input output operations. Although this increase the footprint of each node, the cost offset is compensate by low to zero downtime & horizontal scalability. In this section we will go through the configuration and setup in detail:

## Quorum Node Configuration in HA mode

- Two or more Quorum Nodes serve as one client node.
- The inbound RPC requests from clients will be load balanced to one of these Quorum nodes.
- These nodes will need to share same key for transaction signing and should have shared access to key store directory.
- These nodes need to share the same private state i.e., they connect to same Tessera node(s). This is done by proxy running on each Quorum node listening on local ipc file and directing request to Tessera Q2T http.

## Tessera Node Configuration in HA mode

- Two or more Tessera Nodes serve as Privacy manager for Client Quorum node.
- These nodes share same public/private key pair
- In the server config, the bindingAddress should be the local addresses (their real addresses), but 'advertisedAddress' (serverAddress) needs to be configured to be the proxy
- Add DB replication or mirroring for Tessera private data store and the JDBC connection string to include both Primary DB and DR DB connections to facilitate auto switchover on failure.

## Example Setup

### Proxy Setup on both nodes

```
load_module /usr/lib/nginx/modules/ngx_stream_module.so;
error_log /home/ubuntu/nginx-error.log;
events { }
http {

        # Quorum-to-Tessera http
        upstream q2t {
                server ec2-35-178-250-190.eu-west-2.compute.amazonaws.com:9091;
                server ec2-35-177-214-194.eu-west-2.compute.amazonaws.com:9091;
        }

        server {
                listen unix:/home/ubuntu/tm.ipc;
                location / {
                        proxy_next_upstream error timeout http_404 non_idempotent;
                        proxy_pass http://q2t;
                }
        }
}
```






