package create

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// CreateRequestBody provides operations to call the create method.
type CreateRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The certificateSigningRequest property
    certificateSigningRequest iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintCertificateSigningRequestable;
    // The connectorId property
    connectorId *string;
    // The displayName property
    displayName *string;
    // The hasPhysicalDevice property
    hasPhysicalDevice *bool;
    // The manufacturer property
    manufacturer *string;
    // The model property
    model *string;
    // The physicalDeviceId property
    physicalDeviceId *string;
}
// NewCreateRequestBody instantiates a new createRequestBody and sets the default values.
func NewCreateRequestBody()(*CreateRequestBody) {
    m := &CreateRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCreateRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCreateRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCreateRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CreateRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetCertificateSigningRequest gets the certificateSigningRequest property value. The certificateSigningRequest property
func (m *CreateRequestBody) GetCertificateSigningRequest()(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintCertificateSigningRequestable) {
    if m == nil {
        return nil
    } else {
        return m.certificateSigningRequest
    }
}
// GetConnectorId gets the connectorId property value. The connectorId property
func (m *CreateRequestBody) GetConnectorId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.connectorId
    }
}
// GetDisplayName gets the displayName property value. The displayName property
func (m *CreateRequestBody) GetDisplayName()(*string) {
    if m == nil {
        return nil
    } else {
        return m.displayName
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CreateRequestBody) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["certificateSigningRequest"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreatePrintCertificateSigningRequestFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetCertificateSigningRequest(val.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintCertificateSigningRequestable))
        }
        return nil
    }
    res["connectorId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetConnectorId(val)
        }
        return nil
    }
    res["displayName"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetDisplayName(val)
        }
        return nil
    }
    res["hasPhysicalDevice"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetHasPhysicalDevice(val)
        }
        return nil
    }
    res["manufacturer"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetManufacturer(val)
        }
        return nil
    }
    res["model"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetModel(val)
        }
        return nil
    }
    res["physicalDeviceId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
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
// GetHasPhysicalDevice gets the hasPhysicalDevice property value. The hasPhysicalDevice property
func (m *CreateRequestBody) GetHasPhysicalDevice()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.hasPhysicalDevice
    }
}
// GetManufacturer gets the manufacturer property value. The manufacturer property
func (m *CreateRequestBody) GetManufacturer()(*string) {
    if m == nil {
        return nil
    } else {
        return m.manufacturer
    }
}
// GetModel gets the model property value. The model property
func (m *CreateRequestBody) GetModel()(*string) {
    if m == nil {
        return nil
    } else {
        return m.model
    }
}
// GetPhysicalDeviceId gets the physicalDeviceId property value. The physicalDeviceId property
func (m *CreateRequestBody) GetPhysicalDeviceId()(*string) {
    if m == nil {
        return nil
    } else {
        return m.physicalDeviceId
    }
}
// Serialize serializes information the current object
func (m *CreateRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
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
// SetCertificateSigningRequest sets the certificateSigningRequest property value. The certificateSigningRequest property
func (m *CreateRequestBody) SetCertificateSigningRequest(value iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.PrintCertificateSigningRequestable)() {
    if m != nil {
        m.certificateSigningRequest = value
    }
}
// SetConnectorId sets the connectorId property value. The connectorId property
func (m *CreateRequestBody) SetConnectorId(value *string)() {
    if m != nil {
        m.connectorId = value
    }
}
// SetDisplayName sets the displayName property value. The displayName property
func (m *CreateRequestBody) SetDisplayName(value *string)() {
    if m != nil {
        m.displayName = value
    }
}
// SetHasPhysicalDevice sets the hasPhysicalDevice property value. The hasPhysicalDevice property
func (m *CreateRequestBody) SetHasPhysicalDevice(value *bool)() {
    if m != nil {
        m.hasPhysicalDevice = value
    }
}
// SetManufacturer sets the manufacturer property value. The manufacturer property
func (m *CreateRequestBody) SetManufacturer(value *string)() {
    if m != nil {
        m.manufacturer = value
    }
}
// SetModel sets the model property value. The model property
func (m *CreateRequestBody) SetModel(value *string)() {
    if m != nil {
        m.model = value
    }
}
// SetPhysicalDeviceId sets the physicalDeviceId property value. The physicalDeviceId property
func (m *CreateRequestBody) SetPhysicalDeviceId(value *string)() {
    if m != nil {
        m.physicalDeviceId = value
    }
}
