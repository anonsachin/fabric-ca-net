New Org Binary
==============

* Adding **NEW ORG** follows steps from [Adding New Org](https://hyperledger-fabric.readthedocs.io/en/v2.1.1/channel_update_tutorial.html)

* Doesnt include signconfigtx step as we have only one or here

* Need to set **Vault server address to** `192.168.1.102` Your computer's ip for calling vault from inside the cli.

* Edit the Collection in postman to cutomize the org found in `./collection`.

* This is run inside the GOPATH of the network cli container

* To create the required Directory tree and templates for a new org

* **command** ==> `go build . && ./main && ./main -role=admin`

* **output Directory Structure Of Default Values**

* Running **consul-template** Download Link **==>** [consul-template version 0.19.5](http://releases.hashicorp.com/consul-template/0.19.5/consul-template_0.19.5_linux_amd64.zip)
```
NewOrg
├── admin
│   ├── msp
│   │   ├── cacerts
│   │   │   ├── ca.cert.tpl
│   │   │   └── ca.pem
│   │   ├── config.yaml
│   │   ├── keystore
│   │   │   ├── agent.key
│   │   │   └── agent.key.tpl
│   │   └── signcerts
│   │       ├── agent.crt.tpl
│   │       └── cert.pem
│   └── tls
│       ├── agent.crt.tpl
│       ├── agent.key.tpl
│       ├── ca.cert.tpl
│       ├── ca.crt
│       ├── server.crt
│       └── server.key
├── msp
│   ├── cacerts
│   │   └── ca.pem
│   ├── config.yaml
│   └── tlscacerts
│       └── ca.pem
└── peer
    ├── msp
    │   ├── cacerts
    │   │   ├── ca.cert.tpl
    │   │   └── ca.pem
    │   ├── config.yaml
    │   ├── keystore
    │   │   ├── agent.key
    │   │   └── agent.key.tpl
    │   └── signcerts
    │       ├── agent.crt.tpl
    │       └── cert.pem
    └── tls
        ├── agent.crt.tpl
        ├── agent.key.tpl
        ├── ca.cert.tpl
        ├── ca.crt
        ├── server.crt
        └── server.key
```