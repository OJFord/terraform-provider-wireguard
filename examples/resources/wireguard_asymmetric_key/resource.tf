resource "wireguard_asymmetric_key" "example" {
}

resource "wireguard_preshared_key" "example" {
}

output "wg_public_key" {
  description = "Example's public WireGuard key"
  value       = wireguard_asymmetric_key.example.public_key
}

output "wg_private_key" {
  description = "Example's private WireGuard key"
  value       = wireguard_asymmetric_key.example.private_key
  sensitive   = true
}

output "wg_preshared_key" {
  description = "Example's preshared WireGuard key"
  value       = wireguard_preshared_key.example.preshared_key
}
