package graph
import (
    "strings"
    "errors"
)
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
    switch strings.ToUpper(v) {
        case "TEXTBOX":
            return TEXTBOX_IDENTITYUSERFLOWATTRIBUTEINPUTTYPE, nil
        case "DATETIMEDROPDOWN":
            return DATETIMEDROPDOWN_IDENTITYUSERFLOWATTRIBUTEINPUTTYPE, nil
        case "RADIOSINGLESELECT":
            return RADIOSINGLESELECT_IDENTITYUSERFLOWATTRIBUTEINPUTTYPE, nil
        case "DROPDOWNSINGLESELECT":
            return DROPDOWNSINGLESELECT_IDENTITYUSERFLOWATTRIBUTEINPUTTYPE, nil
        case "EMAILBOX":
            return EMAILBOX_IDENTITYUSERFLOWATTRIBUTEINPUTTYPE, nil
        case "CHECKBOXMULTISELECT":
            return CHECKBOXMULTISELECT_IDENTITYUSERFLOWATTRIBUTEINPUTTYPE, nil
    }
    return 0, errors.New("Unknown IdentityUserFlowAttributeInputType value: " + v)
}
func SerializeIdentityUserFlowAttributeInputType(values []IdentityUserFlowAttributeInputType) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
