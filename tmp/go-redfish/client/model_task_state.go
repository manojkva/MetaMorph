/*
 * Redfish OAPI specification
 *
 * Partial Redfish OAPI specification for a limited client
 *
 * API version: 0.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// TaskState the model 'TaskState'
type TaskState string

// List of TaskState
const (
	TASKSTATE_NEW TaskState = "New"
	TASKSTATE_STARTING TaskState = "Starting"
	TASKSTATE_RUNNING TaskState = "Running"
	TASKSTATE_SUSPENDED TaskState = "Suspended"
	TASKSTATE_INTERRUPTED TaskState = "Interrupted"
	TASKSTATE_PENDING TaskState = "Pending"
	TASKSTATE_STOPPING TaskState = "Stopping"
	TASKSTATE_COMPLETED TaskState = "Completed"
	TASKSTATE_KILLED TaskState = "Killed"
	TASKSTATE_EXCEPTION TaskState = "Exception"
	TASKSTATE_SERVICE TaskState = "Service"
	TASKSTATE_CANCELLING TaskState = "Cancelling"
	TASKSTATE_CANCELLED TaskState = "Cancelled"
)
