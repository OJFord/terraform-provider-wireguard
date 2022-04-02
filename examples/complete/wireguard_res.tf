# terraform plan --var-file=wgvars.tfvars
terraform {
  required_providers {
    wireguard = {
      source  = "OJFord/wireguard"
      version = "~> 0.2.1"
    }
  }
}

provider "wireguard" {
  # Configuration options
}

#####################################
#
# Variables
#
#####################################
variable "server_private_key" {}
variable "server_public_key" {}
variable "server_ip_address" {}
variable "keep_alive" {}
variable "endpoint" {}
variable "listen_port" {}
variable "mtu" {}
variable "dns" {}

variable "peers_map" {
  type = list(object({
    indexCount    = number
    configFile    = string
    description   = string
    address       = list(string)
    public_key    = string
    preshared_key = string
    allowed_ips   = list(string)
  }))
}

#####################################
#
# Outputs
#
#####################################
output "server_public_key" {
  description = "Public Key"
  value       = wireguard_asymmetric_key.server.public_key
}

output "server_private_key" {
  description = "Private Key"
  value       = wireguard_asymmetric_key.server.private_key
  sensitive   = true
}

output "wireguard_server_config" {
  value     = data.wireguard_config_document.server.conf
  sensitive = true
}

#####################################
#
# Templates
#
#####################################
resource "local_file" "server" {
  filename = "outputs/wg.conf"
  content  = <<-EOT
  [Interface]
  Address = ${var.server_ip_address}
  PrivateKey = ${data.wireguard_config_document.server.private_key}
  ListenPort = ${var.listen_port}

  %{for peer in var.peers_map}
  # ${peer.description}
  [Peer]
  PublicKey = ${wireguard_asymmetric_key.peer[peer.indexCount].public_key}
  PresharedKey = ${wireguard_asymmetric_key.peer_preshared[peer.indexCount].public_key}
  AllowedIPs = ${join(",", peer.address)}
  %{endfor}
  EOT
}

data "template_file" "peers_client" {
  count    = length(var.peers_map)
  template = file("templates/client.tftpl")
  vars = {
    wgPeerDescription  = var.peers_map[count.index].description
    wgPeerAddress      = join(",", var.peers_map[count.index].address)
    wgPeerDnsAddress   = join(",", var.dns)
    wgPeerPrivateKey   = wireguard_asymmetric_key.peer[count.index].private_key
    wgPeerAllowedIps   = join(",", var.peers_map[count.index].allowed_ips)
    wgPeerEndpoint     = var.endpoint
    wgServerPort       = var.listen_port
    wgServerPublicKey  = wireguard_asymmetric_key.server.public_key
    wgPeerPresharedKey = wireguard_asymmetric_key.peer_preshared[count.index].public_key
    wgPeerKeepAlive    = var.keep_alive
    wgPeerMTU          = var.mtu
  }
}


#####################################
#
# null_resource
#
#####################################
resource "null_resource" "client_conf" {
  count = length(var.peers_map)

  triggers = {
    template = data.template_file.peers_client[count.index].rendered
  }

  provisioner "local-exec" {
    command = "echo \"${data.template_file.peers_client[count.index].rendered}\" > outputs/${var.peers_map[count.index].configFile}"
  }
}

#####################################
#
# Wireguard Peers / Preshared
#
#####################################
resource "wireguard_asymmetric_key" "peer" {
  count       = length(var.peers_map)
  description = "Peer's public / private key(s) for ${var.peers_map[count.index].description}"
}

resource "wireguard_asymmetric_key" "peer_preshared" {
  count       = length(var.peers_map)
  description = "Peer's preshared public key(s) for ${var.peers_map[count.index].description}"
}

#####################################
#
# Wireguard Server
#
#####################################
resource "wireguard_asymmetric_key" "server" {
  description = "My Wireguard Server Public/Private Keys"
}

#####################################
#
# Wireguard Provisioning
#
#####################################
data "wireguard_config_document" "server" {
  private_key = wireguard_asymmetric_key.server.private_key

  listen_port = var.listen_port
  mtu         = var.mtu
  dns         = var.dns
  description = "My Wireguard Server"

  dynamic "peer" {
    for_each = var.peers_map
    content {

      public_key           = wireguard_asymmetric_key.peer[peer.value.indexCount].public_key
      preshared_key        = wireguard_asymmetric_key.peer_preshared[peer.value.indexCount].public_key
      allowed_ips          = peer.value.allowed_ips
      endpoint             = "${var.endpoint}:${var.listen_port}"
      persistent_keepalive = var.keep_alive
      description          = "Peer # ${peer.value.indexCount} is => ${peer.value.description}"
    }
  }
}

