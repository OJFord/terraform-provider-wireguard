resource "wireguard_preshared_key" "example" {
}

output "wg_preshared_key" {
  description = "Example's preshared WireGuard key"
  value       = wireguard_preshared_key.example.key
  sensitive = true
}
