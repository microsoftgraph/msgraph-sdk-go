package models
import (
    "strings"
    "errors"
)
// Provides operations to manage the policyRoot singleton.
type UnifiedRoleManagementPolicyRuleTargetOperations int

const (
    ALL_UNIFIEDROLEMANAGEMENTPOLICYRULETARGETOPERATIONS UnifiedRoleManagementPolicyRuleTargetOperations = iota
    ACTIVATE_UNIFIEDROLEMANAGEMENTPOLICYRULETARGETOPERATIONS
    DEACTIVATE_UNIFIEDROLEMANAGEMENTPOLICYRULETARGETOPERATIONS
    ASSIGN_UNIFIEDROLEMANAGEMENTPOLICYRULETARGETOPERATIONS
    UPDATE_UNIFIEDROLEMANAGEMENTPOLICYRULETARGETOPERATIONS
    REMOVE_UNIFIEDROLEMANAGEMENTPOLICYRULETARGETOPERATIONS
    EXTEND_UNIFIEDROLEMANAGEMENTPOLICYRULETARGETOPERATIONS
    RENEW_UNIFIEDROLEMANAGEMENTPOLICYRULETARGETOPERATIONS
    UNKNOWNFUTUREVALUE_UNIFIEDROLEMANAGEMENTPOLICYRULETARGETOPERATIONS
)

func (i UnifiedRoleManagementPolicyRuleTargetOperations) String() string {
    return []string{"ALL", "ACTIVATE", "DEACTIVATE", "ASSIGN", "UPDATE", "REMOVE", "EXTEND", "RENEW", "UNKNOWNFUTUREVALUE"}[i]
}
func ParseUnifiedRoleManagementPolicyRuleTargetOperations(v string) (interface{}, error) {
    result := ALL_UNIFIEDROLEMANAGEMENTPOLICYRULETARGETOPERATIONS
    switch strings.ToUpper(v) {
        case "ALL":
            result = ALL_UNIFIEDROLEMANAGEMENTPOLICYRULETARGETOPERATIONS
        case "ACTIVATE":
            result = ACTIVATE_UNIFIEDROLEMANAGEMENTPOLICYRULETARGETOPERATIONS
        case "DEACTIVATE":
            result = DEACTIVATE_UNIFIEDROLEMANAGEMENTPOLICYRULETARGETOPERATIONS
        case "ASSIGN":
            result = ASSIGN_UNIFIEDROLEMANAGEMENTPOLICYRULETARGETOPERATIONS
        case "UPDATE":
            result = UPDATE_UNIFIEDROLEMANAGEMENTPOLICYRULETARGETOPERATIONS
        case "REMOVE":
            result = REMOVE_UNIFIEDROLEMANAGEMENTPOLICYRULETARGETOPERATIONS
        case "EXTEND":
            result = EXTEND_UNIFIEDROLEMANAGEMENTPOLICYRULETARGETOPERATIONS
        case "RENEW":
            result = RENEW_UNIFIEDROLEMANAGEMENTPOLICYRULETARGETOPERATIONS
        case "UNKNOWNFUTUREVALUE":
            result = UNKNOWNFUTUREVALUE_UNIFIEDROLEMANAGEMENTPOLICYRULETARGETOPERATIONS
        default:
            return 0, errors.New("Unknown UnifiedRoleManagementPolicyRuleTargetOperations value: " + v)
    }
    return &result, nil
}
func SerializeUnifiedRoleManagementPolicyRuleTargetOperations(values []UnifiedRoleManagementPolicyRuleTargetOperations) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
