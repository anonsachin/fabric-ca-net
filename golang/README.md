New Org Binary
==============

* Need to set **Vault server address to** `192.168.1.102` Your computer's ip

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
│   │   │   └── ca.cert.tpl
│   │   ├── keystore
│   │   │   └── agent.key.tpl
│   │   └── signcerts
│   │       └── agent.crt.tpl
│   └── tls
│       ├── agent.crt.tpl
│       ├── agent.key.tpl
│       └── ca.cert.tpl
├── msp
│   ├── cacerts
│   │   └── ca.pem
│   └── tlscacerts
│       └── ca.pem
└── peer
    ├── msp
    │   ├── cacerts
    │   │   ├── ca.cert.tpl
    │   │   └── ca.pem
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