package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ComplianceManagementPartnerCollectionResponse 
type ComplianceManagementPartnerCollectionResponse struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // The OdataNextLink property
    odataNextLink *string
    // The value property
    value []ComplianceManagementPartnerable
}
// NewComplianceManagementPartnerCollectionResponse instantiates a new ComplianceManagementPartnerCollectionResponse and sets the default values.
func NewComplianceManagementPartnerCollectionResponse()(*ComplianceManagementPartnerCollectionResponse) {
    m := &ComplianceManagementPartnerCollectionResponse{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateComplianceManagementPartnerCollectionResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateComplianceManagementPartnerCollectionResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewComplianceManagementPartnerCollectionResponse(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *ComplianceManagementPartnerCollectionResponse) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetFieldDeserializers the deserialization information for the current model
func (m *ComplianceManagementPartnerCollectionResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["@odata.nextLink"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataNextLink(val)
        }
        return nil
    }
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(CreateComplianceManagementPartnerFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]ComplianceManagementPartnerable, len(val))
            for i, v := range val {
                res[i] = v.(ComplianceManagementPartnerable)
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetOdataNextLink gets the @odata.nextLink property value. The OdataNextLink property
func (m *ComplianceManagementPartnerCollectionResponse) GetOdataNextLink()(*string) {
    return m.odataNextLink
}
// GetValue gets the value property value. The value property
func (m *ComplianceManagementPartnerCollectionResponse) GetValue()([]ComplianceManagementPartnerable) {
    return m.value
}
// Serialize serializes information the current object
func (m *ComplianceManagementPartnerCollectionResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteStringValue("@odata.nextLink", m.GetOdataNextLink())
        if err != nil {
            return err
        }
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("value", cast)
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
func (m *ComplianceManagementPartnerCollectionResponse) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetOdataNextLink sets the @odata.nextLink property value. The OdataNextLink property
func (m *ComplianceManagementPartnerCollectionResponse) SetOdataNextLink(value *string)() {
    m.odataNextLink = value
}
// SetValue sets the value property value. The value property
func (m *ComplianceManagementPartnerCollectionResponse) SetValue(value []ComplianceManagementPartnerable)() {
    m.value = value
}
