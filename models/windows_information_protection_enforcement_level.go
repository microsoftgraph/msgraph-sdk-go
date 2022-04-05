package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the deviceAppManagement singleton.
type WindowsInformationProtectionEnforcementLevel int

const (
    NOPROTECTION_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL WindowsInformationProtectionEnforcementLevel = iota
    ENCRYPTANDAUDITONLY_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL
    ENCRYPTAUDITANDPROMPT_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL
    ENCRYPTAUDITANDBLOCK_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL
)

func (i WindowsInformationProtectionEnforcementLevel) String() string {
    return []string{"NOPROTECTION", "ENCRYPTANDAUDITONLY", "ENCRYPTAUDITANDPROMPT", "ENCRYPTAUDITANDBLOCK"}[i]
}
func ParseWindowsInformationProtectionEnforcementLevel(v string) (interface{}, error) {
    result := NOPROTECTION_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL
    switch strings.ToUpper(v) {
        case "NOPROTECTION":
            result = NOPROTECTION_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL
        case "ENCRYPTANDAUDITONLY":
            result = ENCRYPTANDAUDITONLY_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL
        case "ENCRYPTAUDITANDPROMPT":
            result = ENCRYPTAUDITANDPROMPT_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL
        case "ENCRYPTAUDITANDBLOCK":
            result = ENCRYPTAUDITANDBLOCK_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL
        default:
            return 0, errors.New("Unknown WindowsInformationProtectionEnforcementLevel value: " + v)
    }
    return &result, nil
}
func SerializeWindowsInformationProtectionEnforcementLevel(values []WindowsInformationProtectionEnforcementLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
