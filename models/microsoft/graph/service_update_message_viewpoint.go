package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type ServiceUpdateMessageViewpoint struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Indicates whether the user archived the message.
    isArchived *bool;
    // Indicates whether the user marked the message as favorite.
    isFavorited *bool;
    // Indicates whether the user read the message.
    isRead *bool;
}
// Instantiates a new serviceUpdateMessageViewpoint and sets the default values.
func NewServiceUpdateMessageViewpoint()(*ServiceUpdateMessageViewpoint) {
    m := &ServiceUpdateMessageViewpoint{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ServiceUpdateMessageViewpoint) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the isArchived property value. Indicates whether the user archived the message.
func (m *ServiceUpdateMessageViewpoint) GetIsArchived()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isArchived
    }
}
// Gets the isFavorited property value. Indicates whether the user marked the message as favorite.
func (m *ServiceUpdateMessageViewpoint) GetIsFavorited()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isFavorited
    }
}
// Gets the isRead property value. Indicates whether the user read the message.
func (m *ServiceUpdateMessageViewpoint) GetIsRead()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.isRead
    }
}
// The deserialization information for the current model
func (m *ServiceUpdateMessageViewpoint) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["isArchived"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsArchived(val)
        return nil
    }
    res["isFavorited"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsFavorited(val)
        return nil
    }
    res["isRead"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        m.SetIsRead(val)
        return nil
    }
    return res
}
func (m *ServiceUpdateMessageViewpoint) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *ServiceUpdateMessageViewpoint) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("isArchived", m.GetIsArchived())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isFavorited", m.GetIsFavorited())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteBoolValue("isRead", m.GetIsRead())
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *ServiceUpdateMessageViewpoint) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the isArchived property value. Indicates whether the user archived the message.
// Parameters:
//  - value : Value to set for the isArchived property.
func (m *ServiceUpdateMessageViewpoint) SetIsArchived(value *bool)() {
    m.isArchived = value
}
// Sets the isFavorited property value. Indicates whether the user marked the message as favorite.
// Parameters:
//  - value : Value to set for the isFavorited property.
func (m *ServiceUpdateMessageViewpoint) SetIsFavorited(value *bool)() {
    m.isFavorited = value
}
// Sets the isRead property value. Indicates whether the user read the message.
// Parameters:
//  - value : Value to set for the isRead property.
func (m *ServiceUpdateMessageViewpoint) SetIsRead(value *bool)() {
    m.isRead = value
}
