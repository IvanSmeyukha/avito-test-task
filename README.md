Starting an application:
> docker-compose up

Requests:

Get user balance (GET)
>http://localhost:8080/api/users/userId

Balance replenishment (PUT)
>http://localhost:8080/api/users/userId

>Body: {
    "amount": 100
}

Funds reservation (PUT)
>http://localhost:8080/api/reserve/userId

>Body: {
    "orderId": 1,
    "serviceId": 1,
    "value": 100
}

Write off revenue (DELETE)
>http://localhost:8080/api/reserve/userId

>Body: {
    "orderId": 1,
    "serviceId": 1,
    "value": 100
}
