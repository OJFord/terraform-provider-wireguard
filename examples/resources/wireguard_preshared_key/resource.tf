resource "wireguard_preshared_key" "example" {
}

output "wg_preshared_keyq" {
  description = "Example's preshared WireGuard key"
  value       = wireguard_preshared_key.example.key
  sensitive = true
}
