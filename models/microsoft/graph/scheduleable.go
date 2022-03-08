package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Scheduleable 
type Scheduleable interface {
    Entityable
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable
    GetEnabled()(*bool)
    GetOfferShiftRequests()([]OfferShiftRequestable)
    GetOfferShiftRequestsEnabled()(*bool)
    GetOpenShiftChangeRequests()([]OpenShiftChangeRequestable)
    GetOpenShifts()([]OpenShiftable)
    GetOpenShiftsEnabled()(*bool)
    GetProvisionStatus()(*OperationStatus)
    GetProvisionStatusCode()(*string)
    GetSchedulingGroups()([]SchedulingGroupable)
    GetShifts()([]Shiftable)
    GetSwapShiftsChangeRequests()([]SwapShiftsChangeRequestable)
    GetSwapShiftsRequestsEnabled()(*bool)
    GetTimeClockEnabled()(*bool)
    GetTimeOffReasons()([]TimeOffReasonable)
    GetTimeOffRequests()([]TimeOffRequestable)
    GetTimeOffRequestsEnabled()(*bool)
    GetTimesOff()([]TimeOffable)
    GetTimeZone()(*string)
    GetWorkforceIntegrationIds()([]string)
    SetEnabled(value *bool)()
    SetOfferShiftRequests(value []OfferShiftRequestable)()
    SetOfferShiftRequestsEnabled(value *bool)()
    SetOpenShiftChangeRequests(value []OpenShiftChangeRequestable)()
    SetOpenShifts(value []OpenShiftable)()
    SetOpenShiftsEnabled(value *bool)()
    SetProvisionStatus(value *OperationStatus)()
    SetProvisionStatusCode(value *string)()
    SetSchedulingGroups(value []SchedulingGroupable)()
    SetShifts(value []Shiftable)()
    SetSwapShiftsChangeRequests(value []SwapShiftsChangeRequestable)()
    SetSwapShiftsRequestsEnabled(value *bool)()
    SetTimeClockEnabled(value *bool)()
    SetTimeOffReasons(value []TimeOffReasonable)()
    SetTimeOffRequests(value []TimeOffRequestable)()
    SetTimeOffRequestsEnabled(value *bool)()
    SetTimesOff(value []TimeOffable)()
    SetTimeZone(value *string)()
    SetWorkforceIntegrationIds(value []string)()
}
