output "wg_public_key" {
  description = "Example's public WireGuard key"
  value       = wireguard_asymmetric_key.example.public_key
}
