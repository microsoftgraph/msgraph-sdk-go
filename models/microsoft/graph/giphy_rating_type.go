package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the collection of group entities.
type GiphyRatingType int

const (
    STRICT_GIPHYRATINGTYPE GiphyRatingType = iota
    MODERATE_GIPHYRATINGTYPE
    UNKNOWNFUTUREVALUE_GIPHYRATINGTYPE
)

func (i GiphyRatingType) String() string {
    return []string{"STRICT", "MODERATE", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseGiphyRatingType(v string) (interface{}, error) {
    result := STRICT_GIPHYRATINGTYPE
    switch strings.ToUpper(v) {
        case "STRICT":
            result = STRICT_GIPHYRATINGTYPE
        case "MODERATE":
            result = MODERATE_GIPHYRATINGTYPE
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_GIPHYRATINGTYPE
        default:
            return 0, errors.New("Unknown GiphyRatingType value: " + v)
    }
    return &result, nil
}
func SerializeGiphyRatingType(values []GiphyRatingType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
