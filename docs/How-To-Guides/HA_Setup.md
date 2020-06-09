# Quorum HA Setup

Quorum architecture allows for true HA setup end to end for heightened availability for various input output operations. Although this increase the footprint of each node, the cost offset is compensate by low to zero downtime & horizontal scalability. In this section we will go through the configuration and setup in detail:

**WARNING**: Below HA setup hasn't been extensively tested and more for testnet as such should not be considered suitable for production use (yet).

## Quorum Node Configuration Requirements:

- Two or more Quorum Nodes serve as one client node.
- The inbound RPC requests from clients will be load balanced to one of these Quorum nodes.
- These nodes will need to share same key for transaction signing and should have shared access to key store directory.
- These nodes need to share the same private state i.e., they connect to same Tessera node(s). This is done by proxy running on each Quorum node listening on local ipc file and directing request to Tessera Q2T http.

## Tessera Node Configuration Requirements:

- Separate Proxy server to redirect/mirror requests to two or more Tessera nodes 
- Two or more Tessera Nodes serve as Privacy manager for Client Quorum node.
- These nodes share same public/private key pair (stored in password protected files or external vaults)
- In the server config, the bindingAddress should be the local addresses (their real addresses), but 'advertisedAddress' (serverAddress) needs to be configured to be the proxy
- Add DB replication or mirroring for Tessera private data store and the JDBC connection string to include both Primary DB and DR DB connections to facilitate auto switchover on failure.


??? info Quorum HA Setup
    ![Quorum Tessera Privacy Flow](https://github.com/jpmorganchase/tessera/raw/master/Tessera%20Privacy%20flow.jpeg)


## Example Setup using nginx Proxy setup

### Proxy Setup on both Quorum nodes

proxy_next_upstream in proxy for q2t endpoints to avoid rejections for valid requests

 
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
    
 

### Proxy server setup

 
    load_module /usr/lib/nginx/modules/ngx_stream_module.so;

    error_log /home/ubuntu/nginx-error.log;
    events { }
    stream {
    #Quorum json-rpc
        upstream quorum {
                server ec2-35-178-250-190.eu-west-2.compute.amazonaws.com:22000;
                server ec2-35-177-214-194.eu-west-2.compute.amazonaws.com:22000 backup;
        }
        server {
                listen 22000;
                proxy_pass quorum;
        }
      }
    http {

        # Third-party server
        upstream thirdparty {
                server ec2-35-178-250-190.eu-west-2.compute.amazonaws.com:9081;
                server ec2-35-177-214-194.eu-west-2.compute.amazonaws.com:9081;
        }
        server {
                listen 9081;
                location / {
                        proxy_next_upstream error timeout http_404 non_idempotent;
                        proxy_pass http://thirdparty;
                }
        }
        # Peer-to-peer server
        upstream p2p {
                server ec2-35-178-250-190.eu-west-2.compute.amazonaws.com:9001;
                server ec2-35-177-214-194.eu-west-2.compute.amazonaws.com:9001;
        }
        upstream p2p-mirror {
                server ec2-35-178-250-190.eu-west-2.compute.amazonaws.com:9001;
                server ec2-35-177-214-194.eu-west-2.compute.amazonaws.com:9001 backup;
        }
        server {
                listen 9001;

                location /resend {
                        proxy_pass http://p2p/resend;
                }
                location /push {
                        proxy_pass http://p2p/push;
                }
                location /partyinfo {
                        mirror /partyinfo-mirror;
                        proxy_pass http://p2p-mirror/partyinfo;
                }
                location /partyinfo-mirror {
                        internal;
                        proxy_pass http://ec2-35-177-214-194.eu-west-2.compute.amazonaws.com:9001/partyinfo;
                }
                location /partyinfo/validate {
                        proxy_pass http://p2p/partyinfo/validate;
                }
                location /upcheck {
                        proxy_pass http://p2p/upcheck;
      }}}




