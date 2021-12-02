package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InferenceClassificationOverride 
type InferenceClassificationOverride struct {
    Entity
    // Specifies how incoming messages from a specific sender should always be classified as. The possible values are: focused, other.
    classifyAs *InferenceClassificationType;
    // The email address information of the sender for whom the override is created.
    senderEmailAddress *EmailAddress;
}
// NewInferenceClassificationOverride instantiates a new inferenceClassificationOverride and sets the default values.
func NewInferenceClassificationOverride()(*InferenceClassificationOverride) {
    m := &InferenceClassificationOverride{
        Entity: *NewEntity(),
    }
    return m
}
// GetClassifyAs gets the classifyAs property value. Specifies how incoming messages from a specific sender should always be classified as. The possible values are: focused, other.
func (m *InferenceClassificationOverride) GetClassifyAs()(*InferenceClassificationType) {
    if m == nil {
        return nil
    } else {
        return m.classifyAs
    }
}
// GetSenderEmailAddress gets the senderEmailAddress property value. The email address information of the sender for whom the override is created.
func (m *InferenceClassificationOverride) GetSenderEmailAddress()(*EmailAddress) {
    if m == nil {
        return nil
    } else {
        return m.senderEmailAddress
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *InferenceClassificationOverride) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["classifyAs"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseInferenceClassificationType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(InferenceClassificationType)
            m.SetClassifyAs(&cast)
        }
        return nil
    }
    res["senderEmailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewEmailAddress() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSenderEmailAddress(val.(*EmailAddress))
        }
        return nil
    }
    return res
}
func (m *InferenceClassificationOverride) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *InferenceClassificationOverride) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetClassifyAs() != nil {
        cast := m.GetClassifyAs().String()
        err = writer.WriteStringValue("classifyAs", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteObjectValue("senderEmailAddress", m.GetSenderEmailAddress())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetClassifyAs sets the classifyAs property value. Specifies how incoming messages from a specific sender should always be classified as. The possible values are: focused, other.
func (m *InferenceClassificationOverride) SetClassifyAs(value *InferenceClassificationType)() {
    if m != nil {
        m.classifyAs = value
    }
}
// SetSenderEmailAddress sets the senderEmailAddress property value. The email address information of the sender for whom the override is created.
func (m *InferenceClassificationOverride) SetSenderEmailAddress(value *EmailAddress)() {
    if m != nil {
        m.senderEmailAddress = value
    }
}
