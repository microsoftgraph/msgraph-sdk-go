package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EntitlementManagementSettings struct {
    Entity
    // 
    durationUntilExternalUserDeletedAfterBlocked *string;
    // One of None, BlockSignIn, or BlockSignInAndDelete.
    externalUserLifecycleAction *AccessPackageExternalUserLifecycleAction;
}
// Instantiates a new entitlementManagementSettings and sets the default values.
func NewEntitlementManagementSettings()(*EntitlementManagementSettings) {
    m := &EntitlementManagementSettings{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the durationUntilExternalUserDeletedAfterBlocked property value. 
func (m *EntitlementManagementSettings) GetDurationUntilExternalUserDeletedAfterBlocked()(*string) {
    if m == nil {
        return nil
    } else {
        return m.durationUntilExternalUserDeletedAfterBlocked
    }
}
// Gets the externalUserLifecycleAction property value. One of None, BlockSignIn, or BlockSignInAndDelete.
func (m *EntitlementManagementSettings) GetExternalUserLifecycleAction()(*AccessPackageExternalUserLifecycleAction) {
    if m == nil {
        return nil
    } else {
        return m.externalUserLifecycleAction
    }
}
// The deserialization information for the current model
func (m *EntitlementManagementSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["durationUntilExternalUserDeletedAfterBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDurationUntilExternalUserDeletedAfterBlocked(val)
        }
        return nil
    }
    res["externalUserLifecycleAction"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseAccessPackageExternalUserLifecycleAction)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(AccessPackageExternalUserLifecycleAction)
            m.SetExternalUserLifecycleAction(&cast)
        }
        return nil
    }
    return res
}
func (m *EntitlementManagementSettings) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *EntitlementManagementSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteStringValue("durationUntilExternalUserDeletedAfterBlocked", m.GetDurationUntilExternalUserDeletedAfterBlocked())
        if err != nil {
            return err
        }
    }
    if m.GetExternalUserLifecycleAction() != nil {
        cast := m.GetExternalUserLifecycleAction().String()
        err = writer.WriteStringValue("externalUserLifecycleAction", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// Sets the durationUntilExternalUserDeletedAfterBlocked property value. 
// Parameters:
//  - value : Value to set for the durationUntilExternalUserDeletedAfterBlocked property.
func (m *EntitlementManagementSettings) SetDurationUntilExternalUserDeletedAfterBlocked(value *string)() {
    m.durationUntilExternalUserDeletedAfterBlocked = value
}
// Sets the externalUserLifecycleAction property value. One of None, BlockSignIn, or BlockSignInAndDelete.
// Parameters:
//  - value : Value to set for the externalUserLifecycleAction property.
func (m *EntitlementManagementSettings) SetExternalUserLifecycleAction(value *AccessPackageExternalUserLifecycleAction)() {
    m.externalUserLifecycleAction = value
}
