package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type InferenceClassificationOverride struct {
    Entity
    // Specifies how incoming messages from a specific sender should always be classified as. The possible values are: focused, other.
    classifyAs *InferenceClassificationType;
    // The email address information of the sender for whom the override is created.
    senderEmailAddress *EmailAddress;
}
// Instantiates a new inferenceClassificationOverride and sets the default values.
func NewInferenceClassificationOverride()(*InferenceClassificationOverride) {
    m := &InferenceClassificationOverride{
        Entity: *NewEntity(),
    }
    return m
}
// Gets the classifyAs property value. Specifies how incoming messages from a specific sender should always be classified as. The possible values are: focused, other.
func (m *InferenceClassificationOverride) GetClassifyAs()(*InferenceClassificationType) {
    if m == nil {
        return nil
    } else {
        return m.classifyAs
    }
}
// Gets the senderEmailAddress property value. The email address information of the sender for whom the override is created.
func (m *InferenceClassificationOverride) GetSenderEmailAddress()(*EmailAddress) {
    if m == nil {
        return nil
    } else {
        return m.senderEmailAddress
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the classifyAs property value. Specifies how incoming messages from a specific sender should always be classified as. The possible values are: focused, other.
// Parameters:
//  - value : Value to set for the classifyAs property.
func (m *InferenceClassificationOverride) SetClassifyAs(value *InferenceClassificationType)() {
    m.classifyAs = value
}
// Sets the senderEmailAddress property value. The email address information of the sender for whom the override is created.
// Parameters:
//  - value : Value to set for the senderEmailAddress property.
func (m *InferenceClassificationOverride) SetSenderEmailAddress(value *EmailAddress)() {
    m.senderEmailAddress = value
}
