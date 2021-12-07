package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EntitlementManagement 
type EntitlementManagement struct {
    Entity
    // 
    accessPackageAssignmentApprovals []Approval;
    // Access packages.
    accessPackages []AccessPackage;
    // Access package assignment requests.
    assignmentRequests []AccessPackageAssignmentRequest;
    // Access package assignments.
    assignments []AccessPackageAssignment;
    // Access package catalogs.
    catalogs []AccessPackageCatalog;
    // Connected organizations.
    connectedOrganizations []ConnectedOrganization;
    // Entitlement management settings.
    settings *EntitlementManagementSettings;
}
// NewEntitlementManagement instantiates a new entitlementManagement and sets the default values.
func NewEntitlementManagement()(*EntitlementManagement) {
    m := &EntitlementManagement{
        Entity: *NewEntity(),
    }
    return m
}
// GetAccessPackageAssignmentApprovals gets the accessPackageAssignmentApprovals property value. 
func (m *EntitlementManagement) GetAccessPackageAssignmentApprovals()([]Approval) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignmentApprovals
    }
}
// GetAccessPackages gets the accessPackages property value. Access packages.
func (m *EntitlementManagement) GetAccessPackages()([]AccessPackage) {
    if m == nil {
        return nil
    } else {
        return m.accessPackages
    }
}
// GetAssignmentRequests gets the assignmentRequests property value. Access package assignment requests.
func (m *EntitlementManagement) GetAssignmentRequests()([]AccessPackageAssignmentRequest) {
    if m == nil {
        return nil
    } else {
        return m.assignmentRequests
    }
}
// GetAssignments gets the assignments property value. Access package assignments.
func (m *EntitlementManagement) GetAssignments()([]AccessPackageAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// GetCatalogs gets the catalogs property value. Access package catalogs.
func (m *EntitlementManagement) GetCatalogs()([]AccessPackageCatalog) {
    if m == nil {
        return nil
    } else {
        return m.catalogs
    }
}
// GetConnectedOrganizations gets the connectedOrganizations property value. Connected organizations.
func (m *EntitlementManagement) GetConnectedOrganizations()([]ConnectedOrganization) {
    if m == nil {
        return nil
    } else {
        return m.connectedOrganizations
    }
}
// GetSettings gets the settings property value. Entitlement management settings.
func (m *EntitlementManagement) GetSettings()(*EntitlementManagementSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EntitlementManagement) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["accessPackageAssignmentApprovals"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewApproval() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Approval, len(val))
            for i, v := range val {
                res[i] = *(v.(*Approval))
            }
            m.SetAccessPackageAssignmentApprovals(res)
        }
        return nil
    }
    res["accessPackages"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackage() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackage, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackage))
            }
            m.SetAccessPackages(res)
        }
        return nil
    }
    res["assignmentRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageAssignmentRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignmentRequest, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackageAssignmentRequest))
            }
            m.SetAssignmentRequests(res)
        }
        return nil
    }
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageAssignment() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageAssignment, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackageAssignment))
            }
            m.SetAssignments(res)
        }
        return nil
    }
    res["catalogs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewAccessPackageCatalog() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]AccessPackageCatalog, len(val))
            for i, v := range val {
                res[i] = *(v.(*AccessPackageCatalog))
            }
            m.SetCatalogs(res)
        }
        return nil
    }
    res["connectedOrganizations"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewConnectedOrganization() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ConnectedOrganization, len(val))
            for i, v := range val {
                res[i] = *(v.(*ConnectedOrganization))
            }
            m.SetConnectedOrganizations(res)
        }
        return nil
    }
    res["settings"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEntitlementManagementSettings() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSettings(val.(*EntitlementManagementSettings))
        }
        return nil
    }
    return res
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
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackageAssignmentApprovals()))
        for i, v := range m.GetAccessPackageAssignmentApprovals() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackageAssignmentApprovals", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAccessPackages()))
        for i, v := range m.GetAccessPackages() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("accessPackages", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignmentRequests()))
        for i, v := range m.GetAssignmentRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignmentRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetAssignments()))
        for i, v := range m.GetAssignments() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("assignments", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetCatalogs()))
        for i, v := range m.GetCatalogs() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("catalogs", cast)
        if err != nil {
            return err
        }
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetConnectedOrganizations()))
        for i, v := range m.GetConnectedOrganizations() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
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
// SetAccessPackageAssignmentApprovals sets the accessPackageAssignmentApprovals property value. 
func (m *EntitlementManagement) SetAccessPackageAssignmentApprovals(value []Approval)() {
    if m != nil {
        m.accessPackageAssignmentApprovals = value
    }
}
// SetAccessPackages sets the accessPackages property value. Access packages.
func (m *EntitlementManagement) SetAccessPackages(value []AccessPackage)() {
    if m != nil {
        m.accessPackages = value
    }
}
// SetAssignmentRequests sets the assignmentRequests property value. Access package assignment requests.
func (m *EntitlementManagement) SetAssignmentRequests(value []AccessPackageAssignmentRequest)() {
    if m != nil {
        m.assignmentRequests = value
    }
}
// SetAssignments sets the assignments property value. Access package assignments.
func (m *EntitlementManagement) SetAssignments(value []AccessPackageAssignment)() {
    if m != nil {
        m.assignments = value
    }
}
// SetCatalogs sets the catalogs property value. Access package catalogs.
func (m *EntitlementManagement) SetCatalogs(value []AccessPackageCatalog)() {
    if m != nil {
        m.catalogs = value
    }
}
// SetConnectedOrganizations sets the connectedOrganizations property value. Connected organizations.
func (m *EntitlementManagement) SetConnectedOrganizations(value []ConnectedOrganization)() {
    if m != nil {
        m.connectedOrganizations = value
    }
}
// SetSettings sets the settings property value. Entitlement management settings.
func (m *EntitlementManagement) SetSettings(value *EntitlementManagementSettings)() {
    if m != nil {
        m.settings = value
    }
}
