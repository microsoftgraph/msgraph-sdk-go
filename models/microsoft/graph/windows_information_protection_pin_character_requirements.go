package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NOTALLOW":
            return NOTALLOW_WINDOWSINFORMATIONPROTECTIONPINCHARACTERREQUIREMENTS, nil
        case "REQUIREATLEASTONE":
            return REQUIREATLEASTONE_WINDOWSINFORMATIONPROTECTIONPINCHARACTERREQUIREMENTS, nil
        case "ALLOW":
            return ALLOW_WINDOWSINFORMATIONPROTECTIONPINCHARACTERREQUIREMENTS, nil
    }
    return 0, errors.New("Unknown WindowsInformationProtectionPinCharacterRequirements value: " + v)
}
func SerializeWindowsInformationProtectionPinCharacterRequirements(values []WindowsInformationProtectionPinCharacterRequirements) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
