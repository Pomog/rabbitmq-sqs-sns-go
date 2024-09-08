# Generate the CA (Certificate Authority) private key
A private key, Base64 encoded file, is part of the public-private key pair.
The private key is kept secret and is used to sign certificates or messages that others can verify
with the corresponding public key.
- **genpkey** - is a more modern and flexible command,  compared to genpkey, for generating private keys.
It can handle a variety of algorithms (e.g., RSA, DSA, ECC) and options.
- **genrsa** - is specific to RSA keys and is mainly used in legacy applications.
```bash
openssl genpkey -algorithm RSA -out ca_key.pem -pkeyopt rsa_keygen_bits:2048
```
- to open RSA Base64 encoded file
```bash
sudo openssl rsa -in ca_key.pem -text
```

# Generate the CA certificate
```bash
openssl req -x509 -new -nodes -key /path/to/certs/ca_key.pem -sha256 -days 365 -out /path/to/certs/ca_certificate.pem \
-subj "/C=US/ST=State/L=City/O=Organization/OU=Unit/CN=rabbitmq"
```

# Generate a private key for the server
```bash
openssl genpkey -algorithm RSA -out /path/to/certs/server_blackbox_key.pem -pkeyopt rsa_keygen_bits:2048
```

# Create a certificate signing request (CSR) for the server
```bash
openssl req -new -key server_blackbox_key.pem -out server_blackbox_csr.pem -subj "/C=US/ST=State/L=City/O=Organization/OU=Unit/CN=rabbitmq"
```

# Sign the server certificate with the CA
```bash
openssl x509 -req -in server_blackbox_csr.pem -CA ca_certificate.pem -CAkey ca_key.pem -CAcreateserial -out server_blackbox_certificate.pem -days 365 -sha256
```
