package graph
import (
    "strings"
    "errors"
)
// 
type MediaDirection int

const (
    INACTIVE_MEDIADIRECTION MediaDirection = iota
    SENDONLY_MEDIADIRECTION
    RECEIVEONLY_MEDIADIRECTION
    SENDRECEIVE_MEDIADIRECTION
)

func (i MediaDirection) String() string {
    return []string{"INACTIVE", "SENDONLY", "RECEIVEONLY", "SENDRECEIVE"}[i]
}
func ParseMediaDirection(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "INACTIVE":
            return INACTIVE_MEDIADIRECTION, nil
        case "SENDONLY":
            return SENDONLY_MEDIADIRECTION, nil
        case "RECEIVEONLY":
            return RECEIVEONLY_MEDIADIRECTION, nil
        case "SENDRECEIVE":
            return SENDRECEIVE_MEDIADIRECTION, nil
    }
    return 0, errors.New("Unknown MediaDirection value: " + v)
}
func SerializeMediaDirection(values []MediaDirection) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
