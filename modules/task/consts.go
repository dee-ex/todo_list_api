package task

import (
    "errors"
)

var SERVICE_ERROR_INVALID_TASK_ID = errors.New("Task id is invalid")

var SERVICE_ERROR_INTERNAL_SERVER_ERROR = errors.New("Internal server errror")

var SERVICE_ERROR_INVALID_OWNER = errors.New("You do not own this task")
