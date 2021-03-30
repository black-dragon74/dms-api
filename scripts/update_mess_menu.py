"""
    Code is poetry
    
    Created by Nick aka black.dragon74
"""

from requests_html import HTMLSession
from bs4 import BeautifulSoup

import json
import datetime

_messURL = "https://mujhub.com"

_menuResp = {
    "last_updated_at": "",
    "last_updated_meal": None,
    "breakfast": [],
    "lunch": [],
    "high_tea": [],
    "dinner": [],
}

_meals = ["breakfast", "lunch", "high_tea", "dinner"]
_xpaths = [
    '//*[@id="root"]/div/div[2]/div[2]/div[2]/div[1]/p',
    '//*[@id="root"]/div/div[2]/div[2]/div[2]/div[2]/p',
    '//*[@id="root"]/div/div[2]/div[2]/div[2]/div[3]/p',
    '//*[@id="root"]/div/div[2]/div[2]/div[2]/div[4]/p'
]
_na_msg = "Not yet updated."


def update_mess_menu():
    try:
        print("Updating mess menu -- {} : {}".format(datetime.datetime.now().date(), datetime.datetime.now().time()))
        session = HTMLSession()
        res = session.get(_messURL)
        print("Request sent. Await response.")

        if not res:
            print("Unable to fetch response.")
            exit(1)

        res.html.render(sleep=5)
        upd_date = res.html.search('MESS MENU - {}<')[0]

        if upd_date is not None:
            _menuResp["last_updated_at"] = upd_date.replace('.', '-')
            print("Found mess menu for date %s" % _menuResp["last_updated_at"])
        else:
            print("Unable to find the last updated date, continue anyways.")

        i = 0
        for meal in _meals:
            currMeal = res.html.xpath(_xpaths[i])[0].text
            _menuResp[meal] = [item.strip() for item in currMeal.split(',')]
            i += 1

        # Figure out the last updated meal
        for meal in _meals:
            if _menuResp[meal] and _menuResp[meal] != [_na_msg]:
                _menuResp["last_updated_meal"] = meal

        print("Found last updated meal as: %s" % _menuResp["last_updated_meal"])

        for key, value in _menuResp.items():
            if not value and key in _meals:
                _menuResp[key] = ["Not updated yet"]
                print("Set status as 'NA' for meal: %s" % key)

        with open('data/mess_menu.json', 'w+') as messFile:
            messFile.write("%s" % json.dumps(_menuResp))
            messFile.close()

            print("Successfully memoized the values.")

        print(" ")
    except Exception as e:
        print("Failed to update the values.\nException caught :: %s" % e)
        print(" ")
        exit(1)


def run():
    update_mess_menu()


if __name__ == '__main__':
    run()
