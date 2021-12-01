package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// PreAuthorizedApplication 
type PreAuthorizedApplication struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The unique identifier for the application.
    appId *string;
    // 
    delegatedPermissionIds []string;
}
// NewPreAuthorizedApplication instantiates a new preAuthorizedApplication and sets the default values.
func NewPreAuthorizedApplication()(*PreAuthorizedApplication) {
    m := &PreAuthorizedApplication{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PreAuthorizedApplication) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAppId gets the appId property value. The unique identifier for the application.
func (m *PreAuthorizedApplication) GetAppId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.appId
    }
}
// GetDelegatedPermissionIds gets the delegatedPermissionIds property value. 
func (m *PreAuthorizedApplication) GetDelegatedPermissionIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.delegatedPermissionIds
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *PreAuthorizedApplication) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["appId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAppId(val)
        }
        return nil
    }
    res["delegatedPermissionIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetDelegatedPermissionIds(res)
        }
        return nil
    }
    return res
}
func (m *PreAuthorizedApplication) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *PreAuthorizedApplication) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAppId sets the appId property value. The unique identifier for the application.
func (m *PreAuthorizedApplication) SetAppId(value *string)() {
    if m != nil {
        m.appId = value
    }
}
// SetDelegatedPermissionIds sets the delegatedPermissionIds property value. 
func (m *PreAuthorizedApplication) SetDelegatedPermissionIds(value []string)() {
    if m != nil {
        m.delegatedPermissionIds = value
    }
}
