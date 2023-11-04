from scraper.flight_analysis.src.google_flight_analysis.scrape import *
import requests
import pandas as pd
from datetime import datetime, timedelta

urlRoutes = "http://localhost:8000/routes"
urlFlights = "http://localhost:8000/flights"


def postFlight(route, date, result):
    headers = {
        "Authorization": "e0a251f8db16f4fc235679339139d483b07660814730598a119c387e9bb72a48",
        "Content-Type": "application/json"
    }
    
    flight_data = {
        "Route": route["ID"],
        "Date": date,
        "Price": result.data["Price ($)"].mean()
    }

    response = requests.post("http://localhost:8000/api/flights", headers=headers, json=flight_data)
    print(response)


resp = requests.get(urlRoutes)
routesData = resp.json()

dates = pd.date_range((datetime.today() + timedelta(1)).strftime("%Y-%m-%d"), (datetime.today() + timedelta(3))).tolist()

for route in routesData:
    for date in dates:
        date = date.strftime("%Y-%m-%d")
        result = Scrape(route["Departure"], route["Arrival"], date)
        ScrapeObjects(result)
        postFlight(route, date, result)


