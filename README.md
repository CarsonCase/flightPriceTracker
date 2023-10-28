# Flight Price Tracker
By Carson Case: carsonpcase@gmail.com

### Purpose
First and foremost. Flight price tracker is meant to be a portfolio project for backend development. It hosts a postgres server to store collected flight prices from amadeus api and serves them in a nice web interface.

# Instructions
1. Clone the repository
2. Make sure you have docker installed with postgresql[https://hub.docker.com/_/postgres]
3. Go to Amadeus[https://developers.amadeus.com] and sign up for an API key and secret
4. Use /env.example to create a .env configuration with your values
5. The makefile in cmd/server holds easy make commands for the DB. Run cd cmd/server then `make dbup` to start your database
6. Go to http://localhost:8080 to see the db adminer dashboard
7. Use `make up` and `make down` to migrate your database up and down. Migrations are stored in /sql/migrations
8. Use `dbdown` to shut off the database
9. cd ../../ to return to root directory
10. Here run `make goFlight`` to build the goFlight executable
11. ./goFlight help to see goFlight commands
12. run `make server` to build server
13. ./server will run the server and print the API key needed for `./goFlight add-route`
14. Use `./goFlight add-route` to publish routes

Note. Sort of the end here. At least until I'm willing to pay for Amadeus API which atm I am not.


