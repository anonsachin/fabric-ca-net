vault {
    address      = "http://127.0.0.1:8200/"
    token        = "myroot"
    grace        = "1s"
    renew_token  = false
    #Default value is true
}

# Orderer MSP
template {
  source      = "/home/sachin/ca-net/network/certs/ordererorg/orderer/msp/cacerts/ca.crt.tpl"
  destination = "/home/sachin/ca-net/network/certs/ordererorg/orderer/msp/cacerts/ca.pem"
  # perms       = 0700
}

template {
  source      = "/home/sachin/ca-net/network/certs/ordererorg/orderer/msp/signcerts/agent.crt.tpl"
  destination = "/home/sachin/ca-net/network/certs/ordererorg/orderer/msp/signcerts/cert.pem"
  # perms       = 0700
}

template {
  source      = "/home/sachin/ca-net/network/certs/ordererorg/orderer/msp/keystore/agent.key.tpl"
  destination = "/home/sachin/ca-net/network/certs/ordererorg/orderer/msp/keystore/agent.key"
  # perms       = 0700
}


# Orderer TLS
template {
  source      = "/home/sachin/ca-net/network/certs/ordererorg/orderer/tls/ca.crt.tpl"
  destination = "/home/sachin/ca-net/network/certs/ordererorg/orderer/tls/ca.crt"
  # perms       = 0700
}

template {
  source      = "/home/sachin/ca-net/network/certs/ordererorg/orderer/tls/agent.crt.tpl"
  destination = "/home/sachin/ca-net/network/certs/ordererorg/orderer/tls/server.crt"
  # perms       = 0700
}

template {
  source      = "/home/sachin/ca-net/network/certs/ordererorg/orderer/tls/agent.key.tpl"
  destination = "/home/sachin/ca-net/network/certs/ordererorg/orderer/tls/server.key"
  # perms       = 0700
}

# Peer MSP
template {
  source      = "/home/sachin/ca-net/network/certs/peerorg/peer/msp/cacerts/ca.crt.tpl"
  destination = "/home/sachin/ca-net/network/certs/peerorg/peer/msp/cacerts/ca.pem"
  # perms       = 0700
}

template {
  source      = "/home/sachin/ca-net/network/certs/peerorg/peer/msp/signcerts/agent.crt.tpl"
  destination = "/home/sachin/ca-net/network/certs/peerorg/peer/msp/signcerts/cert.pem"
  # perms       = 0700
}

template {
  source      = "/home/sachin/ca-net/network/certs/peerorg/peer/msp/keystore/agent.key.tpl"
  destination = "/home/sachin/ca-net/network/certs/peerorg/peer/msp/keystore/agent.key"
  # perms       = 0700
}

# Peer TLS
template {
  source      = "/home/sachin/ca-net/network/certs/peerorg/peer/tls/ca.crt.tpl"
  destination = "/home/sachin/ca-net/network/certs/peerorg/peer/tls/ca.crt"
  # perms       = 0700
}

template {
  source      = "/home/sachin/ca-net/network/certs/peerorg/peer/tls/agent.crt.tpl"
  destination = "/home/sachin/ca-net/network/certs/peerorg/peer/tls/server.crt"
  # perms       = 0700
}

template {
  source      = "/home/sachin/ca-net/network/certs/peerorg/peer/tls/agent.key.tpl"
  destination = "/home/sachin/ca-net/network/certs/peerorg/peer/tls/server.key"
  # perms       = 0700
}

# Peer Admin MSP
template {
  source      = "/home/sachin/ca-net/network/certs/peerorg/admin/msp/cacerts/ca.crt.tpl"
  destination = "/home/sachin/ca-net/network/certs/peerorg/admin/msp/cacerts/ca.pem"
  # perms       = 0700
}

template {
  source      = "/home/sachin/ca-net/network/certs/peerorg/admin/msp/signcerts/agent.crt.tpl"
  destination = "/home/sachin/ca-net/network/certs/peerorg/admin/msp/signcerts/cert.pem"
  # perms       = 0700
}

template {
  source      = "/home/sachin/ca-net/network/certs/peerorg/admin/msp/keystore/agent.key.tpl"
  destination = "/home/sachin/ca-net/network/certs/peerorg/admin/msp/keystore/agent.key"
  # perms       = 0700
}


# Peer Admin TLS
template {
  source      = "/home/sachin/ca-net/network/certs/peerorg/admin/tls/ca.crt.tpl"
  destination = "/home/sachin/ca-net/network/certs/peerorg/admin/tls/ca.crt"
  # perms       = 0700
}

template {
  source      = "/home/sachin/ca-net/network/certs/peerorg/admin/tls/agent.crt.tpl"
  destination = "/home/sachin/ca-net/network/certs/peerorg/admin/tls/server.crt"
  # perms       = 0700
}

template {
  source      = "/home/sachin/ca-net/network/certs/peerorg/admin/tls/agent.key.tpl"
  destination = "/home/sachin/ca-net/network/certs/peerorg/admin/tls/server.key"
  # perms       = 0700
}