package checkmembergroups

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CheckMemberGroupsRequestBody provides operations to call the checkMemberGroups method.
type CheckMemberGroupsRequestBody struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
    // The groupIds property
    groupIds []string;
}
// NewCheckMemberGroupsRequestBody instantiates a new checkMemberGroupsRequestBody and sets the default values.
func NewCheckMemberGroupsRequestBody()(*CheckMemberGroupsRequestBody) {
    m := &CheckMemberGroupsRequestBody{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateCheckMemberGroupsRequestBodyFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCheckMemberGroupsRequestBodyFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCheckMemberGroupsRequestBody(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *CheckMemberGroupsRequestBody) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CheckMemberGroupsRequestBody) GetFieldDeserializers()(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    res["groupIds"] = func (o interface{}, n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetGroupIds(res)
        }
        return nil
    }
    return res
}
// GetGroupIds gets the groupIds property value. The groupIds property
func (m *CheckMemberGroupsRequestBody) GetGroupIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.groupIds
    }
}
// Serialize serializes information the current object
func (m *CheckMemberGroupsRequestBody) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    if m.GetGroupIds() != nil {
        err := writer.WriteCollectionOfStringValues("groupIds", m.GetGroupIds())
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
func (m *CheckMemberGroupsRequestBody) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
// SetGroupIds sets the groupIds property value. The groupIds property
func (m *CheckMemberGroupsRequestBody) SetGroupIds(value []string)() {
    if m != nil {
        m.groupIds = value
    }
}
