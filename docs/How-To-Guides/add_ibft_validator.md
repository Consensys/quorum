# Adding and removing IBFT validators

Over the lifetime of an IBFT network, validators will need to be added and removed as authorities change. 
Here we will showcase adding a new validator to an IBFT network, as well as removing an existing one.

## Adding a node to the validator set

Adding a node to the IBFT validator set is relatively easy once a node is part of the network.
It does not matter whether the node is already online or not, as the process to add the new node as a validator only 
needs the *existing* validators.

!!! warning
    If you are adding multiple validators before they are brought online, make sure you don't go over the BFT limit and cause the chain to stop progressing.

Adding a new validator requires that a majority of existing validators propose the new node to be added. This is 
achieved by calling the `propose` RPC method with the value `true` and replacing the address to your required one:

    ```bash
    $ geth attach /qdata/dd/geth.ipc
    
    > istanbul.propose("0xb131288f355bc27090e542ae0be213c20350b767", true);
    null
    ```
    
This indicates that the current node wishes to add address `0xb131288f355bc27090e542ae0be213c20350b767` as a new 
validator.

### Example

You can find the resources required to run the examples in the 
[quorum-examples](https://github.com/jpmorganchase/quorum-examples/tree/master/examples/ibft_validator_set_changes) 
repository. 

1. The examples use `docker-compose` for the container definitions. If you are following along by copying the commands 
   described, then it is important to set the project name for Docker Compose, or to remember to change the prefix for 
   your directory. See [Docker documentation](https://docs.docker.com/compose/reference/envvars/#compose_project_name) 
   for more details.
   
   To set the project name, run the following:
   ```bash
   $ export COMPOSE_PROJECT_NAME=addnode
   ```
   
2. Bring up the network, which contains 7 nodes, of which 6 are validators.

    ```bash
    $ docker-compose -f ibft-6-validators.yml up
    ```
   
   We will be adding the 7th node as a validator. You may notice in the logs of node 7 messages along the lines of 
   `node7_1  | WARN [01-20|10:37:16.034] Block sealing failed                     err=unauthorized`. This is because 
   the node was started up with minting enabled, but doesn't have the authority to create blocks, and so throws this 
   error.
   
3. Now we need to propose node 7 as a new proposer from the existing nodes.

    !!! note
        Remember, you could do this stage before starting node 7 in your network
        
   We need a majority of existing validators to propose the new node before the changes will take effect.
   
   Lets start with node 1 and see what happens:
   
    ```bash
    # Propose node 7 from node 1
    $ docker exec -it addnode_node1_1 geth --exec 'istanbul.propose("0xb131288f355bc27090e542ae0be213c20350b767", true);' attach /qdata/dd/geth.ipc
    null
   
    # Wait about 5 seconds, and then run:
    $ docker exec -it addnode_node1_1 geth --exec 'istanbul.getSnapshot();' attach /qdata/dd/geth.ipc
    {
      epoch: 30000,
      hash: "0xf814863d809ce3a683ee0a2197b15a8152d2696fc9c4e47cd82d0bd5cdaa3e45",
      number: 269,
      policy: 0,
      tally: {
        0xb131288f355bc27090e542ae0be213c20350b767: {
          authorize: true,
          votes: 1
        }
      },
      validators: ["0x6571d97f340c8495b661a823f2c2145ca47d63c2", "0x8157d4437104e3b8df4451a85f7b2438ef6699ff", "0xb912de287f9b047b4228436e94b5b78e3ee16171", "0xd8dba507e85f116b1f7e231ca8525fc9008a6966", "0xe36cbeb565b061217930767886474e3cde903ac5", "0xf512a992f3fb749857d758ffda1330e590fa915e"],
      votes: [{
          address: "0xb131288f355bc27090e542ae0be213c20350b767",
          authorize: true,
          block: 268,
          validator: "0xd8dba507e85f116b1f7e231ca8525fc9008a6966"
      }]
    }
    ```
   
   Let's break this down.
   Firstly, we proposed the address `0xb131288f355bc27090e542ae0be213c20350b767` to be added; that is what the `true` 
   parameter is for. If we had set it to `false`, that means we want to remove an existing validator with that address.
   
   Secondly, we fetched the current snapshot, which gives us an insight into the current running state of the voting.
   We can see that the new address has 1 vote under the `tally` section, and that one vote is described under the 
   `votes` section. So we know our vote was registered!
   
4. Let's run this from node 2 and see similar results:

    ```bash
    $ docker exec -it addnode_node2_1 geth --exec 'istanbul.propose("0xb131288f355bc27090e542ae0be213c20350b767", true);' attach /qdata/dd/geth.ipc
    null
   
    # Again, you may have to wait 5 - 10 seconds for the snapshot to show the vote
    $ docker exec -it addnode_node2_1 geth --exec 'istanbul.getSnapshot();' attach /qdata/dd/geth.ipc
    {
      epoch: 30000,
      hash: "0x93efcd458f3b875902a4532bb77d5e7ebb701791ea95486ecd58baf682312d74",
      number: 391,
      policy: 0,
      tally: {
        0xb131288f355bc27090e542ae0be213c20350b767: {
          authorize: true,
          votes: 2
        }
      },
      validators: ["0x6571d97f340c8495b661a823f2c2145ca47d63c2", "0x8157d4437104e3b8df4451a85f7b2438ef6699ff", "0xb912de287f9b047b4228436e94b5b78e3ee16171", "0xd8dba507e85f116b1f7e231ca8525fc9008a6966", "0xe36cbeb565b061217930767886474e3cde903ac5", "0xf512a992f3fb749857d758ffda1330e590fa915e"],
      votes: [{
          address: "0xb131288f355bc27090e542ae0be213c20350b767",
          authorize: true,
          block: 388,
          validator: "0xd8dba507e85f116b1f7e231ca8525fc9008a6966"
      }, {
          address: "0xb131288f355bc27090e542ae0be213c20350b767",
          authorize: true,
          block: 390,
          validator: "0x6571d97f340c8495b661a823f2c2145ca47d63c2"
      }]
    }
    ```
   
   True to form, we have the second vote registered!

5. Ok, let's finally vote on nodes 3 and 4.

    ```bash
    $ docker exec -it addnode_node3_1 geth --exec 'istanbul.propose("0xb131288f355bc27090e542ae0be213c20350b767", true);' attach /qdata/dd/geth.ipc
    null
   
    $ docker exec -it addnode_node4_1 geth --exec 'istanbul.propose("0xb131288f355bc27090e542ae0be213c20350b767", true);' attach /qdata/dd/geth.ipc
    null
    ```

6. Now we have a majority of votes, let's check the snapshot again:

    ```bash
    docker exec -it addnode_node1_1 geth --exec 'istanbul.getSnapshot();' attach /qdata/dd/geth.ipc
    {
      epoch: 30000,
      hash: "0xd4234184538297f71f5b7024a2e11f51f06b4f569ebd9e3644abd391b8c66101",
      number: 656,
      policy: 0,
      tally: {},
      validators: ["0x6571d97f340c8495b661a823f2c2145ca47d63c2", "0x8157d4437104e3b8df4451a85f7b2438ef6699ff", "0xb131288f355bc27090e542ae0be213c20350b767", "0xb912de287f9b047b4228436e94b5b78e3ee16171", "0xd8dba507e85f116b1f7e231ca8525fc9008a6966", "0xe36cbeb565b061217930767886474e3cde903ac5", "0xf512a992f3fb749857d758ffda1330e590fa915e"],
      votes: []
    }
    ```
   
   We can see that the votes have now been wiped clean, ready for a new round. Additionally, the address we were adding,
   `0xb131288f355bc27090e542ae0be213c20350b767` now exists within the `validators` list!
   Lastly, the `unauthorized` messages that node 7 was giving before has stopped, as it now has the authority to mint 
   blocks.

## Removing a node from the validator set

Removing a validator is very similar to adding a node, but this time we want to propose nodes with the value `false`, 
to indicate we are deauthorising them. It does not matter whether the node is still online or not, as it doesn't 
require any input from the node being removed.

!!! warning
    Be aware when removing nodes that cross the BFT boundary, e.g. going from 10 validators to 9, as this may impact the chains ability to progress if other nodes are offline

Removing a new validator requires that a majority of existing validators propose the new node to be removed. This is 
achieved by calling the `propose` RPC method with the value `false` and replacing the address to your required one:

    ```bash
    $ geth attach /qdata/dd/geth.ipc
    
    > istanbul.propose("0xb131288f355bc27090e542ae0be213c20350b767", false);
    null
    ```

### Example

You can find the resources required to run the examples in the 
[quorum-examples](https://github.com/jpmorganchase/quorum-examples/tree/master/examples/ibft_validator_set_changes) 
repository. 

1. The examples use `docker-compose` for the container definitions. If you are following along by copying the commands 
   described, then it is important to set the project name for Docker Compose, or to remember to change the prefix for 
   your directory. See [Docker documentation](https://docs.docker.com/compose/reference/envvars/#compose_project_name) 
   for more details.
   
   To set the project name, run the following:
   ```bash
   $ export COMPOSE_PROJECT_NAME=addnode
   ```
   
2. Bring up the network, which contains 7 nodes, of which 6 are validators.

    ```bash
    # Set the environment variable for docker-compose
    $ export COMPOSE_PROJECT_NAME=addnode
    
    # Start the 7 node network, of which 6 are validators
    $ docker-compose -f ibft-6-validators.yml up
    ```
   
3. Now we need to propose node 6 as the node to remove.
        
    !!! note
        We need a majority of existing validators to propose the new node before the changes will take effect.
   
   Lets start with node 1 and see what happens:
   
    ```bash
    # Propose node 7 from node 1
    $ docker exec -it addnode_node1_1 geth --exec 'istanbul.propose("0x8157d4437104e3b8df4451a85f7b2438ef6699ff", false);' attach /qdata/dd/geth.ipc
    null
   
    # Wait about 5 seconds, and then run:
    $ docker exec -it addnode_node1_1 geth --exec 'istanbul.getSnapshot();' attach /qdata/dd/geth.ipc
    {
      epoch: 30000,
      hash: "0xba9f9b72cad90ae8aee39f352b45f21d5ed5535b4479743e3f39b231fd717792",
      number: 140,
      policy: 0,
      tally: {
        0x8157d4437104e3b8df4451a85f7b2438ef6699ff: {
          authorize: false,
          votes: 1
        }
      },
      validators: ["0x6571d97f340c8495b661a823f2c2145ca47d63c2", "0x8157d4437104e3b8df4451a85f7b2438ef6699ff", "0xb912de287f9b047b4228436e94b5b78e3ee16171", "0xd8dba507e85f116b1f7e231ca8525fc9008a6966", "0xe36cbeb565b061217930767886474e3cde903ac5", "0xf512a992f3fb749857d758ffda1330e590fa915e"],
      votes: [{
          address: "0x8157d4437104e3b8df4451a85f7b2438ef6699ff",
          authorize: false,
          block: 136,
          validator: "0xd8dba507e85f116b1f7e231ca8525fc9008a6966"
      }]
    }
    ```
   
   Let's break this down.
   Firstly, we proposed the address `0x8157d4437104e3b8df4451a85f7b2438ef6699ff` to be removed; that is what the 
   `false` parameter is for.
   
   Secondly, we fetched the current snapshot, which gives us an insight into the current running state of the voting.
   We can see that the proposed address has 1 vote under the `tally` section, and that one vote is described under the 
   `votes` section. Here, the `authorize` section is set to `false`, which is inline with our proposal to *remove* the 
   validator.
   
4. We need to get a majority, so let's run the proposal on 3 more nodes:

    ```bash
    $ docker exec -it addnode_node2_1 geth --exec 'istanbul.propose("0x8157d4437104e3b8df4451a85f7b2438ef6699ff", false);' attach /qdata/dd/geth.ipc
    null
   
    $ docker exec -it addnode_node3_1 geth --exec 'istanbul.propose("0x8157d4437104e3b8df4451a85f7b2438ef6699ff", false);' attach /qdata/dd/geth.ipc
    null
   
    $ docker exec -it addnode_node4_1 geth --exec 'istanbul.propose("0x8157d4437104e3b8df4451a85f7b2438ef6699ff", false);' attach /qdata/dd/geth.ipc
    null
    ```
   
5. Let's check the snapshot now all the required votes are in:

    ```bash
    $ docker exec -it addnode_node1_1 geth --exec 'istanbul.getSnapshot();' attach /qdata/dd/geth.ipc
    {
      epoch: 30000,
      hash: "0x25815a32b086926875ea2c44686e4b20effabc731b2b121ebf0e0f395101eea5",
      number: 470,
      policy: 0,
      tally: {},
      validators: ["0x6571d97f340c8495b661a823f2c2145ca47d63c2", "0xb912de287f9b047b4228436e94b5b78e3ee16171", "0xd8dba507e85f116b1f7e231ca8525fc9008a6966", "0xe36cbeb565b061217930767886474e3cde903ac5", "0xf512a992f3fb749857d758ffda1330e590fa915e"],
      votes: []
    }
    ```
   
    The validator has been removed from the `validators` list, and we are left with the other 5 still present. You will 
    also see in the logs of node 6 a message like 
    `node6_1  | WARN [01-20|11:35:52.044] Block sealing failed  err=unauthorized`. This is because it is still minting 
    blocks, but realises it does not have the authority to push them to any of the other nodes on the network (you will
    also see this message for node 7, which was never authorised but still set up to mine).

## See also

- [Adding a new node to the network](/How-To-Guides/adding_nodes)