# InnoTaxiUser

This app is a go-written service, that allows users to register, order taxi, get info about his rating and all CRUD operations
Usage

    Clone and pull this service to you local repository

  git clone git@github.com:infinitss13/InnoTaxiUser.git
  git pull url_to_your_cloned_repo

    Run docker-compose

  docker-compose up

    Run migrations to postgres:

  docker run -v /home/infiniss/GolandProjects/InnoTaxiUser/schema:/schema 
  --network host migrate/migrate -path=/schema -database "postgres://post@localhost:localhost:5432/postgres?sslmode=disable" up
  


## Documentation

All documenations is written in go-swagger documentation cn be accessed by the next url:
http://localhost:8000/swagger/index.html

Service is working on the port 8000 or 8080 if previous is't available. 
External services: 
  1. PostgreSQL - database for users : name, phone, email, hash password, rating
  2. MongoDB - NoSQL database that is used for logging
  3. Redis - NoSQL database that is used for cashing jwt-token, when user sign-out.

All non-constant variables are taken from enviroment variables, that are written in .env file
If you want to change ports of databases or Service, you can change tham in that file.
But also in that case you should change them in docker-compose.yml file


