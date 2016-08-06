# JivoxCensusGo
The API to fetch the census data for population and literacy growth<br />
[![Build Status](https://travis-ci.org/shubhodeep9/JivoxCensusGo.svg?branch=master)](https://travis-ci.org/shubhodeep9/JivoxCensusGo)

## Send request at
http://myffcs.in:6090/census/scrape

## Output format
```json
{
  "Andaman and Nicobar Islands": {
    "population": "380,581",
    "increase": "6.86 %",
    "area": "8,249",
    "density": "46",
    "sex": "876",
    "literacy": "86.63",
    "details": {
      "ActualPopulation": {
        "2001": "356,152",
        "2011": "380,581"
      },
      "ApproximatePopulation": {
        "2001": "3.56 Lakh",
        "2011": "3.81 Lakhs"
      },
      "Area(Km2)": {
        "2001": "8,249",
        "2011": "8,249"
      },
      "Areami2": {
        "2001": "3,185",
        "2011": "3,185"
      },
      "ChildSexRatio": {
        "2001": "957",
        "2011": "968"
      },
      "Density/km2": {
        "2001": "43",
        "2011": "46"
      },
      "Density/mi2": {
        "2001": "112",
        "2011": "119"
      },
      "Female": {
        "2001": "163,180",
        "2011": "177,710"
      },
      "FemaleLiteracy": {
        "2001": "75.24 %",
        "2011": "82.43 %"
      },
      "FemaleLiterate": {
        "2001": "106,304",
        "2011": "129,904"
      },
      "FemalePopulation(0-6Age)": {
        "2001": "21,896",
        "2011": "20,108"
      },
      "Literacy": {
        "2001": "81.30 %",
        "2011": "86.63 %"
      },
      "Male": {
        "2001": "192,972",
        "2011": "202,871"
      },
      "MaleLiteracy": {
        "2001": "86.33 %",
        "2011": "90.27 %"
      },
      "MaleLiterate": {
        "2001": "146,831",
        "2011": "164,377"
      },
      "MalePopulation(0-6Age)": {
        "2001": "22,885",
        "2011": "20,770"
      },
      "PercantageoftotalPopulation": {
        "2001": "0.03%",
        "2011": "0.03%"
      },
      "PopulationGrowth": {
        "2001": "26.94%",
        "2011": "6.86%"
      },
      "SexRatio": {
        "2001": "846",
        "2011": "876"
      },
      "TotalChildPopulation(0-6Age)": {
        "2001": "44,781",
        "2011": "40,878"
      },
      "TotalLiterate": {
        "2001": "253,135",
        "2011": "294,281"
      }
    }
  }, ..
}
```
