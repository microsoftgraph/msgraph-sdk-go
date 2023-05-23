package models
import (
    "errors"
)
// 
type SynchronizationTaskExecutionResult int

const (
    SUCCEEDED_SYNCHRONIZATIONTASKEXECUTIONRESULT SynchronizationTaskExecutionResult = iota
    FAILED_SYNCHRONIZATIONTASKEXECUTIONRESULT
    ENTRYLEVELERRORS_SYNCHRONIZATIONTASKEXECUTIONRESULT
)

func (i SynchronizationTaskExecutionResult) String() string {
    return []string{"Succeeded", "Failed", "EntryLevelErrors"}[i]
}
func ParseSynchronizationTaskExecutionResult(v string) (any, error) {
    result := SUCCEEDED_SYNCHRONIZATIONTASKEXECUTIONRESULT
    switch v {
        case "Succeeded":
            result = SUCCEEDED_SYNCHRONIZATIONTASKEXECUTIONRESULT
        case "Failed":
            result = FAILED_SYNCHRONIZATIONTASKEXECUTIONRESULT
        case "EntryLevelErrors":
            result = ENTRYLEVELERRORS_SYNCHRONIZATIONTASKEXECUTIONRESULT
        default:
            return 0, errors.New("Unknown SynchronizationTaskExecutionResult value: " + v)
    }
    return &result, nil
}
func SerializeSynchronizationTaskExecutionResult(values []SynchronizationTaskExecutionResult) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
