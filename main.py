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

def alreadyExists(route, date):
    response = requests.get("http://localhost:8000/flights")
    for flight in response.json():
        if flight["Date"] == date and flight["Route"] == route:
            print("Already exists: " + str(date) +"\t" + str(route))
            return True

    return False

def main():
    # DAYS IN THE FUTURE TO SCAN
    daysToScan = 2
    apiKey = input("Enter API Key for goFlight server:")
    resp = requests.get(urlRoutes)
    routesData = resp.json()

    dates = pd.date_range((datetime.today() + timedelta(1)), (datetime.today() + timedelta(daysToScan))).tolist()
    failedCount = 0

    for route in routesData:
        for date in dates:
            date = date.strftime("%Y-%m-%d")
            print(route["ID"], date)
            if(alreadyExists(route["ID"], date)):
                failedCount += 1
                continue
            else:
                try:
                    result = Scrape(route["Departure"], route["Arrival"], date)
                    ScrapeObjects(result)
                    postFlight(apiKey, route, date, result)
                except:
                    print()
                    failedCount += 1
        print("For Route: " +str(route) +"\nPublished: "+str(daysToScan - failedCount) + " of: " + str(daysToScan))
        failedCount = 0

if __name__ == "__main__":
    main()


