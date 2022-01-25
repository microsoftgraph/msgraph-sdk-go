package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EntitlementManagementSettings 
type EntitlementManagementSettings struct {
    Entity
    // If externalUserLifecycleAction is blockSignInAndDelete, the duration, typically a number of days, after an external user is blocked from sign in before their account is deleted.
    durationUntilExternalUserDeletedAfterBlocked *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration;
    // One of None, BlockSignIn, or BlockSignInAndDelete.
    externalUserLifecycleAction *AccessPackageExternalUserLifecycleAction;
}
// NewEntitlementManagementSettings instantiates a new entitlementManagementSettings and sets the default values.
func NewEntitlementManagementSettings()(*EntitlementManagementSettings) {
    m := &EntitlementManagementSettings{
        Entity: *NewEntity(),
    }
    return m
}
// GetDurationUntilExternalUserDeletedAfterBlocked gets the durationUntilExternalUserDeletedAfterBlocked property value. If externalUserLifecycleAction is blockSignInAndDelete, the duration, typically a number of days, after an external user is blocked from sign in before their account is deleted.
func (m *EntitlementManagementSettings) GetDurationUntilExternalUserDeletedAfterBlocked()(*i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration) {
    if m == nil {
        return nil
    } else {
        return m.durationUntilExternalUserDeletedAfterBlocked
    }
}
// GetExternalUserLifecycleAction gets the externalUserLifecycleAction property value. One of None, BlockSignIn, or BlockSignInAndDelete.
func (m *EntitlementManagementSettings) GetExternalUserLifecycleAction()(*AccessPackageExternalUserLifecycleAction) {
    if m == nil {
        return nil
    } else {
        return m.externalUserLifecycleAction
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EntitlementManagementSettings) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["durationUntilExternalUserDeletedAfterBlocked"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetISODurationValue()
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
// Serialize serializes information the current object
func (m *EntitlementManagementSettings) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteISODurationValue("durationUntilExternalUserDeletedAfterBlocked", m.GetDurationUntilExternalUserDeletedAfterBlocked())
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
// SetDurationUntilExternalUserDeletedAfterBlocked sets the durationUntilExternalUserDeletedAfterBlocked property value. If externalUserLifecycleAction is blockSignInAndDelete, the duration, typically a number of days, after an external user is blocked from sign in before their account is deleted.
func (m *EntitlementManagementSettings) SetDurationUntilExternalUserDeletedAfterBlocked(value *i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ISODuration)() {
    if m != nil {
        m.durationUntilExternalUserDeletedAfterBlocked = value
    }
}
// SetExternalUserLifecycleAction sets the externalUserLifecycleAction property value. One of None, BlockSignIn, or BlockSignInAndDelete.
func (m *EntitlementManagementSettings) SetExternalUserLifecycleAction(value *AccessPackageExternalUserLifecycleAction)() {
    if m != nil {
        m.externalUserLifecycleAction = value
    }
}
