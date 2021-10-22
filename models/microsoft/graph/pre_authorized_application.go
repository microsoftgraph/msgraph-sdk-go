package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

type PreAuthorizedApplication struct {
    additionalData map[string]interface{};
    appId *string;
    delegatedPermissionIds []string;
}
func NewPreAuthorizedApplication()(*PreAuthorizedApplication) {
    m := &PreAuthorizedApplication{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
func (m *PreAuthorizedApplication) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
func (m *PreAuthorizedApplication) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
func (m *PreAuthorizedApplication) GetDelegatedPermissionIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.delegatedPermissionIds
    }
}
func (m *PreAuthorizedApplication) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["appId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetAppId(val)
        return nil
    }
    res["delegatedPermissionIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        res := make([]string, len(val))
        for i, v := range val {
            res[i] = v.(string)
        }
        m.SetDelegatedPermissionIds(res)
        return nil
    }
    return res
}
func (m *PreAuthorizedApplication) IsNil()(bool) {
    return m == nil
}
func (m *PreAuthorizedApplication) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("appId", m.GetAppId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteCollectionOfStringValues("delegatedPermissionIds", m.GetDelegatedPermissionIds())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
func (m *PreAuthorizedApplication) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
func (m *PreAuthorizedApplication) SetAppId(value *string)() {
    m.appId = value
}
func (m *PreAuthorizedApplication) SetDelegatedPermissionIds(value []string)() {
    m.delegatedPermissionIds = value
}
