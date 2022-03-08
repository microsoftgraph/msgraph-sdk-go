package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the educationRoot singleton.
type EducationAddToCalendarOptions int

const (
    NONE_EDUCATIONADDTOCALENDAROPTIONS EducationAddToCalendarOptions = iota
    STUDENTSANDPUBLISHER_EDUCATIONADDTOCALENDAROPTIONS
    STUDENTSANDTEAMOWNERS_EDUCATIONADDTOCALENDAROPTIONS
    UNKNOWNFUTUREVALUE_EDUCATIONADDTOCALENDAROPTIONS
    STUDENTSONLY_EDUCATIONADDTOCALENDAROPTIONS
)

func (i EducationAddToCalendarOptions) String() string {
    return []string{"NONE", "STUDENTSANDPUBLISHER", "STUDENTSANDTEAMOWNERS", "UNKNOWNFUTUREVALUE", "STUDENTSONLY"}[i]
}
func ParseEducationAddToCalendarOptions(v string) (interface{}, error) {
    result := NONE_EDUCATIONADDTOCALENDAROPTIONS
    switch strings.ToUpper(v) {
        case "NONE":
            result = NONE_EDUCATIONADDTOCALENDAROPTIONS
        case "STUDENTSANDPUBLISHER":
            result = STUDENTSANDPUBLISHER_EDUCATIONADDTOCALENDAROPTIONS
        case "STUDENTSANDTEAMOWNERS":
            result = STUDENTSANDTEAMOWNERS_EDUCATIONADDTOCALENDAROPTIONS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_EDUCATIONADDTOCALENDAROPTIONS
        case "STUDENTSONLY":
            result = STUDENTSONLY_EDUCATIONADDTOCALENDAROPTIONS
        default:
            return 0, errors.New("Unknown EducationAddToCalendarOptions value: " + v)
    }
    return &result, nil
}
func SerializeEducationAddToCalendarOptions(values []EducationAddToCalendarOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
