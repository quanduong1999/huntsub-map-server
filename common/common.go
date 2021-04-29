package common

import (
	"fmt"
	"huntsub/huntsub-map-server/o/report/user"
	"huntsub/huntsub-map-server/x/expo"
	"math"
)

// Distance function returns the distance (in meters) between two points of
//     a given longitude and latitude relatively accurately (using a spherical
//     approximation of the Earth) through the Haversin Distance Formula for
//     great arc distance on a sphere with accuracy for small distances
//
// point coordinates are supplied in degrees and converted into rad. in the func
//
// distance returned is METERS!!!!!!
// http://en.wikipedia.org/wiki/Haversine_formula
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

/*************SEND NOTIFICATION*************/
func SendNotification(tokenDevice, message string) {
	// To check the token is valid
	pushToken, err := expo.NewExponentPushToken(tokenDevice)
	if err != nil {
		panic(err)
	}

	// Create a new Expo SDK client
	client := expo.NewPushClient(nil)
	array := []expo.ExponentPushToken{}
	array = append(array, pushToken)
	// Publish message
	response, err := client.Publish(
		&expo.PushMessage{
			To:       array,
			Body:     message,
			Data:     map[string]string{"withSome": "data"},
			Sound:    "default",
			Title:    "",
			Priority: expo.DefaultPriority,
		},
	)
	// Check errors
	if err != nil {
		panic(err)
		return
	}
	// Validate responses
	if response.ValidateResponse() != nil {
		fmt.Println(response.PushMessage.To, "failed")
	}
}

func CheckExist(key string, data interface{}) int {
	obj, ok := data.(*user.UserReport)
	if !ok {
		return -1
	}

	var index = -1
	for i, c := range obj.AnalysisFields {
		if c.Category == key {
			index = i
		}
	}
	if index > -1 {
		return index
	}

	return -1
}

func GetCategorys(esl float64, data interface{}) []string {
	var strs = []string{}
	obj, ok := data.(*user.UserReport)
	if !ok {
		return nil
	}
	for _, c := range obj.AnalysisFields {
		if c.Percent > esl {
			strs = append(strs, c.Category)
		}
	}
	return strs
}
