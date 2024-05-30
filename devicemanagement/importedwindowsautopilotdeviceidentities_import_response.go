package devicemanagement

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// Deprecated: This class is obsolete. Use ImportedwindowsautopilotdeviceidentitiesImportPostResponseable instead.
type ImportedwindowsautopilotdeviceidentitiesImportResponse struct {
    ImportedwindowsautopilotdeviceidentitiesImportPostResponse
}
// NewImportedwindowsautopilotdeviceidentitiesImportResponse instantiates a new ImportedwindowsautopilotdeviceidentitiesImportResponse and sets the default values.
func NewImportedwindowsautopilotdeviceidentitiesImportResponse()(*ImportedwindowsautopilotdeviceidentitiesImportResponse) {
    m := &ImportedwindowsautopilotdeviceidentitiesImportResponse{
        ImportedwindowsautopilotdeviceidentitiesImportPostResponse: *NewImportedwindowsautopilotdeviceidentitiesImportPostResponse(),
    }
    return m
}
// CreateImportedwindowsautopilotdeviceidentitiesImportResponseFromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
// returns a Parsable when successful
func CreateImportedwindowsautopilotdeviceidentitiesImportResponseFromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewImportedwindowsautopilotdeviceidentitiesImportResponse(), nil
}
// Deprecated: This class is obsolete. Use ImportedwindowsautopilotdeviceidentitiesImportPostResponseable instead.
type ImportedwindowsautopilotdeviceidentitiesImportResponseable interface {
    ImportedwindowsautopilotdeviceidentitiesImportPostResponseable
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable
}
