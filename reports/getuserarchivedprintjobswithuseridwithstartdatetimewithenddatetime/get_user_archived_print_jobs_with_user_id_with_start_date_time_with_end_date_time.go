package getuserarchivedprintjobswithuseridwithstartdatetimewithenddatetime

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime 
type GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime struct {
    // Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
    additionalData map[string]interface{};
}
// NewGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime instantiates a new getUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime and sets the default values.
func NewGetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime()(*GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime) {
    m := &GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime{
    }
    m.SetAdditionalData(make(map[string]interface{}));
    return m
}
// GetAdditionalData gets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime) GetAdditionalData()(map[string]interface{}) {
    if m == nil {
        return nil
    } else {
        return m.additionalData
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := make(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error))
    return res
}
func (m *GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    {
        err := writer.WriteAdditionalData(m.GetAdditionalData())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetAdditionalData sets the additionalData property value. Stores additional data not described in the OpenAPI description found when deserializing. Can be used for serialization as well.
func (m *GetUserArchivedPrintJobsWithUserIdWithStartDateTimeWithEndDateTime) SetAdditionalData(value map[string]interface{})() {
    if m != nil {
        m.additionalData = value
    }
}
