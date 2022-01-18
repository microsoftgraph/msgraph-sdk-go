package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// CloudAppSecuritySessionControl 
type CloudAppSecuritySessionControl struct {
    ConditionalAccessSessionControl
    // Possible values are: mcasConfigured, monitorOnly, blockDownloads. To learn more about these values, Deploy Conditional Access App Control for featured apps.
    cloudAppSecurityType *CloudAppSecuritySessionControlType;
}
// NewCloudAppSecuritySessionControl instantiates a new cloudAppSecuritySessionControl and sets the default values.
func NewCloudAppSecuritySessionControl()(*CloudAppSecuritySessionControl) {
    m := &CloudAppSecuritySessionControl{
        ConditionalAccessSessionControl: *NewConditionalAccessSessionControl(),
    }
    return m
}
// GetCloudAppSecurityType gets the cloudAppSecurityType property value. Possible values are: mcasConfigured, monitorOnly, blockDownloads. To learn more about these values, Deploy Conditional Access App Control for featured apps.
func (m *CloudAppSecuritySessionControl) GetCloudAppSecurityType()(*CloudAppSecuritySessionControlType) {
    if m == nil {
        return nil
    } else {
        return m.cloudAppSecurityType
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CloudAppSecuritySessionControl) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.ConditionalAccessSessionControl.GetFieldDeserializers()
    res["cloudAppSecurityType"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseCloudAppSecuritySessionControlType)
        if err != nil {
            return err
        }
        if val != nil {
            cast := val.(CloudAppSecuritySessionControlType)
            m.SetCloudAppSecurityType(&cast)
        }
        return nil
    }
    return res
}
func (m *CloudAppSecuritySessionControl) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CloudAppSecuritySessionControl) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.ConditionalAccessSessionControl.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetCloudAppSecurityType() != nil {
        cast := m.GetCloudAppSecurityType().String()
        err = writer.WriteStringValue("cloudAppSecurityType", &cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetCloudAppSecurityType sets the cloudAppSecurityType property value. Possible values are: mcasConfigured, monitorOnly, blockDownloads. To learn more about these values, Deploy Conditional Access App Control for featured apps.
func (m *CloudAppSecuritySessionControl) SetCloudAppSecurityType(value *CloudAppSecuritySessionControlType)() {
    if m != nil {
        m.cloudAppSecurityType = value
    }
}
