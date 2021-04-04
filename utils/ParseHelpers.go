package utils

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/black-dragon74/dms-api/types"
	"regexp"
	"strings"
)

// GetValueFromInput returns the `value` attribute of an HTML `input` element
func GetValueFromInput(document *goquery.Document, inputID string) string {
	val, exists := document.Find(inputID).Attr("value")
	if !exists {
		return ""
	}

	return val
}

// ParseHTMLTable parses the HTML table and returns types.HTMLTableModel
func ParseHTMLTable(tbl *goquery.Document, elemID string) (types.HTMLTableModel, error) {
	// First, get the table
	table := tbl.Find(elemID)
	if table.Length() == 0 {
		return types.HTMLTableModel{}, ErrorTableNotFound
	}

	// The DMS tables have no `thead` tag instead the fuckers have put `thead` in `tbody`
	// Means, they use `th` elements in `tbody > tr`
	tableBody := table.Find("tr")
	if tableBody.Length() == 0 {
		return types.HTMLTableModel{}, ErrorTableNoBody
	}

	// Get the headers if they exist
	tableHeader := tableBody.Find("th")
	if tableHeader.Length() == 0 {
		return types.HTMLTableModel{}, ErrorTableNoHeader
	}

	// Read the headers
	headers := make([]string, 0)
	tableHeader.Each(func(i int, selection *goquery.Selection) {
		currStr := toSnakeCase(stripSpecialChars(selection.Text()))
		if currStr != "" {
			headers = append(headers, currStr)
		}
	})

	// Set the header
	retVal := types.HTMLTableModel{
		Header: headers,
	}

	// Now get the `td` values
	dataContainer := table.Find("tr:not(:first-child)")
	if dataContainer.Length() == 0 {
		return types.HTMLTableModel{}, ErrorTableNoData
	}

	// Read the values
	dataArr := make([][]string, 0)
	dataContainer.Each(func(i int, row *goquery.Selection) {
		tArr := make([]string, 0)
		row.Find("td").Each(func(i int, v *goquery.Selection) {
			tArr = append(tArr, strings.TrimSpace(v.Text()))
		})
		dataArr = append(dataArr, tArr)
	})
	retVal.Body = dataArr

	return retVal, nil
}

func SwitchSemester(sem string) (string, error) {
	switch sem {
	case "1":
		return "I", nil
	case "2":
		return "II", nil
	case "3":
		return "III", nil
	case "4":
		return "IV", nil
	case "5":
		return "V", nil
	case "6":
		return "VI", nil
	case "7":
		return "VII", nil
	case "8":
		return "VIII", nil
	case "9":
		return "IX", nil
	case "10":
		return "X", nil
	case "11":
		return "XI", nil
	case "12":
		return "XII", nil
	default:
		return "", ErrorInvalidSemester
	}
}

// StripSpecialChars removes all special characters from `s`
func stripSpecialChars(s string) string {
	stripper := regexp.MustCompile("[^0-9a-zA-Z]+")
	return stripper.ReplaceAllString(s, "")
}

// StripParen returns a string by removing the parentheses
func StripParen(str string) string {
	return regexp.MustCompile(`\(.*?\)`).ReplaceAllString(str, "")
}

// ToSnakeCase converts `s` to snake_case
func toSnakeCase(s string) string {
	matchFirstCap := regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap := regexp.MustCompile("([a-z0-9])([A-Z])")

	snake := matchFirstCap.ReplaceAllString(s, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")

	return strings.ToLower(snake)
}
