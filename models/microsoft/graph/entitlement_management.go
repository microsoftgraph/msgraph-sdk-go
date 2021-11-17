package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EntitlementManagement struct {
    Entity
    // 
    accessPackageAssignmentApprovals []Approval;
    // 
    accessPackages []AccessPackage;
    // 
    assignmentRequests []AccessPackageAssignmentRequest;
    // 
    assignments []AccessPackageAssignment;
    // 
    catalogs []AccessPackageCatalog;
    // 
    connectedOrganizations []ConnectedOrganization;
    // 
    settings *EntitlementManagementSettings;
}
// Instantiates a new entitlementManagement and sets the default values.
func NewEntitlementManagement()(*EntitlementManagement) {
    m := &EntitlementManagement{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the accessPackageAssignmentApprovals property value. 
func (m *EntitlementManagement) GetAccessPackageAssignmentApprovals()([]Approval) {
    if m == nil {
        return nil
    } else {
        return m.accessPackageAssignmentApprovals
    }
}
// Gets the accessPackages property value. 
func (m *EntitlementManagement) GetAccessPackages()([]AccessPackage) {
    if m == nil {
        return nil
    } else {
        return m.accessPackages
    }
}
// Gets the assignmentRequests property value. 
func (m *EntitlementManagement) GetAssignmentRequests()([]AccessPackageAssignmentRequest) {
    if m == nil {
        return nil
    } else {
        return m.assignmentRequests
    }
}
// Gets the assignments property value. 
func (m *EntitlementManagement) GetAssignments()([]AccessPackageAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the catalogs property value. 
func (m *EntitlementManagement) GetCatalogs()([]AccessPackageCatalog) {
    if m == nil {
        return nil
    } else {
        return m.catalogs
    }
}
// Gets the connectedOrganizations property value. 
func (m *EntitlementManagement) GetConnectedOrganizations()([]ConnectedOrganization) {
    if m == nil {
        return nil
    } else {
        return m.connectedOrganizations
    }
}
// Gets the settings property value. 
func (m *EntitlementManagement) GetSettings()(*EntitlementManagementSettings) {
    if m == nil {
        return nil
    } else {
        return m.settings
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the accessPackageAssignmentApprovals property value. 
// Parameters:
//  - value : Value to set for the accessPackageAssignmentApprovals property.
func (m *EntitlementManagement) SetAccessPackageAssignmentApprovals(value []Approval)() {
    m.accessPackageAssignmentApprovals = value
}
// Sets the accessPackages property value. 
// Parameters:
//  - value : Value to set for the accessPackages property.
func (m *EntitlementManagement) SetAccessPackages(value []AccessPackage)() {
    m.accessPackages = value
}
// Sets the assignmentRequests property value. 
// Parameters:
//  - value : Value to set for the assignmentRequests property.
func (m *EntitlementManagement) SetAssignmentRequests(value []AccessPackageAssignmentRequest)() {
    m.assignmentRequests = value
}
// Sets the assignments property value. 
// Parameters:
//  - value : Value to set for the assignments property.
func (m *EntitlementManagement) SetAssignments(value []AccessPackageAssignment)() {
    m.assignments = value
}
// Sets the catalogs property value. 
// Parameters:
//  - value : Value to set for the catalogs property.
func (m *EntitlementManagement) SetCatalogs(value []AccessPackageCatalog)() {
    m.catalogs = value
}
// Sets the connectedOrganizations property value. 
// Parameters:
//  - value : Value to set for the connectedOrganizations property.
func (m *EntitlementManagement) SetConnectedOrganizations(value []ConnectedOrganization)() {
    m.connectedOrganizations = value
}
// Sets the settings property value. 
// Parameters:
//  - value : Value to set for the settings property.
func (m *EntitlementManagement) SetSettings(value *EntitlementManagementSettings)() {
    m.settings = value
}
