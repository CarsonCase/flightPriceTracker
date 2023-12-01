from scraper.flight_analysis.src.google_flight_analysis.scrape import *
import requests
import pandas as pd
from datetime import datetime, timedelta

baseURL = "http://81.4.109.207:8000"
daysToScan = 20

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

    response = requests.post(baseURL+"/api/flights", headers=headers, json=flight_data)
    print(response)

def alreadyExists(route, date):
    response = requests.get(baseURL+"/flights")
    flights = response.json()
    if flights:
        for flight in flights:
            if flight["Date"] == date and flight["Route"] == route:
                print("Already exists: " + str(date) +"\t" + str(route))
                return True

    return False

def main():
    # DAYS IN THE FUTURE TO SCAN
    apiKey = input("Enter API Key for goFlight server:")
    resp = requests.get(baseURL+"/routes")
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
                except OSError as err:
                    print(err)
                    failedCount += 1
                except ValueError as valerr:
                    print("Value Error at: ", date,"\t last elemnt of array: ", dates[:-1])
        print("For Route: " +str(route) +"\nPublished: "+str(daysToScan - failedCount) + " of: " + str(daysToScan))
        failedCount = 0

if __name__ == "__main__":
    main()


