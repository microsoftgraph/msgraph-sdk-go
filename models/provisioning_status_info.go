package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ProvisioningStatusInfo 
type ProvisioningStatusInfo struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The errorInformation property
    errorInformation ProvisioningErrorInfoable;
    // Possible values are: success, warning, failure, skipped, unknownFutureValue.
    status *ProvisioningResult;
}
// NewProvisioningStatusInfo instantiates a new provisioningStatusInfo and sets the default values.
func NewProvisioningStatusInfo()(*ProvisioningStatusInfo) {
    m := &ProvisioningStatusInfo{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateProvisioningStatusInfoFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateProvisioningStatusInfoFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewProvisioningStatusInfo(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ProvisioningStatusInfo) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetErrorInformation gets the errorInformation property value. The errorInformation property
func (m *ProvisioningStatusInfo) GetErrorInformation()(ProvisioningErrorInfoable) {
    if m == nil {
        return nil
    } else {
        return m.errorInformation
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ProvisioningStatusInfo) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["errorInformation"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetObjectValue(CreateProvisioningErrorInfoFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetErrorInformation(val.(ProvisioningErrorInfoable))
        }
        return nil
    }
    res["status"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetEnumValue(ParseProvisioningResult)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetStatus(val.(*ProvisioningResult))
        }
        return nil
    }
    return res
}
// GetStatus gets the status property value. Possible values are: success, warning, failure, skipped, unknownFutureValue.
func (m *ProvisioningStatusInfo) GetStatus()(*ProvisioningResult) {
    if m == nil {
        return nil
    } else {
        return m.status
    }
}
// Serialize serializes information the current object
func (m *ProvisioningStatusInfo) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteObjectValue("errorInformation", m.GetErrorInformation())
        if err != nil {
            return err
        }
    }
    if m.GetStatus() != nil {
        cast := (*m.GetStatus()).String()
        err := writer.WriteStringValue("status", &cast)
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
func (m *ProvisioningStatusInfo) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetErrorInformation sets the errorInformation property value. The errorInformation property
func (m *ProvisioningStatusInfo) SetErrorInformation(value ProvisioningErrorInfoable)() {
    if m != nil {
        m.errorInformation = value
    }
}
// SetStatus sets the status property value. Possible values are: success, warning, failure, skipped, unknownFutureValue.
func (m *ProvisioningStatusInfo) SetStatus(value *ProvisioningResult)() {
    if m != nil {
        m.status = value
    }
}
