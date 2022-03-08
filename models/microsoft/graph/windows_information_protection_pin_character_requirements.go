package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceAppManagement singleton.
type WindowsInformationProtectionPinCharacterRequirements int

const (
    NOTALLOW_WINDOWSINFORMATIONPROTECTIONPINCHARACTERREQUIREMENTS WindowsInformationProtectionPinCharacterRequirements = iota
    REQUIREATLEASTONE_WINDOWSINFORMATIONPROTECTIONPINCHARACTERREQUIREMENTS
    ALLOW_WINDOWSINFORMATIONPROTECTIONPINCHARACTERREQUIREMENTS
)

func (i WindowsInformationProtectionPinCharacterRequirements) String() string {
    return []string{"NOTALLOW", "REQUIREATLEASTONE", "ALLOW"}[i]
}
func ParseWindowsInformationProtectionPinCharacterRequirements(v string) (interface{}, error) {
    result := NOTALLOW_WINDOWSINFORMATIONPROTECTIONPINCHARACTERREQUIREMENTS
    switch strings.ToUpper(v) {
        case "NOTALLOW":
            result = NOTALLOW_WINDOWSINFORMATIONPROTECTIONPINCHARACTERREQUIREMENTS
        case "REQUIREATLEASTONE":
            result = REQUIREATLEASTONE_WINDOWSINFORMATIONPROTECTIONPINCHARACTERREQUIREMENTS
        case "ALLOW":
            result = ALLOW_WINDOWSINFORMATIONPROTECTIONPINCHARACTERREQUIREMENTS
        default:
            return 0, errors.New("Unknown WindowsInformationProtectionPinCharacterRequirements value: " + v)
    }
    return &result, nil
}
func SerializeWindowsInformationProtectionPinCharacterRequirements(values []WindowsInformationProtectionPinCharacterRequirements) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
