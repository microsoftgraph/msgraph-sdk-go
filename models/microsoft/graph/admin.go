package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Admin 
type Admin struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // A container for service communications resources. Read-only.
    serviceAnnouncement *ServiceAnnouncement;
}
// NewAdmin instantiates a new Admin and sets the default values.
func NewAdmin()(*Admin) {
    m := &Admin{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *Admin) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetServiceAnnouncement gets the serviceAnnouncement property value. A container for service communications resources. Read-only.
func (m *Admin) GetServiceAnnouncement()(*ServiceAnnouncement) {
    if m == nil {
        return nil
    } else {
        return m.serviceAnnouncement
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Admin) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["serviceAnnouncement"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewServiceAnnouncement() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetServiceAnnouncement(val.(*ServiceAnnouncement))
        }
        return nil
    }
    return res
}
func (m *Admin) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Admin) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("serviceAnnouncement", m.GetServiceAnnouncement())
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
func (m *Admin) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetServiceAnnouncement sets the serviceAnnouncement property value. A container for service communications resources. Read-only.
func (m *Admin) SetServiceAnnouncement(value *ServiceAnnouncement)() {
    if m != nil {
        m.serviceAnnouncement = value
    }
}
