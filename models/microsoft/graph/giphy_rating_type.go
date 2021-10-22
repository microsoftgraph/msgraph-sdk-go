package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "STRICT":
            return STRICT_GIPHYRATINGTYPE, nil
        case "MODERATE":
            return MODERATE_GIPHYRATINGTYPE, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_GIPHYRATINGTYPE, nil
    }
    return 0, errors.New("Unknown GiphyRatingType value: " + v)
}
func SerializeGiphyRatingType(values []GiphyRatingType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
