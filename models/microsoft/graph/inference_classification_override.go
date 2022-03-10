package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// InferenceClassificationOverride provides operations to manage the collection of drive entities.
type InferenceClassificationOverride struct {
    Entity
    // Specifies how incoming messages from a specific sender should always be classified as. The possible values are: focused, other.
    classifyAs *InferenceClassificationType;
    // The email address information of the sender for whom the override is created.
    senderEmailAddress EmailAddressable;
}
// NewInferenceClassificationOverride instantiates a new inferenceClassificationOverride and sets the default values.
func NewInferenceClassificationOverride()(*InferenceClassificationOverride) {
    m := &InferenceClassificationOverride{
        Entity: *NewEntity(),
    }
    return m
}
// CreateInferenceClassificationOverrideFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateInferenceClassificationOverrideFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewInferenceClassificationOverride(), nil
}
// GetClassifyAs gets the classifyAs property value. Specifies how incoming messages from a specific sender should always be classified as. The possible values are: focused, other.
func (m *InferenceClassificationOverride) GetClassifyAs()(*InferenceClassificationType) {
    if m == nil {
        return nil
    } else {
        return m.classifyAs
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
            m.SetClassifyAs(val.(*InferenceClassificationType))
        }
        return nil
    }
    res["senderEmailAddress"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(CreateEmailAddressFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSenderEmailAddress(val.(EmailAddressable))
        }
        return nil
    }
    return res
}
// GetSenderEmailAddress gets the senderEmailAddress property value. The email address information of the sender for whom the override is created.
func (m *InferenceClassificationOverride) GetSenderEmailAddress()(EmailAddressable) {
    if m == nil {
        return nil
    } else {
        return m.senderEmailAddress
    }
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
        cast := (*m.GetClassifyAs()).String()
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
func (m *InferenceClassificationOverride) SetSenderEmailAddress(value EmailAddressable)() {
    if m != nil {
        m.senderEmailAddress = value
    }
}
