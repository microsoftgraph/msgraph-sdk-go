package graph
import (
    "strings"
    "errors"
)
// 
type WorkforceIntegrationEncryptionProtocol int

const (
    SHAREDSECRET_WORKFORCEINTEGRATIONENCRYPTIONPROTOCOL WorkforceIntegrationEncryptionProtocol = iota
    UNKNOWNFUTUREVALUE_WORKFORCEINTEGRATIONENCRYPTIONPROTOCOL
)

func (i WorkforceIntegrationEncryptionProtocol) String() string {
    return []string{"SHAREDSECRET", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseWorkforceIntegrationEncryptionProtocol(v string) (interface{}, error) {
    result := SHAREDSECRET_WORKFORCEINTEGRATIONENCRYPTIONPROTOCOL
    switch strings.ToUpper(v) {
        case "SHAREDSECRET":
            result = SHAREDSECRET_WORKFORCEINTEGRATIONENCRYPTIONPROTOCOL
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_WORKFORCEINTEGRATIONENCRYPTIONPROTOCOL
        default:
            return 0, errors.New("Unknown WorkforceIntegrationEncryptionProtocol value: " + v)
    }
    return &result, nil
}
func SerializeWorkforceIntegrationEncryptionProtocol(values []WorkforceIntegrationEncryptionProtocol) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
