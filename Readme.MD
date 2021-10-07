# Start postgres


db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=postgres sslmode=disable password=example")


docker-compose -f stack.yml up 
