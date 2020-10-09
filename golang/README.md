New Org Binary
==============

* To create the required Directory tree and templates for a new org
* **command** ==> `go build . && ./main && ./main -role=admin`
* **output Directory Structure**
```
NewOrg
├── admin
│   └── msp
│       ├── cacerts
│       │   └── ca.cert.tpl
│       ├── keystore
│       │   └── agent.key.tpl
│       └── signcerts
│           └── agent.crt.tpl
└── peer
    └── msp
        ├── cacerts
        │   └── ca.cert.tpl
        ├── keystore
        │   └── agent.key.tpl
        └── signcerts
            └── agent.crt.tpl
```