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
- to open RSA Base64 encoded file. ```openssl rsa``` command is used to read and manipulate private keys.
    ```-noout``` suppresses the output of the certificate itself in its encoded form (Base64),
    so only the human-readable details are displayed.
```bash
sudo openssl rsa -in ca_key.pem -text -noout
```

# Generate the self-signed certificate for the Certificate Authority (CA), contains the public key and certificate information. With the entity's identity information provided by ```-subj``` option.
```bash
openssl req -x509 -new -nodes -key ca_key.pem -sha256 -days 365 -out ca_certificate.pem \
-subj "/C=US/ST=State/L=City/O=Organization/OU=Unit/CN=rabbitmq"
```
- to open the Base64 encoded ```X.509``` public key certificate information file. 
```bash
openssl x509 -in ca_certificate.pem -text -noout
```

# Generate a private key for the server
```bash
openssl genpkey -algorithm RSA -out server_blackbox_key.pem -pkeyopt rsa_keygen_bits:2048
```

# Create a certificate signing request (CSR), as new file ```server_blackbox_csr.pem```, for the server based on previously created private key ```server_blackbox_key.pem```. This CSR does not contain the private key, but contain generated public key. The server’s private key ```server_blackbox_key.pem``` is used to digitally sign the CSR.
```bash
openssl req -new -key server_blackbox_key.pem -out server_blackbox_csr.pem -subj "/C=US/ST=State/L=City/O=Organization/OU=Unit/CN=rabbitmq"
```
- to open the Base64 encoded certificate signing request
```bash
sudo openssl req -in server_blackbox_csr.pem -text -noout
```

# Sign the server certificate with the CA. Signin includes hashing all the relevant data in the CSR using CA's private key, this encrypted hash becomes the digital signature of the certificate. ```-CAcreateserial``` creates a unique serial number for the certificate.
```bash
openssl x509 -req -in server_blackbox_csr.pem -CA ca_certificate.pem -CAkey ca_key.pem -CAcreateserial -out server_blackbox_certificate.pem -days 365 -sha256
```
- to open the Base64 encoded signed certificate.
```bash
sudo openssl x509 -in server_blackbox_certificate.pem -text -noout
```

# The Client uses the CA's public key to decrypt the signature on the signed certificate, and gets the original hash (the one the CA encrypted). Then hashes the certificate's content (excluding the signature) to create a new hash value.
If the new hash value matches the decrypted hash, this proves that:
* The certificate was signed by the CA (because only the CA’s private key could create the signature).
* The certificate’s content (public key, identity info) has not been changed.

# This scheme not included SAN (Subject Alternative Name) in the certificate.