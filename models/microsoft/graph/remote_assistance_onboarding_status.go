package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NOTONBOARDED":
            return NOTONBOARDED_REMOTEASSISTANCEONBOARDINGSTATUS, nil
        case "ONBOARDING":
            return ONBOARDING_REMOTEASSISTANCEONBOARDINGSTATUS, nil
        case "ONBOARDED":
            return ONBOARDED_REMOTEASSISTANCEONBOARDINGSTATUS, nil
    }
    return 0, errors.New("Unknown RemoteAssistanceOnboardingStatus value: " + v)
}
func SerializeRemoteAssistanceOnboardingStatus(values []RemoteAssistanceOnboardingStatus) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
