package create

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// CreateRequestBody 
type CreateRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // 
    certificateSigningRequest *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintCertificateSigningRequest;
    // 
    connectorId *string;
    // 
    displayName *string;
    // 
    hasPhysicalDevice *bool;
    // 
    manufacturer *string;
    // 
    model *string;
    // 
    physicalDeviceId *string;
}
// NewCreateRequestBody instantiates a new createRequestBody and sets the default values.
func NewCreateRequestBody()(*CreateRequestBody) {
    m := &CreateRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreateRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCertificateSigningRequest gets the certificateSigningRequest property value. 
func (m *CreateRequestBody) GetCertificateSigningRequest()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintCertificateSigningRequest) {
    if m == nil {
        return nil
    } else {
        return m.certificateSigningRequest
    }
}
// GetConnectorId gets the connectorId property value. 
func (m *CreateRequestBody) GetConnectorId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.connectorId
    }
}
// GetDisplayName gets the displayName property value. 
func (m *CreateRequestBody) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetHasPhysicalDevice gets the hasPhysicalDevice property value. 
func (m *CreateRequestBody) GetHasPhysicalDevice()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasPhysicalDevice
    }
}
// GetManufacturer gets the manufacturer property value. 
func (m *CreateRequestBody) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// GetModel gets the model property value. 
func (m *CreateRequestBody) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// GetPhysicalDeviceId gets the physicalDeviceId property value. 
func (m *CreateRequestBody) GetPhysicalDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.physicalDeviceId
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreateRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    res["certificateSigningRequest"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetObjectValue(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.NewPrintCertificateSigningRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateSigningRequest(val.(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintCertificateSigningRequest))
        }
        return nil
    }
    res["connectorId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorId(val)
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
    res["hasPhysicalDevice"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasPhysicalDevice(val)
        }
        return nil
    }
    res["manufacturer"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["physicalDeviceId"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetPhysicalDeviceId(val)
        }
        return nil
    }
    return res
}
func (m *CreateRequestBody) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *CreateRequestBody) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("certificateSigningRequest", m.GetCertificateSigningRequest())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("connectorId", m.GetConnectorId())
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
        err := writer.WriteBoolValue("hasPhysicalDevice", m.GetHasPhysicalDevice())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("manufacturer", m.GetManufacturer())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("model", m.GetModel())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("physicalDeviceId", m.GetPhysicalDeviceId())
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
func (m *CreateRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetCertificateSigningRequest sets the certificateSigningRequest property value. 
func (m *CreateRequestBody) SetCertificateSigningRequest(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintCertificateSigningRequest)() {
    if m != nil {
        m.certificateSigningRequest = value
    }
}
// SetConnectorId sets the connectorId property value. 
func (m *CreateRequestBody) SetConnectorId(value *string)() {
    if m != nil {
        m.connectorId = value
    }
}
// SetDisplayName sets the displayName property value. 
func (m *CreateRequestBody) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetHasPhysicalDevice sets the hasPhysicalDevice property value. 
func (m *CreateRequestBody) SetHasPhysicalDevice(value *bool)() {
    if m != nil {
        m.hasPhysicalDevice = value
    }
}
// SetManufacturer sets the manufacturer property value. 
func (m *CreateRequestBody) SetManufacturer(value *string)() {
    if m != nil {
        m.manufacturer = value
    }
}
// SetModel sets the model property value. 
func (m *CreateRequestBody) SetModel(value *string)() {
    if m != nil {
        m.model = value
    }
}
// SetPhysicalDeviceId sets the physicalDeviceId property value. 
func (m *CreateRequestBody) SetPhysicalDeviceId(value *string)() {
    if m != nil {
        m.physicalDeviceId = value
    }
}
