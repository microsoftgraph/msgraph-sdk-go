package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "ATTESTED":
            return ATTESTED_ATTESTATIONLEVEL, nil
        case "NOTATTESTED":
            return NOTATTESTED_ATTESTATIONLEVEL, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_ATTESTATIONLEVEL, nil
    }
    return 0, errors.New("Unknown AttestationLevel value: " + v)
}
func SerializeAttestationLevel(values []AttestationLevel) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
