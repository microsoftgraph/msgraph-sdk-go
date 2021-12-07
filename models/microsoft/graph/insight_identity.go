package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InsightIdentity 
type InsightIdentity struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The email address of the user who shared the item.
    address *string;
    // The display name of the user who shared the item.
    displayName *string;
    // The id of the user who shared the item.
    id *string;
}
// NewInsightIdentity instantiates a new insightIdentity and sets the default values.
func NewInsightIdentity()(*InsightIdentity) {
    m := &InsightIdentity{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InsightIdentity) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAddress gets the address property value. The email address of the user who shared the item.
func (m *InsightIdentity) GetAddress()(*string) {
    if m == nil {
        return nil
    } else {
        return m.address
    }
}
// GetDisplayName gets the displayName property value. The display name of the user who shared the item.
func (m *InsightIdentity) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetId gets the id property value. The id of the user who shared the item.
func (m *InsightIdentity) GetId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.id
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InsightIdentity) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["address"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetAddress(val)
        }
        return nil
    }
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["id"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetId(val)
        }
        return nil
    }
    return res
}
func (m *InsightIdentity) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *InsightIdentity) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("address", m.GetAddress())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("id", m.GetId())
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
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *InsightIdentity) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAddress sets the address property value. The email address of the user who shared the item.
func (m *InsightIdentity) SetAddress(value *string)() {
    if m != nil {
        m.address = value
    }
}
// SetDisplayName sets the displayName property value. The display name of the user who shared the item.
func (m *InsightIdentity) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetId sets the id property value. The id of the user who shared the item.
func (m *InsightIdentity) SetId(value *string)() {
    if m != nil {
        m.id = value
    }
}
