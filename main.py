import csv
import json

with open('data.json') as f:
  covidData = json.load(f)

dailyCase = covidData['cases_time_series']

csvData = [['Date', 'Case']]

for data in dailyCase:
      csvData.append([data['date'], data['dailyconfirmed']])

with open("output.csv", "w") as f:
    writer = csv.writer(f)
    writer.writerows(csvData)