package twitterhelper

import (
	"encoding/json"
	"fmt"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

func TopTrending(client *twitter.Client) []string {

	fmt.Printf("======================================================")
	fmt.Println("Fetching top trending hastags lists from INDIA......")

	TrendListName := make([]string, 5, 5)
	// trending lists of india
	trendlist, _, err := client.Trends.Place(int64(2282863), &twitter.TrendsPlaceParams{WOEID: 2282863})
	trendJSON, _ := json.MarshalIndent(trendlist, "", " ")
	//fmt.Println(string(trendJSON))

	// Unmarshaling JSON DATA
	err = json.Unmarshal(trendJSON, &trendlist)
	if err != nil {
		log.Error("Error in Unmarshaling")
	}

	//fmt.Println(trendlist[0].Trends[0].Name)

	for Ids := range trendlist[0].Trends {

		TrendListName = append(TrendListName, trendlist[0].Trends[Ids].Name)
	}

	fmt.Printf("======================================================")
	fmt.Println("Completed retrieving", len(TrendListName), "top trending hastags lists from INDIA")

	return TrendListName

}
