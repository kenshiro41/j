from bs4 import BeautifulSoup
import requests
import json
import lxml
import os
import psycopg2
import numpy as np
import re


section = int
home = str
home_score = int
away = str
away_score = int
date = str
visitor = str
ko = str
stadium = str
data = []

target_url = 'https://data.j-league.or.jp/SFTP01/'
req = requests.get(target_url)
soup = BeautifulSoup(req.content, 'lxml')
main_contents = soup.find('div', class_='tab-contents-main').stripped_strings
sec = soup.find('h2', class_='tab-title').stripped_strings
findint = int

for s in sec:
  findint = re.findall(r'\d+', s)

# insert some values in array
for index, value in enumerate(main_contents):
  if index == 0:
    continue
  data.append(value)
arr = list(np.array_split(data, 9))

# inser values to postgres
connection = psycopg2.connect(
    'postgresql://name:pass@127.0.0.1:5432/jleague')
section = int(findint[2])
for i in range(len(arr)):
  li = arr[i]
  home = li[0]
  home_score = li[1]
  away = li[4]
  away_score = li[3]
  date = li[6]
  visitor = li[8]
  ko = li[10]
  stadium = li[12]
  try:
    with connection.cursor() as cur:
      cur.execute(
          'INSERT INTO result.match_result (section, home, home_score, away, away_score, date, visitor, ko, stadium) VALUES (%s, %s, %s, %s, %s, %s, %s, %s, %s)',
          (
              section,
              home,
              home_score,
              away,
              away_score,
              date,
              visitor,
              ko,
              stadium)
      )
    connection.commit()
    print('success')
  except psycopg2.DatabaseError as err:
    print(err)
