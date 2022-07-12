package getapplicablecontenttypesforlistwithlistid

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetApplicableContentTypesForListWithListIdMember1 provides operations to call the getApplicableContentTypesForList method.
type GetApplicableContentTypesForListWithListIdMember1 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
}
// NewGetApplicableContentTypesForListWithListIdMember1 instantiates a new getApplicableContentTypesForListWithListIdMember1 and sets the default values.
func NewGetApplicableContentTypesForListWithListIdMember1()(*GetApplicableContentTypesForListWithListIdMember1) {
    m := &GetApplicableContentTypesForListWithListIdMember1{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetApplicableContentTypesForListWithListIdMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetApplicableContentTypesForListWithListIdMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetApplicableContentTypesForListWithListIdMember1(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetApplicableContentTypesForListWithListIdMember1) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetApplicableContentTypesForListWithListIdMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *GetApplicableContentTypesForListWithListIdMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetApplicableContentTypesForListWithListIdMember1) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
