"""
    Code is poetry
    Created by black.dragon74 aka Nick

    Script to update the faculty details JSON
"""

# Required imports
import requests
from bs4 import BeautifulSoup
import json
import html
import datetime

# Constants
CONTACTS_URL = "https://jaipur.manipal.edu/muj/academics/faculty-list.html"
CONTACT_FIELD_ID = "facultyList"
BASE_URL = "https://jaipur.manipal.edu"

def selectid(attrval):
    obj = {'id': attrval}
    return obj


def run():
    print("Updating the faculties.json file as on %s " % datetime.datetime.now().date())

    with requests.session() as updateSession:

        requestData = updateSession.get(CONTACTS_URL).content
        updateSession.close()

        result = []

        requestSoup = BeautifulSoup(requestData, 'html.parser')

        # jsonData = getValueFromInput(requestSoup, CONTACT_FIELD_ID)
        jsonData = requestSoup.find('ul', selectid(CONTACT_FIELD_ID))["data-faculty-object"]

        if not jsonData:
            print("No JSON data found to decode from. Aborting.")
            exit(1)

        jsonSerialize = json.loads(jsonData)

        counter = 1
        for data in jsonSerialize:
            # Skip empty names, IDK why but there are such cases present on the website
            if not data["title"]:
                continue

            curr_fac = dict()
            curr_fac["id"] = counter
            curr_fac["name"] = data["title"]
            curr_fac["designation"] = data["designation"]

            # A bit of extra work for the phone number
            # Replace the country code
            phone_number = sanitize(data["phone"]).strip().replace('+91', '').replace(' ', '').replace('-', '')

            # We do not need more than 10 digit phone numbers
            if len(phone_number) > 10:
                phone_number = phone_number[2:]

            # Okay, ready to roll
            curr_fac["phone"] = phone_number
            curr_fac["department"] = sanitize(data["departmentText"])
            curr_fac["email"] = sanitize(data["email"])
            curr_fac["image"] = BASE_URL + data["thumbnailImagePath"]

            # Append to the array
            result.append(curr_fac)
            counter = counter + 1

        # Now that we are done with the data, time to write it to the file
        with open('data/faculties.json', 'w+') as output_file:
            output_file.write("%s" % json.dumps(result))
            output_file.close()

        # Success
        print("Successfully updated data with %s values" % counter)


# Fuction to sanitize and select only one email address
def sanitize(data):
    return html.unescape(data).split(';')[0]


# Make the script executable
if __name__ == '__main__':
    run()


