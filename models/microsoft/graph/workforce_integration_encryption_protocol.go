package graph
import (
    "strings"
    "errors"
)
type WorkforceIntegrationEncryptionProtocol int

const (
    SHAREDSECRET_WORKFORCEINTEGRATIONENCRYPTIONPROTOCOL WorkforceIntegrationEncryptionProtocol = iota
    UNKNOWNFUTUREVALUE_WORKFORCEINTEGRATIONENCRYPTIONPROTOCOL
)

func (i WorkforceIntegrationEncryptionProtocol) String() string {
    return []string{"SHAREDSECRET", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseWorkforceIntegrationEncryptionProtocol(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "SHAREDSECRET":
            return SHAREDSECRET_WORKFORCEINTEGRATIONENCRYPTIONPROTOCOL, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_WORKFORCEINTEGRATIONENCRYPTIONPROTOCOL, nil
    }
    return 0, errors.New("Unknown WorkforceIntegrationEncryptionProtocol value: " + v)
}
func SerializeWorkforceIntegrationEncryptionProtocol(values []WorkforceIntegrationEncryptionProtocol) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
