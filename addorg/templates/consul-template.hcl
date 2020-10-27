vault {
    address      = "VAULTHOST"
    token        = "myroot"
    grace        = "1s"
    renew_token  = false
    #Default value is true
}

# MSP
template {
  source      = "BASEPATH/ORG/ROLE/msp/cacerts/ca.cert.tpl"
  destination = "BASEPATH/ORG/ROLE/msp/cacerts/ca.pem"
  # perms       = 0700
}

template {
  source      = "BASEPATH/ORG/ROLE/msp/signcerts/agent.crt.tpl"
  destination = "BASEPATH/ORG/ROLE/msp/signcerts/cert.pem"
  # perms       = 0700
}

template {
  source      = "BASEPATH/ORG/ROLE/msp/keystore/agent.key.tpl"
  destination = "BASEPATH/ORG/ROLE/msp/keystore/agent.key"
  # perms       = 0700
}


# TLS
template {
  source      = "BASEPATH/ORG/ROLE/tls/ca.cert.tpl"
  destination = "BASEPATH/ORG/ROLE/tls/ca.crt"
  # perms       = 0700
}

template {
  source      = "BASEPATH/ORG/ROLE/tls/agent.crt.tpl"
  destination = "BASEPATH/ORG/ROLE/tls/server.crt"
  # perms       = 0700
}

template {
  source      = "BASEPATH/ORG/ROLE/tls/agent.key.tpl"
  destination = "BASEPATH/ORG/ROLE/tls/server.key"
  # perms       = 0700
}