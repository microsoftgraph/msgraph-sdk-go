package assignlicense

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// AssignLicenseRequestBody provides operations to call the assignLicense method.
type AssignLicenseRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The addLicenses property
    addLicenses []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssignedLicenseable;
    // The removeLicenses property
    removeLicenses []string;
}
// NewAssignLicenseRequestBody instantiates a new assignLicenseRequestBody and sets the default values.
func NewAssignLicenseRequestBody()(*AssignLicenseRequestBody) {
    m := &AssignLicenseRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateAssignLicenseRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateAssignLicenseRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewAssignLicenseRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *AssignLicenseRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetAddLicenses gets the addLicenses property value. The addLicenses property
func (m *AssignLicenseRequestBody) GetAddLicenses()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssignedLicenseable) {
    if m == nil {
        return nil
    } else {
        return m.addLicenses
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *AssignLicenseRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["addLicenses"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateAssignedLicenseFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssignedLicenseable, len(val))
            for i, v := range val {
                res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssignedLicenseable)
            }
            m.SetAddLicenses(res)
        }
        return nil
    }
    res["removeLicenses"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetRemoveLicenses(res)
        }
        return nil
    }
    return res
}
// GetRemoveLicenses gets the removeLicenses property value. The removeLicenses property
func (m *AssignLicenseRequestBody) GetRemoveLicenses()([]string) {
    if m == nil {
        return nil
    } else {
        return m.removeLicenses
    }
}
// Serialize serializes information the current object
func (m *AssignLicenseRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetAddLicenses() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetAddLicenses()))
        for i, v := range m.GetAddLicenses() {
            cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
        }
        err := writer.WriteCollectionOfObjectValues("addLicenses", cast)
        if err != nil {
            return err
        }
    }
    if m.GetRemoveLicenses() != nil {
        err := writer.WriteCollectionOfStringValues("removeLicenses", m.GetRemoveLicenses())
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
func (m *AssignLicenseRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetAddLicenses sets the addLicenses property value. The addLicenses property
func (m *AssignLicenseRequestBody) SetAddLicenses(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.AssignedLicenseable)() {
    if m != nil {
        m.addLicenses = value
    }
}
// SetRemoveLicenses sets the removeLicenses property value. The removeLicenses property
func (m *AssignLicenseRequestBody) SetRemoveLicenses(value []string)() {
    if m != nil {
        m.removeLicenses = value
    }
}
