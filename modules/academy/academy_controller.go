package academy

import (
	"context"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/machinebox/graphql"
)

type AcademyController struct{}

type Query struct {
	Query      string `json:"query"`
	Start      int    `json:"start"`
	Display    int    `json:"display"`
	Adult      bool   `json:"adult"`
	Spq        bool   `json:"spq"`
	QueryRank  string `json:"queryRank"`
	DeviceType string `json:"deviceType"`
}

func (AcademyController) CrawlNaver(c echo.Context) error {
	client := graphql.NewClient("https://pcmap-api.place.naver.com/place/graphql")
	req := graphql.NewRequest(`
	query getPlacesList($input: PlacesInput) {
		businesses: places(input: $input) {
		  items {
			id
			name
			normalizedName
			category
			detailCid {
			  c0
			  c1
			  c2
			  c3
			  __typename
			}
			categoryCodeList
			dbType
			distance
			roadAddress
			address
			fullAddress
			commonAddress
			bookingUrl
			phone
			virtualPhone
			businessHours
			daysOff
			imageUrl
			imageCount
			x
			y
			poiInfo {
			  polyline {
				shapeKey {
				  id
				  name
				  version
				  __typename
				}
				boundary {
				  minX
				  minY
				  maxX
				  maxY
				  __typename
				}
				details {
				  totalDistance
				  arrivalAddress
				  departureAddress
				  __typename
				}
				__typename
			  }
			  polygon {
				shapeKey {
				  id
				  name
				  version
				  __typename
				}
				boundary {
				  minX
				  minY
				  maxX
				  maxY
				  __typename
				}
				__typename
			  }
			  __typename
			}
			subwayId     
			isPublicGas
			isDelivery
			isTableOrder
			isPreOrder
			isTakeOut
			isCvsDelivery
			hasBooking
			naverBookingCategory
			bookingDisplayName
			bookingBusinessId
			bookingVisitId
			bookingPickupId
			easyOrder {
			  easyOrderId
			  easyOrderCid
			  businessHours {
				weekday {
				  start
				  end
				  __typename
				}
				weekend {
				  start
				  end
				  __typename
				}
				__typename
			  }
			  __typename
			}
			baemin {
			  businessHours {
				deliveryTime {
				  start
				  end
				  __typename
				}
				closeDate {
				  start
				  end
				  __typename
				}
				temporaryCloseDate {
				  start
				  end
				  __typename
				}
				__typename
			  }
			  __typename
			}
			yogiyo {
			  businessHours {
				actualDeliveryTime {
				  start
				  end
				  __typename
				}
				bizHours {
				  start
				  end
				  __typename
				}
				__typename
			  }
			  __typename
			}
			isPollingStation
			hasNPay
			talktalkUrl
			visitorReviewCount
			visitorReviewScore
			blogCafeReviewCount
			bookingReviewCount
			streetPanorama {
			  id
			  pan
			  tilt
			  lat
			  lon
			  __typename
			}
			naverBookingHubId
			bookingHubUrl
			bookingHubButtonName
			newOpening
			newBusinessHours {
			  status
			  description
			  dayOff
			  dayOffDescription
			  __typename
			}
		  }
		}
	  }
	  
	`)
	// start 1 or 51
	q := Query{Query: "양천 요가학원", Start: 51, Display: 50, Adult: false, Spq: false, QueryRank: "", DeviceType: "pcmap"}
	req.Var("input", q)
	ctx := context.Background()
	var respData map[string]interface{}
	if err := client.Run(ctx, req, &respData); err != nil {
		log.Fatal(err)
	}
	return c.JSON(200, respData)
}
