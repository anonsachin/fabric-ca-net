New Org Binary
==============

* Need to set **Vault server address to** `192.168.1.102` Your computer's ip

* This is run inside the GOPATH of the network cli container

* To create the required Directory tree and templates for a new org
* **command** ==> `go build . && ./main && ./main -role=admin`
* **output Directory Structure Of Default Values**
```
NewOrg
├── admin
│   ├── msp
│   │   ├── cacerts
│   │   │   └── ca.cert.tpl
│   │   ├── keystore
│   │   │   └── agent.key.tpl
│   │   └── signcerts
│   │       └── agent.crt.tpl
│   └── tls
│       ├── agent.crt.tpl
│       ├── agent.key.tpl
│       └── ca.cert.tpl
└── peer
    ├── msp
    │   ├── cacerts
    │   │   └── ca.cert.tpl
    │   ├── keystore
    │   │   └── agent.key.tpl
    │   └── signcerts
    │       └── agent.crt.tpl
    └── tls
        ├── agent.crt.tpl
        ├── agent.key.tpl
        └── ca.cert.tpl
```