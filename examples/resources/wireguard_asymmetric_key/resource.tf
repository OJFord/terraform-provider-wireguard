resource "wireguard_asymmetric_key" "example" {
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
