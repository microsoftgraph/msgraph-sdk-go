package graph
import (
    "strings"
    "errors"
)
// 
type TeamworkActivityTopicSource int

const (
    ENTITYURL_TEAMWORKACTIVITYTOPICSOURCE TeamworkActivityTopicSource = iota
    TEXT_TEAMWORKACTIVITYTOPICSOURCE
)

func (i TeamworkActivityTopicSource) String() string {
    return []string{"ENTITYURL", "TEXT"}[i]
}
func ParseTeamworkActivityTopicSource(v string) (interface{}, error) {
    switch strings.ToUpper(v) {
        case "ENTITYURL":
            return ENTITYURL_TEAMWORKACTIVITYTOPICSOURCE, nil
        case "TEXT":
            return TEXT_TEAMWORKACTIVITYTOPICSOURCE, nil
    }
    return 0, errors.New("Unknown TeamworkActivityTopicSource value: " + v)
}
func SerializeTeamworkActivityTopicSource(values []TeamworkActivityTopicSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
