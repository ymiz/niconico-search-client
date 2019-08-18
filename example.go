package main

import (
	"log"
	"net/http"
	"nicosearch/live"
	"nicosearch/live/fields"
	"nicosearch/live/filters"
	"nicosearch/live/sort"
	"nicosearch/live/targets"
	"time"
)

func main() {
	http.DefaultClient.Timeout = 10 * time.Second

	resp, err := live.Find(live.SearchParameter{
		Keywords: "ゲーム",
		Targets: targets.Targets{
			Field: []targets.Field{
				targets.Tags,
			},
		},
		Fields: fields.Fields{
			Field: []fields.Field{
				fields.ContentId,
				fields.Title,
				fields.ViewCounter,
			},
		},
		Filters: []filters.Filter{
			{
				Field:     filters.ProviderType,
				Condition: "0",
				Value:     "community",
			},
			{
				Field:     filters.LiveStatus,
				Condition: "0",
				Value:     "onair",
			},
		},
		Sort: sort.Sort{
			Field:   sort.ViewCounter,
			OrderBy: sort.Desc,
		},
		Context:   "niconico-search-client",
		UserAgent: "niconico-search-client",
	})
	if err != nil {
		log.Println("live search error.", err)
		return
	}
	log.Println(resp)
}
