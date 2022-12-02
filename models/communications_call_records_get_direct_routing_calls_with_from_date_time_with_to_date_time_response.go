package models

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// CommunicationsCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse provides operations to call the getDirectRoutingCalls method.
type CommunicationsCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse struct {
    BaseCollectionPaginationCountResponse
}
// NewCommunicationsCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse instantiates a new CommunicationsCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse and sets the default values.
func NewCommunicationsCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse()(*CommunicationsCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse) {
    m := &CommunicationsCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse{
        BaseCollectionPaginationCountResponse: *NewBaseCollectionPaginationCountResponse(),
    }
    return m
}
// CreateCommunicationsCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateCommunicationsCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewCommunicationsCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse(), nil
}
// GetFieldDeserializers the deserialization information for the current model
func (m *CommunicationsCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := m.BaseCollectionPaginationCountResponse.GetFieldDeserializers()
    return res
}
// Serialize serializes information the current object
func (m *CommunicationsCallRecordsGetDirectRoutingCallsWithFromDateTimeWithToDateTimeResponse) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    err := m.BaseCollectionPaginationCountResponse.Serialize(writer)
    if err != nil {
        return err
    }
    return nil
}
