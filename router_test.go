package yno

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchRouterRequestStructure(t *testing.T) {
	testCases := []string{
		`{"Query":{"Where":{"$eq":{"SerialNumber":"M5B123456"}}}}`,
		`{"Query":{"Where":{"$pm":{"ModelName":"RTX"}}}}`,
		`{"Query":{"Where":{"$in":{"ModelName":["RTX830","RTX1210"]}}}}`,
		`{"Query":{"Where":{"$inArray":{"AssignedUsers":["customer01","customer02"]}}}}`,
		`{"Query":{"Where":{"$pmInArray":{"AssignedUsers":["customer","user"]}}}}`,
		`{"Query":{"Where":{"$and":[{"$eq":{"DeviceStatus":"Online"}},{"$in":{"ModelName":["RTX830","RTX1210"]}}]}}}`,
		`{"Query":{"Where":{"$or":[{"$inArray":{"AssignedLabels":["浜松","東京"]}},{"$pm":{"DeviceDescription":"device"}}]}}}`,
		`{"Query":{"Where":{"$or":[{"$and":[{"$inArray":{"AssignedUsers":["user1","user2"]}},{"$eq":{"DeviceDescription":"東京本社"}}]},{"$and":[{"$inArray":{"AssignedUsers":["user"]}},{"$eq":{"DeviceDescription":"浜松支店"}}]}]}}}`,
		`{"PageToken":"1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890ABCDEFGHIJKLMNOPQRSTUVWXYZ12345"}`,
	}

	for _, tc := range testCases {
		var actual SearchRouterRequest
		err := json.Unmarshal([]byte(tc), &actual)
		assert.NoError(t, err)

		result, err := json.Marshal(actual)
		assert.NoError(t, err)
		assert.Equal(t, []byte(tc), result)
	}
}
