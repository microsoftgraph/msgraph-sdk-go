package getprinterarchivedprintjobswithprinteridwithstartdatetimewithenddatetime

import (
    i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91 "github.com/microsoft/kiota-abstractions-go/serialization"
)

// GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeMember1 provides operations to call the getPrinterArchivedPrintJobs method.
type GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeMember1 struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{}
}
// NewGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeMember1 instantiates a new getPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeMember1 and sets the default values.
func NewGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeMember1()(*GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeMember1) {
    m := &GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeMember1{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// CreateGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeMember1FromDiscriminatorValue creates a new instance of the appropriate class based on discriminator value
func CreateGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeMember1FromDiscriminatorValue(parseNode i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.Parsable, error) {
    return NewGetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeMember1(), nil
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeMember1) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeMember1) GetFieldDeserializers()(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error)) {
    res := make(map[string]func(i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.ParseNode)(error))
    return res
}
// Serialize serializes information the current object
func (m *GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeMember1) Serialize(writer i878a80d2330e89d26896388a3f487eef27b0a0e6c010c493bf80be1452208f91.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetPrinterArchivedPrintJobsWithPrinterIdWithStartDateTimeWithEndDateTimeMember1) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
