terraform {
  required_providers {
    wireguard = {
      source = "OJFord/wireguard"
    }
  }
}

resource "wireguard_asymmetric_key" "peer1" {}
resource "wireguard_asymmetric_key" "peer2" {}
resource "wireguard_asymmetric_key" "peer3" {}
