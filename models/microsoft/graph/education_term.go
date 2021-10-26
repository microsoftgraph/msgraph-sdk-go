package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// 
type EducationTerm struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // Display name of the term.
    displayName *string;
    // End of the term.
    endDate *string;
    // ID of term in the syncing system.
    externalId *string;
    // Start of the term.
    startDate *string;
}
// Instantiates a new educationTerm and sets the default values.
func NewEducationTerm()(*EducationTerm) {
    m := &EducationTerm{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationTerm) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the displayName property value. Display name of the term.
func (m *EducationTerm) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the endDate property value. End of the term.
func (m *EducationTerm) GetEndDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.endDate
    }
}
// Gets the externalId property value. ID of term in the syncing system.
func (m *EducationTerm) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// Gets the startDate property value. Start of the term.
func (m *EducationTerm) GetStartDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.startDate
    }
}
// The deserialization information for the current model
func (m *EducationTerm) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["displayName"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetDisplayName(val)
        return nil
    }
    res["endDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetEndDate(val)
        return nil
    }
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetExternalId(val)
        return nil
    }
    res["startDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        m.SetStartDate(val)
        return nil
    }
    return res
}
func (m *EducationTerm) IsNil()(bool) {
    return m == nil
}
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
func (m *EducationTerm) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("displayName", m.GetDisplayName())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("endDate", m.GetEndDate())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("externalId", m.GetExternalId())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("startDate", m.GetStartDate())
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
func (m *EducationTerm) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the displayName property value. Display name of the term.
// Parameters:
//  - value : Value to set for the displayName property.
func (m *EducationTerm) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the endDate property value. End of the term.
// Parameters:
//  - value : Value to set for the endDate property.
func (m *EducationTerm) SetEndDate(value *string)() {
    m.endDate = value
}
// Sets the externalId property value. ID of term in the syncing system.
// Parameters:
//  - value : Value to set for the externalId property.
func (m *EducationTerm) SetExternalId(value *string)() {
    m.externalId = value
}
// Sets the startDate property value. Start of the term.
// Parameters:
//  - value : Value to set for the startDate property.
func (m *EducationTerm) SetStartDate(value *string)() {
    m.startDate = value
}
