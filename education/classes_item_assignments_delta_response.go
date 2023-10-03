package education

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// ClassesItemAssignmentsDeltaResponse 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ClassesItemAssignmentsDeltaResponse struct {
    ClassesItemAssignmentsDeltaGetResponse
}
// NewClassesItemAssignmentsDeltaResponse instantiates a new ClassesItemAssignmentsDeltaResponse and sets the default values.
func NewClassesItemAssignmentsDeltaResponse()(*ClassesItemAssignmentsDeltaResponse) {
    m := &ClassesItemAssignmentsDeltaResponse{
        ClassesItemAssignmentsDeltaGetResponse: *NewClassesItemAssignmentsDeltaGetResponse(),
    }
    return m
}
// CreateClassesItemAssignmentsDeltaResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateClassesItemAssignmentsDeltaResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewClassesItemAssignmentsDeltaResponse(), nil
}
// ClassesItemAssignmentsDeltaResponseable 
// Deprecated: This class is obsolete. Use deltaGetResponse instead.
type ClassesItemAssignmentsDeltaResponseable interface {
    ClassesItemAssignmentsDeltaGetResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
