# Flight Price Tracker
By Carson Case: carsonpcase@gmail.com

### Purpose
First and foremost. Flight price tracker is meant to be a portfolio project for backend development. It hosts a postgres server to store collected flight prices from amadeus api and serves them in a nice web interface.

# Instructions
1. Clone the repository
2. Make sure you have docker installed with postgresql[https://hub.docker.com/_/postgres]
3. Go to Amadeus[https://developers.amadeus.com] and sign up for an API key and secret
4. Use /env.example to create a .env configuration with your values
5. Run `make dbup` to start your database
6. Go to localhost:8080 to see the db adminer dashboard
7. Use `make up` and `make down` to migrate your database up and down. Migrations are stored in /sql/migrations
8. Use `dbdown` to shut off the database

