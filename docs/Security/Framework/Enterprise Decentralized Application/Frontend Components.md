## Frontend
An **Enterprise Decentralized Application (eDApp)** is nothing more than typical Decentralized Application (dApp). Unlike dApps eDApps runs on an Enterprise Ethereum Client such as Quorum. Similarly to 
dApps eDApps are architecture is usually architect using either two or three tier representation of Client/Server Architecture. 

In **three tire architecture** the client contains user facing frontend built with traditional web, mobile, or desktop based application, and communicates with Backend server using standard transport protocols TCP/UDP. In dApp and eDApp Server side components 
can implement some of the business logic of the main application purpose in a traditional approach (Microservice Architecture, Service Oriented Architecture ..etc), however part or all of the server side components will
communicate with a ledger by connecting to the ledgers Remote Produce Call (RPC) to use Use/Manage Smart Contracts that supports overall application business use cases.

From another hand in **two tire architecture** the frontend component might communicate directly with the ledger through the RPC interface to interact with the Smart Contracts.
Both two or three tire architecture have advantage and disadvantages. For instance in a two tire architecture the network topology might require the direct exposure of the ledger to its user something which is considered to be an insecure practice. However this model provide true decentralization availability.  

As any traditional application, eDApps Client/Server component has to follow application security best practices. Its recommended to use a Secure Software Development Lifecycle (SSDLC) to implement any application. In general Secure Development Lifecycle include the following phases :

- Risk Assessment Phase.
- Threat Modeling, and Secure Design Review Phase.
- Static Source Code Analysis Phase.
- Security Testing, and Manual Secure Source Review Phase.
- Security Assessment of Configuration of Deployment Pipeline Phase.   
