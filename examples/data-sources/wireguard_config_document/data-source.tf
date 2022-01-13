data "wireguard_config_document" "peer1" {
  private_key = wireguard_asymmetric_key.peer1.private_key
  listen_port = 1234
  dns = [
    "1.1.1.1",
    "1.0.0.1",
    "2606:4700:4700:0:0:0:0:64",
    "2606:4700:4700:0:0:0:0:6400",
  ]

  peer {
    public_key = wireguard_asymmetric_key.peer2.public_key
    allowed_ips = [
      "0.0.0.0/0",
    ]
    persistent_keepalive = 25
  }

  peer {
    public_key = wireguard_asymmetric_key.peer3.public_key
    endpoint   = "example.com:51820"
    allowed_ips = [
      "::/0",
    ]
  }
}

output "peer1" {
  value     = data.wireguard_config_document.peer1.conf
  sensitive = true
}
