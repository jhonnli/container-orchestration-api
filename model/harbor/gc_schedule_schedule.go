/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.6.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package harbor

type GcScheduleSchedule struct {

	// The schedule type. The valid values are daily， weekly and None. 'None' means to cancel the schedule.
	Type_ string `json:"type,omitempty"`

	// Optional, only used when the type is weekly. The valid values are 1-7.
	Weekday int32 `json:"weekday,omitempty"`

	// The time offset with the UTC 00:00 in seconds.
	Offtime int64 `json:"offtime,omitempty"`
}
