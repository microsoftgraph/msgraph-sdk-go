package models
import (
    "errors"
)
type MultiTenantOrganizationMemberRole int

const (
    OWNER_MULTITENANTORGANIZATIONMEMBERROLE MultiTenantOrganizationMemberRole = iota
    MEMBER_MULTITENANTORGANIZATIONMEMBERROLE
    UNKNOWNFUTUREVALUE_MULTITENANTORGANIZATIONMEMBERROLE
)

func (i MultiTenantOrganizationMemberRole) String() string {
    return []string{"owner", "member", "unknownFutureValue"}[i]
}
func ParseMultiTenantOrganizationMemberRole(v string) (any, error) {
    result := OWNER_MULTITENANTORGANIZATIONMEMBERROLE
    switch v {
        case "owner":
            result = OWNER_MULTITENANTORGANIZATIONMEMBERROLE
        case "member":
            result = MEMBER_MULTITENANTORGANIZATIONMEMBERROLE
        case "unknownFutureValue":
            result = UNKNOWNFUTUREVALUE_MULTITENANTORGANIZATIONMEMBERROLE
        default:
            return 0, errors.New("Unknown MultiTenantOrganizationMemberRole value: " + v)
    }
    return &result, nil
}
func SerializeMultiTenantOrganizationMemberRole(values []MultiTenantOrganizationMemberRole) []string {
    result := make([]string, len(values))
    for i, v := range values {
        result[i] = v.String()
    }
    return result
}
func (i MultiTenantOrganizationMemberRole) isMultiValue() bool {
    return false
}
