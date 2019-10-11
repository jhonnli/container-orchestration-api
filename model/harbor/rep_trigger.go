/*
 * Harbor API
 *
 * These APIs provide services for manipulating Harbor project.
 *
 * API version: 1.6.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package harbor

type RepTrigger struct {

	// The replication policy trigger kind. The valid values are manual, immediate and schedule.
	Kind string `json:"kind,omitempty"`

	ScheduleParam *ScheduleParam `json:"schedule_param,omitempty"`
}
