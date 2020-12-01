package anxcloud

import (
	"testing"

	"github.com/anexia-it/go-anxcloud/pkg/ipam/prefix"
	"github.com/google/go-cmp/cmp"
)

// expanders tests

// flatteners tests

func TestFlattenNetworkPrefixLocations(t *testing.T) {
	cases := []struct {
		Input          []prefix.Location
		ExpectedOutput []interface{}
	}{
		{
			[]prefix.Location{
				{
					ID:        "52b5f6b2fd3a4a7eaaedf1a7c0191234",
					Name:      "AT, Vienna, Test",
					Code:      "ANX-8888",
					Country:   "AT",
					Latitude:  "0.0",
					Longitude: "0.0",
					CityCode:  "VIE",
				},
				{
					ID:        "72c5f6b2fd3a4a7eaaedf1a7c0194321",
					Name:      "AT, Vienna, Test2",
					Code:      "ANX-8889",
					Country:   "AT",
					Latitude:  "1.1",
					Longitude: "1.1",
					CityCode:  "VIE",
				},
			},
			[]interface{}{
				map[string]interface{}{
					"identifier": "52b5f6b2fd3a4a7eaaedf1a7c0191234",
					"name":       "AT, Vienna, Test",
					"code":       "ANX-8888",
					"country":    "AT",
					"lat":        "0.0",
					"lon":        "0.0",
					"city_code":  "VIE",
				},
				map[string]interface{}{
					"identifier": "72c5f6b2fd3a4a7eaaedf1a7c0194321",
					"name":       "AT, Vienna, Test2",
					"code":       "ANX-8889",
					"country":    "AT",
					"lat":        "1.1",
					"lon":        "1.1",
					"city_code":  "VIE",
				},
			},
		},
		{
			[]prefix.Location{},
			[]interface{}{},
		},
	}

	for _, tc := range cases {
		output := flattenNetworkPrefixLocations(tc.Input)
		if diff := cmp.Diff(tc.ExpectedOutput, output); diff != "" {
			t.Fatalf("Unexpected output from expander: mismatch (-want +got):\n%s", diff)
		}
	}
}
