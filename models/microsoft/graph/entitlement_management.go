package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EntitlementManagement provides operations to manage the identityGovernance singleton.
type EntitlementManagement struct {
    Entity
    // Approval stages for assignment requests.
    accessPackageAssignmentApprovals []Approvalable;
    // Represents access package objects.
    accessPackages []AccessPackageable;
    // 
    assignmentPolicies []AccessPackageAssignmentPolicyable;
    // Represents access package assignment requests created by or on behalf of a user.
    assignmentRequests []AccessPackageAssignmentRequestable;
    // Represents the grant of an access package to a subject (user or group).
    assignments []AccessPackageAssignmentable;
    // Represents a group of access packages.
    catalogs []AccessPackageCatalogable;
    // Represents references to a directory or domain of another organization whose users can request access.
    connectedOrganizations []ConnectedOrganizationable;
    // Represents the settings that control the behavior of Azure AD entitlement management.
    settings EntitlementManagementSettingsable;
}
// NewEntitlementManagement instantiates a new entitlementManagement and sets the default values.
func NewEntitlementManagement()(*EntitlementManagement) {
    m := &EntitlementManagement{
        Entity: *NewEntity(),
    }
    return m
}
// CreateEntitlementManagementFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateEntitlementManagementFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewEntitlementManagement(), nil
}
// GetAccessPackageAssignmentApprovals gets the accessPackageAssignmentApprovals property value. Approval stages for assignment requests.
func (m *EntitlementManagement) GetAccessPackageAssignmentApprovals()([]Approvalable) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignmentApprovals
    }
}
// GetAccessPackages gets the accessPackages property value. Represents access package objects.
func (m *EntitlementManagement) GetAccessPackages()([]AccessPackageable) {
    if m == nil {
        return nil
    } else {
        return m.accessPackages
    }
}
// GetAssignmentPolicies gets the assignmentPolicies property value. 
func (m *EntitlementManagement) GetAssignmentPolicies()([]AccessPackageAssignmentPolicyable) {
    if m == nil {
        return nil
    } else {
        return m.assignmentPolicies
    }
}
// GetAssignmentRequests gets the assignmentRequests property value. Represents access package assignment requests created by or on behalf of a user.
func (m *EntitlementManagement) GetAssignmentRequests()([]AccessPackageAssignmentRequestable) {
    if m == nil {
        return nil
    } else {
        return m.assignmentRequests
    }
}
// GetAssignments gets the assignments property value. Represents the grant of an access package to a subject (user or group).
func (m *EntitlementManagement) GetAssignments()([]AccessPackageAssignmentable) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetCatalogs gets the catalogs property value. Represents a group of access packages.
func (m *EntitlementManagement) GetCatalogs()([]AccessPackageCatalogable) {
    if m == nil {
        return nil
    } else {
        return m.catalogs
    }
}
// GetConnectedOrganizations gets the connectedOrganizations property value. Represents references to a directory or domain of another organization whose users can request access.
func (m *EntitlementManagement) GetConnectedOrganizations()([]ConnectedOrganizationable) {
    if m == nil {
        return nil
    } else {
        return m.connectedOrganizations
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EntitlementManagement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageAssignmentApprovals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateApprovalFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Approvalable, len(val))
            for i, v := range val {
                res[i] = v.(Approvalable)
            }
            m.SetAccessPackageAssignmentApprovals(res)
        }
        return nil
    }
    res["accessPackages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageable)
            }
            m.SetAccessPackages(res)
        }
        return nil
    }
    res["assignmentPolicies"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAssignmentPolicyFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignmentPolicyable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageAssignmentPolicyable)
            }
            m.SetAssignmentPolicies(res)
        }
        return nil
    }
    res["assignmentRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAssignmentRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignmentRequestable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageAssignmentRequestable)
            }
            m.SetAssignmentRequests(res)
        }
        return nil
    }
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageAssignmentFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignmentable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageAssignmentable)
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["catalogs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateAccessPackageCatalogFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageCatalogable, len(val))
            for i, v := range val {
                res[i] = v.(AccessPackageCatalogable)
            }
            m.SetCatalogs(res)
        }
        return nil
    }
    res["connectedOrganizations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateConnectedOrganizationFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConnectedOrganizationable, len(val))
            for i, v := range val {
                res[i] = v.(ConnectedOrganizationable)
            }
            m.SetConnectedOrganizations(res)
        }
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateEntitlementManagementSettingsFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(EntitlementManagementSettingsable))
        }
        return nil
    }
    return res
}
// GetSettings gets the settings property value. Represents the settings that control the behavior of Azure AD entitlement management.
func (m *EntitlementManagement) GetSettings()(EntitlementManagementSettingsable) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
func (m *EntitlementManagement) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EntitlementManagement) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetAccessPackageAssignmentApprovals() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageAssignmentApprovals()))
        for i, v := range m.GetAccessPackageAssignmentApprovals() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageAssignmentApprovals", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAccessPackages() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackages()))
        for i, v := range m.GetAccessPackages() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackages", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentPolicies() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignmentPolicies()))
        for i, v := range m.GetAssignmentPolicies() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignmentPolicies", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignmentRequests() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignmentRequests()))
        for i, v := range m.GetAssignmentRequests() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignmentRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetAssignments() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    if m.GetCatalogs() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCatalogs()))
        for i, v := range m.GetCatalogs() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("catalogs", cast)
        if err != nil {
            return err
        }
    }
    if m.GetConnectedOrganizations() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConnectedOrganizations()))
        for i, v := range m.GetConnectedOrganizations() {
            cast[i] = v.(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable)
        }
        err = writer.WriteCollectionOfObjectValues("connectedOrganizations", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("settings", m.GetSettings())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAccessPackageAssignmentApprovals sets the accessPackageAssignmentApprovals property value. Approval stages for assignment requests.
func (m *EntitlementManagement) SetAccessPackageAssignmentApprovals(value []Approvalable)() {
    if m != nil {
        m.accessPackageAssignmentApprovals = value
    }
}
// SetAccessPackages sets the accessPackages property value. Represents access package objects.
func (m *EntitlementManagement) SetAccessPackages(value []AccessPackageable)() {
    if m != nil {
        m.accessPackages = value
    }
}
// SetAssignmentPolicies sets the assignmentPolicies property value. 
func (m *EntitlementManagement) SetAssignmentPolicies(value []AccessPackageAssignmentPolicyable)() {
    if m != nil {
        m.assignmentPolicies = value
    }
}
// SetAssignmentRequests sets the assignmentRequests property value. Represents access package assignment requests created by or on behalf of a user.
func (m *EntitlementManagement) SetAssignmentRequests(value []AccessPackageAssignmentRequestable)() {
    if m != nil {
        m.assignmentRequests = value
    }
}
// SetAssignments sets the assignments property value. Represents the grant of an access package to a subject (user or group).
func (m *EntitlementManagement) SetAssignments(value []AccessPackageAssignmentable)() {
    if m != nil {
        m.assignments = value
    }
}
// SetCatalogs sets the catalogs property value. Represents a group of access packages.
func (m *EntitlementManagement) SetCatalogs(value []AccessPackageCatalogable)() {
    if m != nil {
        m.catalogs = value
    }
}
// SetConnectedOrganizations sets the connectedOrganizations property value. Represents references to a directory or domain of another organization whose users can request access.
func (m *EntitlementManagement) SetConnectedOrganizations(value []ConnectedOrganizationable)() {
    if m != nil {
        m.connectedOrganizations = value
    }
}
// SetSettings sets the settings property value. Represents the settings that control the behavior of Azure AD entitlement management.
func (m *EntitlementManagement) SetSettings(value EntitlementManagementSettingsable)() {
    if m != nil {
        m.settings = value
    }
}
