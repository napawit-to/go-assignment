# Start postgres


db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=example")


docker-compose -f stack.yml up 


# APIs

curl -i http://127.0.0.1:8081/transaction/withdraw

curl -i http://127.0.0.1:8081/transaction/deposit

curl -i http://127.0.0.1:8081/customer

curl -i http://127.0.0.1:8081/account/{id}
