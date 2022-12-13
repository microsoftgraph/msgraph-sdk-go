package models

import (
    i2bacd9b8d8db2e77ee2b5c5ccb19d679c36f920b8fee9d786a0adafff458afcd "github.com/google/UUID"
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// AssignedLicense 
type AssignedLicense struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
    // A collection of the unique identifiers for plans that have been disabled.
    disabledPlans []UUID
    // The OdataType property
    odataType *string
    // The unique identifier for the SKU.
    skuId *UUID
}
// NewAssignedLicense instantiates a new assignedLicense and sets the default values.
func NewAssignedLicense()(*AssignedLicense) {
    m := &AssignedLicense{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAssignedLicenseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignedLicenseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignedLicense(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignedLicense) GetAdditionalData()(map[string]interface{}) {
    return m.additionalData
}
// GetDisabledPlans gets the disabledPlans property value. A collection of the unique identifiers for plans that have been disabled.
func (m *AssignedLicense) GetDisabledPlans()([]UUID) {
    return m.disabledPlans
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignedLicense) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["disabledPlans"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("uUID")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]UUID, len(val))
            for i, v := range val {
                res[i] = *(v.(*UUID))
            }
            m.SetDisabledPlans(res)
        }
        return nil
    }
    res["@odata.type"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOdataType(val)
        }
        return nil
    }
    res["skuId"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetUUIDValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSkuId(val)
        }
        return nil
    }
    return res
}
// GetOdataType gets the @odata.type property value. The OdataType property
func (m *AssignedLicense) GetOdataType()(*string) {
    return m.odataType
}
// GetSkuId gets the skuId property value. The unique identifier for the SKU.
func (m *AssignedLicense) GetSkuId()(*UUID) {
    return m.skuId
}
// Serialize serializes information the current object
func (m *AssignedLicense) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetDisabledPlans() != nil {
        err := writer.WriteCollectionOfUUIDValues("disabledPlans", m.GetDisabledPlans())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteStringValue("@odata.type", m.GetOdataType())
        if err != nil {
            return err
        }
    }
    {
        err := writer.WriteUUIDValue("skuId", m.GetSkuId())
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
func (m *AssignedLicense) SetAdditionalData(value map[string]interface{})() {
    m.additionalData = value
}
// SetDisabledPlans sets the disabledPlans property value. A collection of the unique identifiers for plans that have been disabled.
func (m *AssignedLicense) SetDisabledPlans(value []UUID)() {
    m.disabledPlans = value
}
// SetOdataType sets the @odata.type property value. The OdataType property
func (m *AssignedLicense) SetOdataType(value *string)() {
    m.odataType = value
}
// SetSkuId sets the skuId property value. The unique identifier for the SKU.
func (m *AssignedLicense) SetSkuId(value *UUID)() {
    m.skuId = value
}
