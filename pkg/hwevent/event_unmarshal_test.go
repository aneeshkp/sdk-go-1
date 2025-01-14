package hwevent_test

import (
	"encoding/json"

	"testing"
	"time"

	"github.com/redhat-cne/sdk-go/pkg/hwevent"
	"github.com/stretchr/testify/require"

	"github.com/google/go-cmp/cmp"
	"github.com/redhat-cne/sdk-go/pkg/types"
)

func TestUnMarshal(t *testing.T) {
	now := types.Timestamp{Time: time.Now().UTC()}
	_type := "HW_EVENT"
	version := "v1"
	id := "ABC-1234"

	testCases := map[string]struct {
		body    []byte
		want    *hwevent.Event
		wantErr error
	}{

		"struct Data fan": {
			body: mustJSONMarshal(t, map[string]interface{}{
				"data": map[string]interface{}{
					"data":    JSON_EVENT_TMP0100,
					"version": version,
				},
				"id":         id,
				"time":       now.Format(time.RFC3339Nano),
				"type":       _type,
				"dataSchema": nil,
			}),
			want: &hwevent.Event{
				ID:         id,
				Type:       _type,
				Time:       &now,
				DataSchema: nil,
				Data: &hwevent.Data{
					Version: version,
					Data:    &REDFISH_EVENT_TMP0100,
				},
			},
			wantErr: nil,
		},
	}

	for n, tc := range testCases {
		t.Run(n, func(t *testing.T) {
			got := &hwevent.Event{}
			err := json.Unmarshal(tc.body, got)

			if tc.wantErr != nil || err != nil {
				if diff := cmp.Diff(tc.wantErr, err); diff != "" {
					t.Errorf("unexpected error (-want, +got) = %v", diff)
				}
				return
			}

			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("unexpected event (-want, +got) = %v", diff)
			}
		})
	}
}

func mustJSONMarshal(tb testing.TB, body interface{}) []byte {
	b, err := json.Marshal(body)
	require.NoError(tb, err)
	return b
}
