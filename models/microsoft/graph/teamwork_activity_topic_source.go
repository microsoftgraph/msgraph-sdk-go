package graph
import (
    "strings"
    "errors"
)
// Provides operations to call the sendActivityNotification method.
type TeamworkActivityTopicSource int

const (
    ENTITYURL_TEAMWORKACTIVITYTOPICSOURCE TeamworkActivityTopicSource = iota
    TEXT_TEAMWORKACTIVITYTOPICSOURCE
)

func (i TeamworkActivityTopicSource) String() string {
    return []string{"ENTITYURL", "TEXT"}[i]
}
func ParseTeamworkActivityTopicSource(v string) (interface{}, error) {
    result := ENTITYURL_TEAMWORKACTIVITYTOPICSOURCE
    switch strings.ToUpper(v) {
        case "ENTITYURL":
            result = ENTITYURL_TEAMWORKACTIVITYTOPICSOURCE
        case "TEXT":
            result = TEXT_TEAMWORKACTIVITYTOPICSOURCE
        default:
            return 0, errors.New("Unknown TeamworkActivityTopicSource value: " + v)
    }
    return &result, nil
}
func SerializeTeamworkActivityTopicSource(values []TeamworkActivityTopicSource) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
