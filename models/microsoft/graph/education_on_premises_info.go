package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationOnPremisesInfo 
type EducationOnPremisesInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Unique identifier for the user object in Active Directory.
    immutableId *string;
}
// NewEducationOnPremisesInfo instantiates a new educationOnPremisesInfo and sets the default values.
func NewEducationOnPremisesInfo()(*EducationOnPremisesInfo) {
    m := &EducationOnPremisesInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationOnPremisesInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetImmutableId gets the immutableId property value. Unique identifier for the user object in Active Directory.
func (m *EducationOnPremisesInfo) GetImmutableId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.immutableId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationOnPremisesInfo) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["immutableId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetImmutableId(val)
        }
        return nil
    }
    return res
}
func (m *EducationOnPremisesInfo) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *EducationOnPremisesInfo) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("immutableId", m.GetImmutableId())
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
func (m *EducationOnPremisesInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetImmutableId sets the immutableId property value. Unique identifier for the user object in Active Directory.
func (m *EducationOnPremisesInfo) SetImmutableId(value *string)() {
    if m != nil {
        m.immutableId = value
    }
}
