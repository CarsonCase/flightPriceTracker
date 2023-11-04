from scraper.flight_analysis.src.google_flight_analysis.scrape import *
import requests
import pandas as pd
from datetime import datetime, timedelta

urlRoutes = "http://localhost:8000/routes"
urlFlights = "http://localhost:8000/flights"

# todo: get api key from argument
def postFlight(apiKey, route, date, result):
    headers = {
        "Authorization": apiKey,
        "Content-Type": "application/json"
    }
    
    flight_data = {
        "Route": route["ID"],
        "Date": date,
        "Price": result.data["Price ($)"].mean()
    }

    response = requests.post("http://localhost:8000/api/flights", headers=headers, json=flight_data)
    print(response)

def main():
    apiKey = input("Enter API Key for goFlight server:")
    resp = requests.get(urlRoutes)
    routesData = resp.json()

    dates = pd.date_range((datetime.today() + timedelta(1)).strftime("%Y-%m-%d"), (datetime.today() + timedelta(3))).tolist()

    for route in routesData:
        for date in dates:
            date = date.strftime("%Y-%m-%d")
            result = Scrape(route["Departure"], route["Arrival"], date)
            ScrapeObjects(result)
            postFlight(apiKey, route, date, result)

if __name__ == "__main__":
    main()


