package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "NOPROTECTION":
            return NOPROTECTION_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL, nil
        case "ENCRYPTANDAUDITONLY":
            return ENCRYPTANDAUDITONLY_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL, nil
        case "ENCRYPTAUDITANDPROMPT":
            return ENCRYPTAUDITANDPROMPT_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL, nil
        case "ENCRYPTAUDITANDBLOCK":
            return ENCRYPTAUDITANDBLOCK_WINDOWSINFORMATIONPROTECTIONENFORCEMENTLEVEL, nil
    }
    return 0, errors.New("Unknown WindowsInformationProtectionEnforcementLevel value: " + v)
}
func SerializeWindowsInformationProtectionEnforcementLevel(values []WindowsInformationProtectionEnforcementLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
