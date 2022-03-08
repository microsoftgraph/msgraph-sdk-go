package verifywindowsenrollmentautodiscoverywithdomainname

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponse provides operations to call the verifyWindowsEnrollmentAutoDiscovery method.
type VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    value *bool;
}
// NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponse instantiates a new verifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponse and sets the default values.
func NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponse()(*VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponse) {
    m := &VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponseFromDiscriminatorValue(parseNode i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, error) {
    return NewVerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponse) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponse) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["value"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetValue(val)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. 
func (m *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponse) GetValue()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.value
    }
}
func (m *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponse) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponse) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteBoolValue("value", m.GetValue())
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
func (m *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponse) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetValue sets the value property value. 
func (m *VerifyWindowsEnrollmentAutoDiscoveryWithDomainNameResponse) SetValue(value *bool)() {
    if m != nil {
        m.value = value
    }
}
