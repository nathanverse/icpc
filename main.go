package main

/*
*
Problem: Campaign Matching in Programmatic Advertising
We are running the following programmatic advertising campaigns:
Campaign 1:
- Location: United States of America, Canada
- Device OS: Windows, iOS
- Banner format(widthxheight): 320x50, 300x600
- Bid price: 1 USD

Campaign 2:
- Location: United States of America
- Banner formats (widthxheight): 320x50, 728x90
- Contextual targeting: Sports
- Bid price: 2 USD

The system loads all campaign targeting configurations and receives ad requests to evaluate and bid.
A campaign is considered a suitable candidate only if all targeting criteria match the incoming ad request.


//
	terms [Badminton, Soccer, Shirt]
//

// example.com/abc/xyz.html -> Sports

Example ad request format in JSON:
	{
	    "id: "a9486130-8e03-473d-9a8e-b92fd496069b",
	    "device": {
	        "os": "Windows 11",
	        "geo: {
	            "country": "US"
	        }
	    },
	    "site": {
	        "url": "example.com/abc/xyz.html"
	    },
	    "banner": {
	        "floorPrice": 2,
	        "floorCurrency": "USD",
	        "format": [
	            {
	                "w": 320,
	                "h": 30
	            },
	            {
	                "w": 300,
	                "h": 50
	            }
	        ]
	    }
	}

Task:
- Design and implement a function that matches an incoming ad request with suitable campaign(s).
*/

//func maxSumAfterPartitioning(arr []int, k int) int {
//	dp := make([]int, len(arr)+1)
//	curMax := -1
//	for j := len(arr) - 1; j >= 0; j-- {
//		if len(arr)-j <= k {
//			curMax = max(arr[j], curMax)
//			dp[j] = curMax * (len(arr) - j)
//			continue
//		}
//
//		curMax = arr[j]
//		dpJ := arr[j] + dp[j+1]
//		for i := 1; i < k; i++ {
//			curMax = max(curMax, arr[j+i])
//			maxDpI := curMax * (i + 1)
//			dpJ = max(maxDpI+dp[j+i+1], dpJ)
//		}
//
//		dp[j] = dpJ
//	}
//
//	return dp[0]
//}

Campaign 1:
- Location: United States of America, Canada
- Device OS: Windows, iOS
- Banner format(widthxheight): 320x50, 300x600
- Bid price: 1 USD

Campaign 2:
- Location: United States of America
- Banner formats (widthxheight): 320x50, 728x90
- Contextual targeting: Sports
- Bid price: 2 USD


type Campaign struct {
	Id int
	Locations []string
	DeviceOSs []string
	BannerFormats []string
	Bid float64
	Contexts []string
}

type Format struct {
	W int
	H int
}

type Configuration struct {
	Id string
	DeviceOS string
	Country string
	Site string
	FloorPrice float64
	Formats []Format
}

func SaveStringMap(storedMap map[string][]int, elements[]string, campaignId int){
	for _, l := range elements {
		if _, ok := storedMap[l]; !ok {
			storedMap[l] = []int{}
		}

		storedMap[l] = append(storedMap[l], campaignId)
	}
}

func matchCampaigns(config Configuration, campaigns []Campaign) []int {
	locations := map[string][]int{}
	deviceOss := map[string][]int{}
	banners := map[string][]int{}
	for _, c := range campaigns {
		SaveStringMap(locations, c.Locations, c.Id)
		SaveStringMap(deviceOss, c.Locations, c.Id)
		SaveStringMap(banners, c.BannerFormats, c.Id)
	}

	locationMatchCampaigns := locations[config.Country] // 1, 2, 3, 4
	deviceOsMatchCampaigns := deviceOss[config.Country] // 1, 2, 3, 4
	bannersMatchCampaigns := banners[config.Country] // 3, 2, 1aasaa asds assd

	//mapp := map[int]struct{}{}
	//mergeMapp := map[int]struct{}{}


}

//func main() {
//	fmt.Println(maxSumAfterPartitioning([]int{1, 4, 1, 5, 7, 3, 6, 1, 9, 9, 3}, 4))
//}
