package rolemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242 "github.com/microsoftgraph/msgraph-sdk-go/models"
)

type DirectoryRoleassignmentschedulesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse struct {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseCollectionPaginationCountResponse
}
// NewDirectoryRoleassignmentschedulesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse instantiates a new DirectoryRoleassignmentschedulesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse and sets the default values.
func NewDirectoryRoleassignmentschedulesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse()(*DirectoryRoleassignmentschedulesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse) {
    m := &DirectoryRoleassignmentschedulesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse{
        BaseCollectionPaginationCountResponse: *iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateDirectoryRoleassignmentschedulesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateDirectoryRoleassignmentschedulesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewDirectoryRoleassignmentschedulesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
// returns a map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error) when successful
func (m *DirectoryRoleassignmentschedulesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    res["value"] = func (n i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.CreateUnifiedRoleAssignmentScheduleFromDiscriminatorValue)
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleable, len(val))
            for i, v := range val {
                if v != nil {
                    res[i] = v.(iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleable)
                }
            }
            m.SetValue(res)
        }
        return nil
    }
    return res
}
// GetValue gets the value property value. The value property
// returns a []UnifiedRoleAssignmentScheduleable when successful
func (m *DirectoryRoleassignmentschedulesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse) GetValue()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleable) {
    val, err := m.GetBackingStore().Get("value")
    if err != nil {
        panic(err)
    }
    if val != nil {
        return val.([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleable)
    }
    return nil
}
// Serialize serializes information the current object
func (m *DirectoryRoleassignmentschedulesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    if m.GetValue() != nil {
        cast := make([]i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, len(m.GetValue()))
        for i, v := range m.GetValue() {
            if v != nil {
                cast[i] = v.(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable)
            }
        }
        err = writer.WriteCollectionOfObjectValues("value", cast)
        if err != nil {
            return err
        }
    }
    return nil
}
// SetValue sets the value property value. The value property
func (m *DirectoryRoleassignmentschedulesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponse) SetValue(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleable)() {
    err := m.GetBackingStore().Set("value", value)
    if err != nil {
        panic(err)
    }
}
type DirectoryRoleassignmentschedulesFilterbycurrentuserwithonFilterByCurrentUserWithOnGetResponseable interface {
    iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.BaseCollectionPaginationCountResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
    GetValue()([]iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleable)
    SetValue(value []iadcd81124412c61e647227ecfc4449d8bba17de0380ddda76f641a29edf2b242.UnifiedRoleAssignmentScheduleable)()
}
