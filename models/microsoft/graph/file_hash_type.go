package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "UNKNOWN":
            return UNKNOWN_FILEHASHTYPE, nil
        case "SHA1":
            return SHA1_FILEHASHTYPE, nil
        case "SHA256":
            return SHA256_FILEHASHTYPE, nil
        case "MD5":
            return MD5_FILEHASHTYPE, nil
        case "AUTHENTICODEHASH256":
            return AUTHENTICODEHASH256_FILEHASHTYPE, nil
        case "LSHASH":
            return LSHASH_FILEHASHTYPE, nil
        case "CTPH":
            return CTPH_FILEHASHTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_FILEHASHTYPE, nil
    }
    return 0, errors.New("Unknown FileHashType value: " + v)
}
func SerializeFileHashType(values []FileHashType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
