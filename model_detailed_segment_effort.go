/*
 * Strava API v3
 *
 * The [Swagger Playground](https://developers.strava.com/playground) is the easiest way to familiarize yourself with the Strava API by submitting HTTP requests and observing the responses before you write any client code. It will show what a response will look like with different endpoints depending on the authorization scope you receive from your athletes. To use the Playground, go to https://www.strava.com/settings/api and change your “Authorization Callback Domain” to developers.strava.com. Please note, we only support Swagger 2.0. There is a known issue where you can only select one scope at a time. For more information, please check the section “client code” at https://developers.strava.com/docs.
 *
 * API version: 3.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package strava

import (
	"time"
)

type DetailedSegmentEffort struct {
	// The unique identifier of this effort
	Id int64 `json:"id,omitempty"`
	// The effort's elapsed time
	ElapsedTime int32 `json:"elapsed_time,omitempty"`
	// The time at which the effort was started.
	StartDate time.Time `json:"start_date,omitempty"`
	// The time at which the effort was started in the local timezone.
	StartDateLocal time.Time `json:"start_date_local,omitempty"`
	// The effort's distance in meters
	Distance float32 `json:"distance,omitempty"`
	// Whether this effort is the current best on the leaderboard
	IsKom bool `json:"is_kom,omitempty"`
	// The name of the segment on which this effort was performed
	Name     string        `json:"name,omitempty"`
	Activity *MetaActivity `json:"activity,omitempty"`
	Athlete  *MetaAthlete  `json:"athlete,omitempty"`
	// The effort's moving time
	MovingTime int32 `json:"moving_time,omitempty"`
	// The start index of this effort in its activity's stream
	StartIndex int32 `json:"start_index,omitempty"`
	// The end index of this effort in its activity's stream
	EndIndex int32 `json:"end_index,omitempty"`
	// The effort's average cadence
	AverageCadence float32 `json:"average_cadence,omitempty"`
	// The average wattage of this effort
	AverageWatts float32 `json:"average_watts,omitempty"`
	// For riding efforts, whether the wattage was reported by a dedicated recording device
	DeviceWatts bool `json:"device_watts,omitempty"`
	// The heart heart rate of the athlete during this effort
	AverageHeartrate float32 `json:"average_heartrate,omitempty"`
	// The maximum heart rate of the athlete during this effort
	MaxHeartrate float32         `json:"max_heartrate,omitempty"`
	Segment      *SummarySegment `json:"segment,omitempty"`
	// The rank of the effort on the global leaderboard if it belongs in the top 10 at the time of upload
	KomRank int32 `json:"kom_rank,omitempty"`
	// The rank of the effort on the athlete's leaderboard if it belongs in the top 3 at the time of upload
	PrRank int32 `json:"pr_rank,omitempty"`
	// Whether this effort should be hidden when viewed within an activity
	Hidden bool `json:"hidden,omitempty"`
}