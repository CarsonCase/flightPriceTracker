# Flight Price Tracker
By Carson Case: carsonpcase@gmail.com

### Purpose
First and foremost. Flight price tracker is meant to be a portfolio project for backend development. It hosts a postgres server to store collected flight prices from amadeus api and serves them in a nice web interface.

# Instructions
1. Clone the repository
2. Make sure you have docker installed with postgresql[https://hub.docker.com/_/postgres]
3. Go to Amadeus[https://developers.amadeus.com] and sign up for an API key and secret
4. Use /env.example to create a .env configuration with your values
5. The makefile holds easy make commands for the DB. Run `make dbup` to start your database
6. Go to http://localhost:8080 to see the db adminer dashboard
7. Use `make up` and `make down` to migrate your database up and down. Migrations are stored in cmd/server/sql/migrations
8. Use `dbdown` to shut off the database
9. Run `make goFlight`` to build the goFlight executable
10. ./goFlight help to see goFlight commands
11. run `make server` to build server
12. ./server will run the server and print the API key needed for `./goFlight add-route`
13. Use `./goFlight add-route` to publish routes. The scraper will publish flight price data for each route logged
14. scraper/flight_analysis contains the selenium scraper to scrape Kayak for average prices of flights. Make sure that chromedriver["https://skolo.online/documents/webscrapping/#step-2-install-chromedriver] is installed
15. Run `python main.py` to begin the scraper. You may also run this on a different machine from the scraper if you adjust URL in the script.

You now have a SQL database filled with Routes and cooresponding flight price data, as well as a server serving this database

# Frontend
The repo for a simple frontend interface is here[https://github.com/CarsonCase/FlightPriceTrackerWeb]

This is what you can expect from running the entire codebase together. Note there are some hardcoded "localhosts" so you may need to change them if you're looking to properly host this system.
![Image of the frontend](<Screenshot from 2023-11-13 14-17-36.png>)