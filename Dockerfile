# Use the official RabbitMQ image as the base image
FROM rabbitmq:3.13-management

# Copy your configuration files into the container
COPY configs/rabbitmq/rabbitmq.conf /etc/rabbitmq/rabbitmq.conf
COPY configs/rabbitmq/definitions.json /etc/rabbitmq/definitions.json

# Expose RabbitMQ ports
EXPOSE 5672 15672
