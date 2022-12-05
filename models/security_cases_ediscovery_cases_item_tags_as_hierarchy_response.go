package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// SecurityCasesEdiscoveryCasesItemTagsAsHierarchyResponse provides operations to call the asHierarchy method.
type SecurityCasesEdiscoveryCasesItemTagsAsHierarchyResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewSecurityCasesEdiscoveryCasesItemTagsAsHierarchyResponse instantiates a new SecurityCasesEdiscoveryCasesItemTagsAsHierarchyResponse and sets the default values.
func NewSecurityCasesEdiscoveryCasesItemTagsAsHierarchyResponse()(*SecurityCasesEdiscoveryCasesItemTagsAsHierarchyResponse) {
    m := &SecurityCasesEdiscoveryCasesItemTagsAsHierarchyResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateSecurityCasesEdiscoveryCasesItemTagsAsHierarchyResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateSecurityCasesEdiscoveryCasesItemTagsAsHierarchyResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewSecurityCasesEdiscoveryCasesItemTagsAsHierarchyResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *SecurityCasesEdiscoveryCasesItemTagsAsHierarchyResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *SecurityCasesEdiscoveryCasesItemTagsAsHierarchyResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
