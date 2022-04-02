server_public_key  = "NOT_A_REAL_KEY"
server_private_key = "NOT_A_REAL_KEY"
server_ip_address  = "10.100.0.1/24"
keep_alive         = 25
endpoint           = "myserver.amazonaws.com"
listen_port        = 41240
mtu                = 1480
dns                = ["10.100.0.1"]

peers_map = [
  {
    indexCount    = 0
    description   = "Samsung S20"
    configFile    = "sams20.conf"
    address       = ["10.100.0.2/32"]
    allowed_ips   = ["0.0.0.0/0"]
    preshared_key = "N/A"
    public_key    = "N/A"
  },
  {
    indexCount    = 1
    description   = "MBP 15"
    configFile    = "mbp15.conf"
    address       = ["10.100.0.3/32"]
    allowed_ips   = ["0.0.0.0/0"]
    preshared_key = "N/A"
    public_key    = "N/A"
  },
]
