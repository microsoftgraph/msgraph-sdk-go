package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the security singleton.
type FileHashType int

const (
    UNKNOWN_FILEHASHTYPE FileHashType = iota
    SHA1_FILEHASHTYPE
    SHA256_FILEHASHTYPE
    MD5_FILEHASHTYPE
    AUTHENTICODEHASH256_FILEHASHTYPE
    LSHASH_FILEHASHTYPE
    CTPH_FILEHASHTYPE
    UNKNOWNFUTUREVALUE_FILEHASHTYPE
)

func (i FileHashType) String() string {
    return []string{"UNKNOWN", "SHA1", "SHA256", "MD5", "AUTHENTICODEHASH256", "LSHASH", "CTPH", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseFileHashType(v string) (interface{}, error) {
    result := UNKNOWN_FILEHASHTYPE
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            result = UNKNOWN_FILEHASHTYPE
        case "SHA1":
            result = SHA1_FILEHASHTYPE
        case "SHA256":
            result = SHA256_FILEHASHTYPE
        case "MD5":
            result = MD5_FILEHASHTYPE
        case "AUTHENTICODEHASH256":
            result = AUTHENTICODEHASH256_FILEHASHTYPE
        case "LSHASH":
            result = LSHASH_FILEHASHTYPE
        case "CTPH":
            result = CTPH_FILEHASHTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_FILEHASHTYPE
        default:
            return 0, errors.New("Unknown FileHashType value: " + v)
    }
    return &result, nil
}
func SerializeFileHashType(values []FileHashType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
