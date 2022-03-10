package graph
import (
    "strings"
    "errors"
)
// Provides operations to manage the identityContainer singleton.
type IdentityUserFlowAttributeInputType int

const (
    TEXTBOX_IDENTITYUSERFLOWATTRIBUTEINPUTTYPE IdentityUserFlowAttributeInputType = iota
    DATETIMEDROPDOWN_IDENTITYUSERFLOWATTRIBUTEINPUTTYPE
    RADIOSINGLESELECT_IDENTITYUSERFLOWATTRIBUTEINPUTTYPE
    DROPDOWNSINGLESELECT_IDENTITYUSERFLOWATTRIBUTEINPUTTYPE
    EMAILBOX_IDENTITYUSERFLOWATTRIBUTEINPUTTYPE
    CHECKBOXMULTISELECT_IDENTITYUSERFLOWATTRIBUTEINPUTTYPE
)

func (i IdentityUserFlowAttributeInputType) String() string {
    return []string{"TEXTBOX", "DATETIMEDROPDOWN", "RADIOSINGLESELECT", "DROPDOWNSINGLESELECT", "EMAILBOX", "CHECKBOXMULTISELECT"}[i]
}
func ParseIdentityUserFlowAttributeInputType(v string) (interface{}, error) {
    result := TEXTBOX_IDENTITYUSERFLOWATTRIBUTEINPUTTYPE
    switch strings.ToUpper(v) {
        case "TEXTBOX":
            result = TEXTBOX_IDENTITYUSERFLOWATTRIBUTEINPUTTYPE
        case "DATETIMEDROPDOWN":
            result = DATETIMEDROPDOWN_IDENTITYUSERFLOWATTRIBUTEINPUTTYPE
        case "RADIOSINGLESELECT":
            result = RADIOSINGLESELECT_IDENTITYUSERFLOWATTRIBUTEINPUTTYPE
        case "DROPDOWNSINGLESELECT":
            result = DROPDOWNSINGLESELECT_IDENTITYUSERFLOWATTRIBUTEINPUTTYPE
        case "EMAILBOX":
            result = EMAILBOX_IDENTITYUSERFLOWATTRIBUTEINPUTTYPE
        case "CHECKBOXMULTISELECT":
            result = CHECKBOXMULTISELECT_IDENTITYUSERFLOWATTRIBUTEINPUTTYPE
        default:
            return 0, errors.New("Unknown IdentityUserFlowAttributeInputType value: " + v)
    }
    return &result, nil
}
func SerializeIdentityUserFlowAttributeInputType(values []IdentityUserFlowAttributeInputType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
