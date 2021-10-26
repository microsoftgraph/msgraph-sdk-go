package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type TargetedManagedAppConfiguration struct {
    ManagedAppConfiguration
    // List of apps to which the policy is deployed.
    apps []ManagedMobileApp;
    // Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
    assignments []TargetedManagedAppPolicyAssignment;
    // Count of apps to which the current policy is deployed.
    deployedAppCount *int32;
    // Navigation property to deployment summary of the configuration.
    deploymentSummary *ManagedAppPolicyDeploymentSummary;
    // Indicates if the policy is deployed to any inclusion groups or not.
    isAssigned *bool;
}
// Instantiates a new targetedManagedAppConfiguration and sets the default values.
func NewTargetedManagedAppConfiguration()(*TargetedManagedAppConfiguration) {
    m := &TargetedManagedAppConfiguration{
        ManagedAppConfiguration: *NewManagedAppConfiguration(),
    }
    return m
}
// Gets the apps property value. List of apps to which the policy is deployed.
func (m *TargetedManagedAppConfiguration) GetApps()([]ManagedMobileApp) {
    if m == nil {
        return nil
    } else {
        return m.apps
    }
}
// Gets the assignments property value. Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
func (m *TargetedManagedAppConfiguration) GetAssignments()([]TargetedManagedAppPolicyAssignment) {
    if m == nil {
        return nil
    } else {
        return m.assignments
    }
}
// Gets the deployedAppCount property value. Count of apps to which the current policy is deployed.
func (m *TargetedManagedAppConfiguration) GetDeployedAppCount()(*int32) {
    if m == nil {
        return nil
    } else {
        return m.deployedAppCount
    }
}
// Gets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
func (m *TargetedManagedAppConfiguration) GetDeploymentSummary()(*ManagedAppPolicyDeploymentSummary) {
    if m == nil {
        return nil
    } else {
        return m.deploymentSummary
    }
}
// Gets the isAssigned property value. Indicates if the policy is deployed to any inclusion groups or not.
func (m *TargetedManagedAppConfiguration) GetIsAssigned()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isAssigned
    }
}
// The deserialization information for the current model
func (m *TargetedManagedAppConfiguration) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ManagedAppConfiguration.GetFieldDeserializers()
    res["apps"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedMobileApp() })
        if err != nil {
            return err
        }
        res := make([]ManagedMobileApp, len(val))
        for i, v := range val {
            res[i] = *(v.(*ManagedMobileApp))
        }
        m.SetApps(res)
        return nil
    }
    res["assignments"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTargetedManagedAppPolicyAssignment() })
        if err != nil {
            return err
        }
        res := make([]TargetedManagedAppPolicyAssignment, len(val))
        for i, v := range val {
            res[i] = *(v.(*TargetedManagedAppPolicyAssignment))
        }
        m.SetAssignments(res)
        return nil
    }
    res["deployedAppCount"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetInt32Value()
        if err != nil {
            return err
        }
        m.SetDeployedAppCount(val)
        return nil
    }
    res["deploymentSummary"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewManagedAppPolicyDeploymentSummary() })
        if err != nil {
            return err
        }
        m.SetDeploymentSummary(val.(*ManagedAppPolicyDeploymentSummary))
        return nil
    }
    res["isAssigned"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsAssigned(val)
        return nil
    }
    return res
}
func (m *TargetedManagedAppConfiguration) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *TargetedManagedAppConfiguration) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ManagedAppConfiguration.Serialize(writer)
    if err != nil {
        return err
    }
    {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetApps()))
        for i, v := range m.GetApps() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("apps", cast)
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
        err = writer.WriteInt32Value("deployedAppCount", m.GetDeployedAppCount())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("deploymentSummary", m.GetDeploymentSummary())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("isAssigned", m.GetIsAssigned())
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the apps property value. List of apps to which the policy is deployed.
// Parameters:
//  - value : Value to set for the apps property.
func (m *TargetedManagedAppConfiguration) SetApps(value []ManagedMobileApp)() {
    m.apps = value
}
// Sets the assignments property value. Navigation property to list of inclusion and exclusion groups to which the policy is deployed.
// Parameters:
//  - value : Value to set for the assignments property.
func (m *TargetedManagedAppConfiguration) SetAssignments(value []TargetedManagedAppPolicyAssignment)() {
    m.assignments = value
}
// Sets the deployedAppCount property value. Count of apps to which the current policy is deployed.
// Parameters:
//  - value : Value to set for the deployedAppCount property.
func (m *TargetedManagedAppConfiguration) SetDeployedAppCount(value *int32)() {
    m.deployedAppCount = value
}
// Sets the deploymentSummary property value. Navigation property to deployment summary of the configuration.
// Parameters:
//  - value : Value to set for the deploymentSummary property.
func (m *TargetedManagedAppConfiguration) SetDeploymentSummary(value *ManagedAppPolicyDeploymentSummary)() {
    m.deploymentSummary = value
}
// Sets the isAssigned property value. Indicates if the policy is deployed to any inclusion groups or not.
// Parameters:
//  - value : Value to set for the isAssigned property.
func (m *TargetedManagedAppConfiguration) SetIsAssigned(value *bool)() {
    m.isAssigned = value
}
