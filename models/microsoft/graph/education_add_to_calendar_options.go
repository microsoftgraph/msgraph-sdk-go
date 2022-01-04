package graph
import (
    "strings"
    "errors"
)
// 
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
    switch strings.ToUpper(v) {
        case "NONE":
            return NONE_EDUCATIONADDTOCALENDAROPTIONS, nil
        case "STUDENTSANDPUBLISHER":
            return STUDENTSANDPUBLISHER_EDUCATIONADDTOCALENDAROPTIONS, nil
        case "STUDENTSANDTEAMOWNERS":
            return STUDENTSANDTEAMOWNERS_EDUCATIONADDTOCALENDAROPTIONS, nil
        case "UNKNOWNFUTUREVALUE":
            return UNKNOWNFUTUREVALUE_EDUCATIONADDTOCALENDAROPTIONS, nil
        case "STUDENTSONLY":
            return STUDENTSONLY_EDUCATIONADDTOCALENDAROPTIONS, nil
    }
    return 0, errors.New("Unknown EducationAddToCalendarOptions value: " + v)
}
func SerializeEducationAddToCalendarOptions(values []EducationAddToCalendarOptions) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
