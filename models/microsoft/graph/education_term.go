package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// EducationTerm 
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
// NewEducationTerm instantiates a new educationTerm and sets the default values.
func NewEducationTerm()(*EducationTerm) {
    m := &EducationTerm{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationTerm) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetDisplayName gets the displayName property value. Display name of the term.
func (m *EducationTerm) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetEndDate gets the endDate property value. End of the term.
func (m *EducationTerm) GetEndDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.endDate
    }
}
// GetExternalId gets the externalId property value. ID of term in the syncing system.
func (m *EducationTerm) GetExternalId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.externalId
    }
}
// GetStartDate gets the startDate property value. Start of the term.
func (m *EducationTerm) GetStartDate()(*string) {
    if m == nil {
        return nil
    } else {
        return m.startDate
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *EducationTerm) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
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
    res["endDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEndDate(val)
        }
        return nil
    }
    res["externalId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetExternalId(val)
        }
        return nil
    }
    res["startDate"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStartDate(val)
        }
        return nil
    }
    return res
}
func (m *EducationTerm) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
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
// SetAdditionalData sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *EducationTerm) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetDisplayName sets the displayName property value. Display name of the term.
func (m *EducationTerm) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetEndDate sets the endDate property value. End of the term.
func (m *EducationTerm) SetEndDate(value *string)() {
    if m != nil {
        m.endDate = value
    }
}
// SetExternalId sets the externalId property value. ID of term in the syncing system.
func (m *EducationTerm) SetExternalId(value *string)() {
    if m != nil {
        m.externalId = value
    }
}
// SetStartDate sets the startDate property value. Start of the term.
func (m *EducationTerm) SetStartDate(value *string)() {
    if m != nil {
        m.startDate = value
    }
}
