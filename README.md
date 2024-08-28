# rabbitmq-sqs-sns-go
This repository was created for learning RabbitMQ and event-driven architecture by building microservices in Go. It includes testing simple SQS and SNS services, as well as an introduction to the Advanced Message Queuing Protocol (AMQP).
- [AMQP](https://www.rabbitmq.com/tutorials/amqp-concepts)

# Installing RabbitMQ using the community Docker image:
```bash
# latest RabbitMQ 3.13
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.13-management
```
## User creation by RabbitMQCLI from the docker container
```bash
docker exec rabbitmq rabbitmqctl add_user admin password
```
adding an administrator tag to the new user
```bash
docker exec rabbitmq rabbitmqctl set_user_tags admin administrator
```
removing the guest user
```bash
docker exec rabbitmq rabbitmqctl delete_user guest
```
show users
```bash
docker exec rabbitmq rabbitmqctl list_users
```

## Creating the vhost
```bash
docker exec rabbitmq rabbitmqctl add_vhost customers
```

## Add full access to configure, write, and read on the customer vhost for user admin
```bash
docker exec rabbitmq rabbitmqctl set_permissions -p customers admin ".*" ".*" ".*"
```

## Declaration of the Topic exchange using the rabbitmqadmin.
```bash
docker exec rabbitmq rabbitmqadmin declare exchange --vhost=customers name=customer_events type=topic -u admin -p password durable=true
```

## Giving the user permission to send on this exchange, allow posting on the vhost customers on the exchange customer_events on any routing key starting with customers.
```bash
docker exec rabbitmq rabbitmqctl set_topic_permissions -p customers admin customer_events "^customers.*" "^customers.*"
```



