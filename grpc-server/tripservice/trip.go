package trip

import (
	"context"
	trippb "coolcar/proto/gen/go"
)

// Service implements trip service.
type Service struct{}

// GetTrip TODO
func (*Service) GetTrip(c context.Context, req *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	return &trippb.GetTripResponse{
		Id: req.Id,
		Trip: &trippb.Trip{
			Start:       "abc",
			End:         "def",
			DurationSec: 3600,
			FeeCent:     10000,
			StartPos: &trippb.Location{
				Latitude:  30,
				Longitude: 120,
			},
			EndPos: &trippb.Location{
				Latitude:  35,
				Longitude: 120,
			},
			PathLocations: []*trippb.Location{
				{
					Latitude:  35,
					Longitude: 120,
				},
				{
					Latitude:  35,
					Longitude: 120,
				},
			},
			Status: trippb.TripStatus_IN_PROGRESS,
		},
	}, nil
}
