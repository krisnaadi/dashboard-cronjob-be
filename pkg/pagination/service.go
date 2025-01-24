package pagination

import (
	"net/url"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/tech-djoin/wallet-djoin-service/internal/constant"
)

func GetStartRowAndLimit(page int) (int, int) {
	if page > 0 {
		return ((constant.LIMIT * page) - constant.LIMIT), constant.LIMIT
	}
	return 0, constant.LIMIT
}

func GetPrevPage(c echo.Context, page int) string {
	if page > 1 {
		originalParams := c.QueryParams()

		// Create a new set of query parameters
		newParams := url.Values{}

		// Copy the original query parameters to the new parameters
		for key, values := range originalParams {
			for _, value := range values {
				newParams.Add(key, value)
			}
		}

		newParams.Set("page", strconv.Itoa((page - 1)))

		return newParams.Encode()
	}

	return ""
}

func GetNextPage(c echo.Context, page int, endrow int, maxData int64) string {
	if endrow < int(maxData) {
		originalParams := c.QueryParams()

		// Create a new set of query parameters
		newParams := url.Values{}

		// Copy the original query parameters to the new parameters
		for key, values := range originalParams {
			for _, value := range values {
				newParams.Add(key, value)
			}
		}
		newParams.Set("page", strconv.Itoa((page + 1)))

		newParams.Encode()

		return newParams.Encode()
	}

	return ""
}
