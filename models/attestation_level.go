package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the drive singleton.
type AttestationLevel int

const (
    ATTESTED_ATTESTATIONLEVEL AttestationLevel = iota
    NOTATTESTED_ATTESTATIONLEVEL
    UNKNOWNFUTUREVALUE_ATTESTATIONLEVEL
)

func (i AttestationLevel) String() string {
    return []string{"ATTESTED", "NOTATTESTED", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseAttestationLevel(v string) (interface{}, error) {
    result := ATTESTED_ATTESTATIONLEVEL
    switch strings.ToUpper(v) {
        case "ATTESTED":
            result = ATTESTED_ATTESTATIONLEVEL
        case "NOTATTESTED":
            result = NOTATTESTED_ATTESTATIONLEVEL
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_ATTESTATIONLEVEL
        default:
            return 0, errors.New("Unknown AttestationLevel value: " + v)
    }
    return &result, nil
}
func SerializeAttestationLevel(values []AttestationLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
