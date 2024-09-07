# Use the official RabbitMQ image as the base image
FROM rabbitmq:3.13-management

# Create a directory for certificates in the container
RUN mkdir -p /certs

# Copy the certificates to the container
COPY certs/ca_certificate.pem /certs/ca_certificate.pem
COPY certs/server_blackbox_certificate.pem /certs/server_blackbox_certificate.pem
COPY certs/server_blackbox_key.pem /certs/server_blackbox_key.pem

# Copy your configuration files into the container
COPY configs/rabbitmq/rabbitmq.conf /etc/rabbitmq/rabbitmq.conf
COPY configs/rabbitmq/definitions.json /etc/rabbitmq/definitions.json

# Expose RabbitMQ ports
EXPOSE 5672 15672
