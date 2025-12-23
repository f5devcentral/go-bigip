package bigip

import (
	"context"
	"net/http"
	"time"

	"github.com/stretchr/testify/assert"
)

// Utilizes the NetTestSuite from net_test.go

func (s *NetTestSuite) TestTrafficMatchingCriteria() {
	s.ResponseFunc = func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
  "kind": "tm:ltm:traffic-matching-criteria:traffic-matching-criteriacollectionstate",
	"selfLink": "https://localhost/mgmt/tm/ltm/traffic-matching-criteria?ver=1.2.3.4",
	"items": [
	  {
			"kind": "tm:ltm:traffic-matching-criteria:traffic-matching-criteriastate",
			"name": "tmc-foo",
			"partition": "/Common",
			"fullPatch": "/Common/tmc-foo",
			"generation": 1,
			"selfLink": "https://localhost/mgmt/tm/ltm/traffic-matching-criteria/~Common~tmc-foo?ver=1.2.3.4",
			"destinationAddressInline": "192.168.1.100",
			"destinationPortInline": "443",
			"protocol": "tcp",
			"routeDomain": "any",
			"sourceAddressInline": "0.0.0.0",
			"sourceAddressList": "/Common/addresslist-foo",
			"sourceAddressListReference": {
			  "link": "https://localhost/mgmt/tm/net/address-list/~Common~addresslist-foo?ver=1.2.3.4"
			},
			"sourcePortInline": 0
		},
	  {
			"kind": "tm:ltm:traffic-matching-criteria:traffic-matching-criteriastate",
			"name": "tmc-bar",
			"partition": "/Common",
			"fullPatch": "/Common/tmc-bar",
			"generation": 1,
			"selfLink": "https://localhost/mgmt/tm/ltm/traffic-matching-criteria/~Common~tmc-bar?ver=1.2.3.4",
			"destinationAddressInline": "192.168.1.100",
			"destinationPortInline": "443",
			"protocol": "tcp",
			"routeDomain": "any",
			"sourceAddressInline": "0.0.0.0",
			"sourceAddressList": "/Common/addresslist-foo",
			"sourceAddressListReference": {
			  "link": "https://localhost/mgmt/tm/net/address-list/~Common~addresslist-foo?ver=1.2.3.4"
			},
			"sourcePortInline": 0
		}]}`))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	trafficMatchingCriterias, err := s.Client.TrafficMatchingCriterias(ctx)

	assert.Nil(s.T(), err)
	assertRestCall(s, "GET", "/mgmt/tm/ltm/traffic-matching-criteria", "")
	assert.Equal(s.T(), 2, len(trafficMatchingCriterias.TrafficMatchingCriterias))
	assert.Equal(s.T(), "tmc-foo", trafficMatchingCriterias.TrafficMatchingCriterias[0].Name)
	assert.Equal(s.T(), "tmc-bar", trafficMatchingCriterias.TrafficMatchingCriterias[1].Name)
}

func (s *NetTestSuite) TestGetTrafficMatchingCriteria() {
	s.ResponseFunc = func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
			"kind": "tm:ltm:traffic-matching-criteria:traffic-matching-criteriastate",
			"name": "tmc-foo",
			"partition": "/Common",
			"fullPatch": "/Common/tmc-foo",
			"generation": 1,
			"selfLink": "https://localhost/mgmt/tm/ltm/traffic-matching-criteria/~Common~tmc-foo?ver=1.2.3.4",
			"destinationAddressInline": "192.168.1.100",
			"destinationPortInline": "443",
			"protocol": "tcp",
			"routeDomain": "any",
			"sourceAddressInline": "0.0.0.0",
			"sourceAddressList": "/Common/addresslist-foo",
			"sourceAddressListReference": {
			  "link": "https://localhost/mgmt/tm/net/address-list/~Common~addresslist-foo?ver=1.2.3.4"
			},
			"sourcePortInline": 0
		}`))
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	trafficMatchingCriteria, err := s.Client.GetTrafficMatchingCriteria(ctx, "tmc-foo")

	assert.Nil(s.T(), err)
	assertRestCall(s, "GET", "/mgmt/tm/ltm/traffic-matching-criteria/tmc-foo", "")
	assert.Equal(s.T(), "tmc-foo", trafficMatchingCriteria.Name)
	assert.Equal(s.T(), "/Common/addresslist-foo", trafficMatchingCriteria.SourceAddressList)
}

func (s *NetTestSuite) TestAddTrafficMatchingCriteria() {
	someTrafficMatchingCriteria := &TrafficMatchingCriteria{
		Name: "tmc-foo",
		Partition: "Common",
		DestinationAddressInline: "192.168.1.100",
		DestinationPortInline: "443",
		Protocol: "tcp",
		SourceAddressInline: "0.0.0.0",
		SourceAddressList: "/Common/addresslist-foo",
		SourcePortInline: 0,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.Client.AddTrafficMatchingCriteria(ctx, someTrafficMatchingCriteria)

	assert.Nil(s.T(), err)
	assertRestCall(s, "POST", "/mgmt/tm/ltm/traffic-matching-criteria", `{"name":"tmc-foo","partition":"Common","destinationAddressInline":"192.168.1.100","destinationPortInline":"443","protocol":"tcp","sourceAddressInline":"0.0.0.0","sourceAddressList":"/Common/addresslist-foo"}`)
}

func (s *NetTestSuite) TestModifyTrafficMatchingCriteria() {
	someTrafficMatchingCriteriaMod := &TrafficMatchingCriteria{
		SourceAddressList: "/Common/addresslist-bar",
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.Client.ModifyTrafficMatchingCriteria(ctx, "tmc-foo", someTrafficMatchingCriteriaMod)

	assert.Nil(s.T(), err)
	assertRestCall(s, "PATCH", "/mgmt/tm/ltm/traffic-matching-criteria/tmc-foo", `{"sourceAddressList":"/Common/addresslist-bar"}`)
}

func (s *NetTestSuite) TestDeleteTrafficMatchingCriteria() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.Client.DeleteTrafficMatchingCriteria(ctx, "tmc-foo")

	assert.Nil(s.T(), err)
	assertRestCall(s, "DELETE", "/mgmt/tm/ltm/traffic-matching-criteria/tmc-foo", "")
}

