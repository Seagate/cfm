/*
 * Redfish
 *
 * This contains the definition of a Redfish service.
 *
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package redfishapi

type TaskV173TaskState string

// List of TaskV173TaskState
const (
	TASKV173TASKSTATE_NEW         TaskV173TaskState = "New"
	TASKV173TASKSTATE_STARTING    TaskV173TaskState = "Starting"
	TASKV173TASKSTATE_RUNNING     TaskV173TaskState = "Running"
	TASKV173TASKSTATE_SUSPENDED   TaskV173TaskState = "Suspended"
	TASKV173TASKSTATE_INTERRUPTED TaskV173TaskState = "Interrupted"
	TASKV173TASKSTATE_PENDING     TaskV173TaskState = "Pending"
	TASKV173TASKSTATE_STOPPING    TaskV173TaskState = "Stopping"
	TASKV173TASKSTATE_COMPLETED   TaskV173TaskState = "Completed"
	TASKV173TASKSTATE_KILLED      TaskV173TaskState = "Killed"
	TASKV173TASKSTATE_EXCEPTION   TaskV173TaskState = "Exception"
	TASKV173TASKSTATE_SERVICE     TaskV173TaskState = "Service"
	TASKV173TASKSTATE_CANCELLING  TaskV173TaskState = "Cancelling"
	TASKV173TASKSTATE_CANCELLED   TaskV173TaskState = "Cancelled"
)

// AssertTaskV173TaskStateRequired checks if the required fields are not zero-ed
func AssertTaskV173TaskStateRequired(obj TaskV173TaskState) error {
	return nil
}

// AssertTaskV173TaskStateConstraints checks if the values respects the defined constraints
func AssertTaskV173TaskStateConstraints(obj TaskV173TaskState) error {
	return nil
}