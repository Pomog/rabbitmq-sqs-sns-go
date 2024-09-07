# Generate the CA (Certificate Authority) private key
openssl genpkey -algorithm RSA -out /path/to/certs/ca_key.pem -pkeyopt rsa_keygen_bits:2048

# Generate the CA certificate
openssl req -x509 -new -nodes -key /path/to/certs/ca_key.pem -sha256 -days 365 -out /path/to/certs/ca_certificate.pem \
-subj "/C=US/ST=State/L=City/O=Organization/OU=Unit/CN=your-ca-name"

# Generate a private key for the server
openssl genpkey -algorithm RSA -out /path/to/certs/server_blackbox_key.pem -pkeyopt rsa_keygen_bits:2048

# Create a certificate signing request (CSR) for the server
openssl req -new -key /path/to/certs/server_blackbox_key.pem -out /path/to/certs/server_blackbox_csr.pem \
-subj "/C=US/ST=State/L=City/O=Organization/OU=Unit/CN=your-server-name"

# Sign the server certificate with the CA
openssl x509 -req -in /path/to/certs/server_blackbox_csr.pem -CA /path/to/certs/ca_certificate.pem -CAkey /path/to/certs/ca_key.pem \
-CAcreateserial -out /path/to/certs/server_blackbox_certificate.pem -days 365 -sha256
