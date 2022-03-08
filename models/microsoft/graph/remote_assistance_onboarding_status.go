package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceManagement singleton.
type RemoteAssistanceOnboardingStatus int

const (
    NOTONBOARDED_REMOTEASSISTANCEONBOARDINGSTATUS RemoteAssistanceOnboardingStatus = iota
    ONBOARDING_REMOTEASSISTANCEONBOARDINGSTATUS
    ONBOARDED_REMOTEASSISTANCEONBOARDINGSTATUS
)

func (i RemoteAssistanceOnboardingStatus) String() string {
    return []string{"NOTONBOARDED", "ONBOARDING", "ONBOARDED"}[i]
}
func ParseRemoteAssistanceOnboardingStatus(v string) (interface{}, error) {
    result := NOTONBOARDED_REMOTEASSISTANCEONBOARDINGSTATUS
    switch strings.ToUpper(v) {
        case "NOTONBOARDED":
            result = NOTONBOARDED_REMOTEASSISTANCEONBOARDINGSTATUS
        case "ONBOARDING":
            result = ONBOARDING_REMOTEASSISTANCEONBOARDINGSTATUS
        case "ONBOARDED":
            result = ONBOARDED_REMOTEASSISTANCEONBOARDINGSTATUS
        default:
            return 0, errors.New("Unknown RemoteAssistanceOnboardingStatus value: " + v)
    }
    return &result, nil
}
func SerializeRemoteAssistanceOnboardingStatus(values []RemoteAssistanceOnboardingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
