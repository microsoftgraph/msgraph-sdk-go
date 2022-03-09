package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
    i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557 "github.com/microsoftgraph/msgraph-sdk-go/models/microsoft/graph/callrecords"
)

// CloudCommunicationsable 
type CloudCommunicationsable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetCallRecords()([]i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.CallRecordable)
    GetCalls()([]Callable)
    GetOnlineMeetings()([]OnlineMeetingable)
    GetPresences()([]Presenceable)
    SetCallRecords(value []i6afae973b07adf053fd7fc51b2f43be439d7fa83efb2b91811395e1128338557.CallRecordable)()
    SetCalls(value []Callable)()
    SetOnlineMeetings(value []OnlineMeetingable)()
    SetPresences(value []Presenceable)()
}
