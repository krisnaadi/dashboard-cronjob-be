package formatter

import (
	"fmt"
	"time"

	"github.com/dustin/go-humanize"
)

func CurrencyFormat(amount float64) string {
	if amount == float64(int64(amount)) {
		return humanize.FormatInteger("#.###,", int(amount))
	}

	return humanize.FormatFloat("#.###,##", amount)
}

func TimeToUnixTime(t *time.Time) int64 {
	return t.Unix()
}

func FormattedDateToString(dateString string, layout string, format string) string {
	// Parse the input date string into a time.Time value
	// layout example time.RFC3339Nano
	date, err := time.Parse(layout, dateString)
	if err != nil {
		fmt.Println("Error parsing date formattedDateToString() function: ", err)
		return dateString
	}

	// Format the time.Time value
	// example "Monday, 2 January 2006 15:04"
	formattedDate := date.Format(format)

	return formattedDate
}

func FullDateTimeFormat(dateTime time.Time) string {
	return dateTime.Format("2006-01-02 15:04:05")
}
