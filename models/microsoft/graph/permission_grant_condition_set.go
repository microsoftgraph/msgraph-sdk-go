package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PermissionGrantConditionSet struct {
    Entity
    clientApplicationIds []string;
    clientApplicationPublisherIds []string;
    clientApplicationsFromVerifiedPublisherOnly *bool;
    clientApplicationTenantIds []string;
    permissionClassification *string;
    permissions []string;
    permissionType *PermissionType;
    resourceApplication *string;
}
func NewPermissionGrantConditionSet()(*PermissionGrantConditionSet) {
    m := &PermissionGrantConditionSet{
        Entity: *NewEntity(),
    }
    return m
}
func (m *PermissionGrantConditionSet) GetClientApplicationIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.clientApplicationIds
    }
}
func (m *PermissionGrantConditionSet) GetClientApplicationPublisherIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.clientApplicationPublisherIds
    }
}
func (m *PermissionGrantConditionSet) GetClientApplicationsFromVerifiedPublisherOnly()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.clientApplicationsFromVerifiedPublisherOnly
    }
}
func (m *PermissionGrantConditionSet) GetClientApplicationTenantIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.clientApplicationTenantIds
    }
}
func (m *PermissionGrantConditionSet) GetPermissionClassification()(*string) {
    if m == nil {
        return nil
    } else {
        return m.permissionClassification
    }
}
func (m *PermissionGrantConditionSet) GetPermissions()([]string) {
    if m == nil {
        return nil
    } else {
        return m.permissions
    }
}
func (m *PermissionGrantConditionSet) GetPermissionType()(*PermissionType) {
    if m == nil {
        return nil
    } else {
        return m.permissionType
    }
}
func (m *PermissionGrantConditionSet) GetResourceApplication()(*string) {
    if m == nil {
        return nil
    } else {
        return m.resourceApplication
    }
}
func (m *PermissionGrantConditionSet) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["clientApplicationIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetClientApplicationIds(res)
        return nil
    }
    res["clientApplicationPublisherIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetClientApplicationPublisherIds(res)
        return nil
    }
    res["clientApplicationsFromVerifiedPublisherOnly"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetClientApplicationsFromVerifiedPublisherOnly(val)
        return nil
    }
    res["clientApplicationTenantIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetClientApplicationTenantIds(res)
        return nil
    }
    res["permissionClassification"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetPermissionClassification(val)
        return nil
    }
    res["permissions"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetPermissions(res)
        return nil
    }
    res["permissionType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParsePermissionType)
        if err != nil {
            return err
        }
        cast := val.(PermissionType)
        m.SetPermissionType(&cast)
        return nil
    }
    res["resourceApplication"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetResourceApplication(val)
        return nil
    }
    return res
}
func (m *PermissionGrantConditionSet) IsNil()(bool) {
    return m == nil
}
func (m *PermissionGrantConditionSet) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteCollectionOfStringValues("clientApplicationIds", m.GetClientApplicationIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("clientApplicationPublisherIds", m.GetClientApplicationPublisherIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("clientApplicationsFromVerifiedPublisherOnly", m.GetClientApplicationsFromVerifiedPublisherOnly())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("clientApplicationTenantIds", m.GetClientApplicationTenantIds())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("permissionClassification", m.GetPermissionClassification())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteCollectionOfStringValues("permissions", m.GetPermissions())
        if err != nil {
            return err
        }
    }
    if m.GetPermissionType() != nil {
        cast := m.GetPermissionType().String()
        err = writer.WriteStringValue("permissionType", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("resourceApplication", m.GetResourceApplication())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *PermissionGrantConditionSet) SetClientApplicationIds(value []string)() {
    m.clientApplicationIds = value
}
func (m *PermissionGrantConditionSet) SetClientApplicationPublisherIds(value []string)() {
    m.clientApplicationPublisherIds = value
}
func (m *PermissionGrantConditionSet) SetClientApplicationsFromVerifiedPublisherOnly(value *bool)() {
    m.clientApplicationsFromVerifiedPublisherOnly = value
}
func (m *PermissionGrantConditionSet) SetClientApplicationTenantIds(value []string)() {
    m.clientApplicationTenantIds = value
}
func (m *PermissionGrantConditionSet) SetPermissionClassification(value *string)() {
    m.permissionClassification = value
}
func (m *PermissionGrantConditionSet) SetPermissions(value []string)() {
    m.permissions = value
}
func (m *PermissionGrantConditionSet) SetPermissionType(value *PermissionType)() {
    m.permissionType = value
}
func (m *PermissionGrantConditionSet) SetResourceApplication(value *string)() {
    m.resourceApplication = value
}
