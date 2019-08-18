package live

import (
	"github.com/ymiz/niconico-search-client/common"
	"github.com/ymiz/niconico-search-client/live/fields"
	"github.com/ymiz/niconico-search-client/live/filters"
	"github.com/ymiz/niconico-search-client/live/sort"
	"github.com/ymiz/niconico-search-client/live/targets"
	"net/url"
	"strconv"
)

type SearchParameter struct {
	Keywords  string
	Targets   targets.Targets
	Fields    fields.Fields
	Filters   []filters.Filter
	Sort      sort.Sort
	Offset    int
	Limit     int
	Context   string
	UserAgent string
}

func (s SearchParameter) GenerateUrlValues() url.Values {
	values := url.Values{}
	values.Set("q", s.Keywords)
	values.Set("targets", s.Targets.ToString())
	values.Set("_sort", s.Sort.ToString())
	values.Set("_context", s.Context)
	if s.Fields.HasField() {
		values.Set("fields", s.Fields.ToString())
	}
	if len(s.Filters) > 0 {
		for _, filter := range s.Filters {
			values.Set(filter.GenerateKeyString(), filter.Value)
		}
	}
	values.Set("_offset", strconv.Itoa(s.Offset))
	if s.Limit > 0 {
		values.Set("_limit", strconv.Itoa(s.Limit))
	}
	return values
}

func Find(param SearchParameter) (*common.Response, error) {
	urlValues := param.GenerateUrlValues()
	return common.Get(common.LiveEndPoint, param.UserAgent, urlValues)
}
