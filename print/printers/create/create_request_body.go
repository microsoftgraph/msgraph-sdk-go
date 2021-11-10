package create

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph"
)

// 
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
// Instantiates a new createRequestBody and sets the default values.
func NewCreateRequestBody()(*CreateRequestBody) {
    m := &CreateRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// Gets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreateRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// Gets the certificateSigningRequest property value. 
func (m *CreateRequestBody) GetCertificateSigningRequest()(*i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintCertificateSigningRequest) {
    if m == nil {
        return nil
    } else {
        return m.certificateSigningRequest
    }
}
// Gets the connectorId property value. 
func (m *CreateRequestBody) GetConnectorId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.connectorId
    }
}
// Gets the displayName property value. 
func (m *CreateRequestBody) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// Gets the hasPhysicalDevice property value. 
func (m *CreateRequestBody) GetHasPhysicalDevice()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasPhysicalDevice
    }
}
// Gets the manufacturer property value. 
func (m *CreateRequestBody) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// Gets the model property value. 
func (m *CreateRequestBody) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// Gets the physicalDeviceId property value. 
func (m *CreateRequestBody) GetPhysicalDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.physicalDeviceId
    }
}
// The deserialization information for the current model
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
// Serializes information the current object
// Parameters:
//  - writer : Serialization writer to use to serialize this model
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
// Sets the AdditionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
// Parameters:
//  - value : Value to set for the AdditionalData property.
func (m *CreateRequestBody) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// Sets the certificateSigningRequest property value. 
// Parameters:
//  - value : Value to set for the certificateSigningRequest property.
func (m *CreateRequestBody) SetCertificateSigningRequest(value *i4a838ef194e4c99e9f2c63ba10dab9cb120a89367c1d4ab0daa63bb424e20d87.PrintCertificateSigningRequest)() {
    m.certificateSigningRequest = value
}
// Sets the connectorId property value. 
// Parameters:
//  - value : Value to set for the connectorId property.
func (m *CreateRequestBody) SetConnectorId(value *string)() {
    m.connectorId = value
}
// Sets the displayName property value. 
// Parameters:
//  - value : Value to set for the displayName property.
func (m *CreateRequestBody) SetDisplayName(value *string)() {
    m.displayName = value
}
// Sets the hasPhysicalDevice property value. 
// Parameters:
//  - value : Value to set for the hasPhysicalDevice property.
func (m *CreateRequestBody) SetHasPhysicalDevice(value *bool)() {
    m.hasPhysicalDevice = value
}
// Sets the manufacturer property value. 
// Parameters:
//  - value : Value to set for the manufacturer property.
func (m *CreateRequestBody) SetManufacturer(value *string)() {
    m.manufacturer = value
}
// Sets the model property value. 
// Parameters:
//  - value : Value to set for the model property.
func (m *CreateRequestBody) SetModel(value *string)() {
    m.model = value
}
// Sets the physicalDeviceId property value. 
// Parameters:
//  - value : Value to set for the physicalDeviceId property.
func (m *CreateRequestBody) SetPhysicalDeviceId(value *string)() {
    m.physicalDeviceId = value
}
