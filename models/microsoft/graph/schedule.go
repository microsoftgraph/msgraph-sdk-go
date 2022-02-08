package graph

import (
    i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55 "github.com/microsoft/kiota/abstractions/go/serialization"
)

// Schedule 
type Schedule struct {
    Entity
    // Indicates whether the schedule is enabled for the team. Required.
    enabled *bool;
    // 
    offerShiftRequests []OfferShiftRequest;
    // Indicates whether offer shift requests are enabled for the schedule.
    offerShiftRequestsEnabled *bool;
    // 
    openShiftChangeRequests []OpenShiftChangeRequest;
    // 
    openShifts []OpenShift;
    // Indicates whether open shifts are enabled for the schedule.
    openShiftsEnabled *bool;
    // The status of the schedule provisioning. The possible values are notStarted, running, completed, failed.
    provisionStatus *OperationStatus;
    // Additional information about why schedule provisioning failed.
    provisionStatusCode *string;
    // The logical grouping of users in the schedule (usually by role).
    schedulingGroups []SchedulingGroup;
    // The shifts in the schedule.
    shifts []Shift;
    // 
    swapShiftsChangeRequests []SwapShiftsChangeRequest;
    // Indicates whether swap shifts requests are enabled for the schedule.
    swapShiftsRequestsEnabled *bool;
    // Indicates whether time clock is enabled for the schedule.
    timeClockEnabled *bool;
    // The set of reasons for a time off in the schedule.
    timeOffReasons []TimeOffReason;
    // 
    timeOffRequests []TimeOffRequest;
    // Indicates whether time off requests are enabled for the schedule.
    timeOffRequestsEnabled *bool;
    // The instances of times off in the schedule.
    timesOff []TimeOff;
    // Indicates the time zone of the schedule team using tz database format. Required.
    timeZone *string;
    // 
    workforceIntegrationIds []string;
}
// NewSchedule instantiates a new schedule and sets the default values.
func NewSchedule()(*Schedule) {
    m := &Schedule{
        Entity: *NewEntity(),
    }
    return m
}
// GetEnabled gets the enabled property value. Indicates whether the schedule is enabled for the team. Required.
func (m *Schedule) GetEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.enabled
    }
}
// GetOfferShiftRequests gets the offerShiftRequests property value. 
func (m *Schedule) GetOfferShiftRequests()([]OfferShiftRequest) {
    if m == nil {
        return nil
    } else {
        return m.offerShiftRequests
    }
}
// GetOfferShiftRequestsEnabled gets the offerShiftRequestsEnabled property value. Indicates whether offer shift requests are enabled for the schedule.
func (m *Schedule) GetOfferShiftRequestsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.offerShiftRequestsEnabled
    }
}
// GetOpenShiftChangeRequests gets the openShiftChangeRequests property value. 
func (m *Schedule) GetOpenShiftChangeRequests()([]OpenShiftChangeRequest) {
    if m == nil {
        return nil
    } else {
        return m.openShiftChangeRequests
    }
}
// GetOpenShifts gets the openShifts property value. 
func (m *Schedule) GetOpenShifts()([]OpenShift) {
    if m == nil {
        return nil
    } else {
        return m.openShifts
    }
}
// GetOpenShiftsEnabled gets the openShiftsEnabled property value. Indicates whether open shifts are enabled for the schedule.
func (m *Schedule) GetOpenShiftsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.openShiftsEnabled
    }
}
// GetProvisionStatus gets the provisionStatus property value. The status of the schedule provisioning. The possible values are notStarted, running, completed, failed.
func (m *Schedule) GetProvisionStatus()(*OperationStatus) {
    if m == nil {
        return nil
    } else {
        return m.provisionStatus
    }
}
// GetProvisionStatusCode gets the provisionStatusCode property value. Additional information about why schedule provisioning failed.
func (m *Schedule) GetProvisionStatusCode()(*string) {
    if m == nil {
        return nil
    } else {
        return m.provisionStatusCode
    }
}
// GetSchedulingGroups gets the schedulingGroups property value. The logical grouping of users in the schedule (usually by role).
func (m *Schedule) GetSchedulingGroups()([]SchedulingGroup) {
    if m == nil {
        return nil
    } else {
        return m.schedulingGroups
    }
}
// GetShifts gets the shifts property value. The shifts in the schedule.
func (m *Schedule) GetShifts()([]Shift) {
    if m == nil {
        return nil
    } else {
        return m.shifts
    }
}
// GetSwapShiftsChangeRequests gets the swapShiftsChangeRequests property value. 
func (m *Schedule) GetSwapShiftsChangeRequests()([]SwapShiftsChangeRequest) {
    if m == nil {
        return nil
    } else {
        return m.swapShiftsChangeRequests
    }
}
// GetSwapShiftsRequestsEnabled gets the swapShiftsRequestsEnabled property value. Indicates whether swap shifts requests are enabled for the schedule.
func (m *Schedule) GetSwapShiftsRequestsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.swapShiftsRequestsEnabled
    }
}
// GetTimeClockEnabled gets the timeClockEnabled property value. Indicates whether time clock is enabled for the schedule.
func (m *Schedule) GetTimeClockEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.timeClockEnabled
    }
}
// GetTimeOffReasons gets the timeOffReasons property value. The set of reasons for a time off in the schedule.
func (m *Schedule) GetTimeOffReasons()([]TimeOffReason) {
    if m == nil {
        return nil
    } else {
        return m.timeOffReasons
    }
}
// GetTimeOffRequests gets the timeOffRequests property value. 
func (m *Schedule) GetTimeOffRequests()([]TimeOffRequest) {
    if m == nil {
        return nil
    } else {
        return m.timeOffRequests
    }
}
// GetTimeOffRequestsEnabled gets the timeOffRequestsEnabled property value. Indicates whether time off requests are enabled for the schedule.
func (m *Schedule) GetTimeOffRequestsEnabled()(*bool) {
    if m == nil {
        return nil
    } else {
        return m.timeOffRequestsEnabled
    }
}
// GetTimesOff gets the timesOff property value. The instances of times off in the schedule.
func (m *Schedule) GetTimesOff()([]TimeOff) {
    if m == nil {
        return nil
    } else {
        return m.timesOff
    }
}
// GetTimeZone gets the timeZone property value. Indicates the time zone of the schedule team using tz database format. Required.
func (m *Schedule) GetTimeZone()(*string) {
    if m == nil {
        return nil
    } else {
        return m.timeZone
    }
}
// GetWorkforceIntegrationIds gets the workforceIntegrationIds property value. 
func (m *Schedule) GetWorkforceIntegrationIds()([]string) {
    if m == nil {
        return nil
    } else {
        return m.workforceIntegrationIds
    }
}
// GetFieldDeserializers the deserialization information for the current model
func (m *Schedule) GetFieldDeserializers()(map[string]func(interface{}, i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode)(error)) {
    res := m.Entity.GetFieldDeserializers()
    res["enabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetEnabled(val)
        }
        return nil
    }
    res["offerShiftRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOfferShiftRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OfferShiftRequest, len(val))
            for i, v := range val {
                res[i] = *(v.(*OfferShiftRequest))
            }
            m.SetOfferShiftRequests(res)
        }
        return nil
    }
    res["offerShiftRequestsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOfferShiftRequestsEnabled(val)
        }
        return nil
    }
    res["openShiftChangeRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOpenShiftChangeRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OpenShiftChangeRequest, len(val))
            for i, v := range val {
                res[i] = *(v.(*OpenShiftChangeRequest))
            }
            m.SetOpenShiftChangeRequests(res)
        }
        return nil
    }
    res["openShifts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewOpenShift() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]OpenShift, len(val))
            for i, v := range val {
                res[i] = *(v.(*OpenShift))
            }
            m.SetOpenShifts(res)
        }
        return nil
    }
    res["openShiftsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetOpenShiftsEnabled(val)
        }
        return nil
    }
    res["provisionStatus"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetEnumValue(ParseOperationStatus)
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisionStatus(val.(*OperationStatus))
        }
        return nil
    }
    res["provisionStatusCode"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetProvisionStatusCode(val)
        }
        return nil
    }
    res["schedulingGroups"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSchedulingGroup() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SchedulingGroup, len(val))
            for i, v := range val {
                res[i] = *(v.(*SchedulingGroup))
            }
            m.SetSchedulingGroups(res)
        }
        return nil
    }
    res["shifts"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewShift() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]Shift, len(val))
            for i, v := range val {
                res[i] = *(v.(*Shift))
            }
            m.SetShifts(res)
        }
        return nil
    }
    res["swapShiftsChangeRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewSwapShiftsChangeRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]SwapShiftsChangeRequest, len(val))
            for i, v := range val {
                res[i] = *(v.(*SwapShiftsChangeRequest))
            }
            m.SetSwapShiftsChangeRequests(res)
        }
        return nil
    }
    res["swapShiftsRequestsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetSwapShiftsRequestsEnabled(val)
        }
        return nil
    }
    res["timeClockEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeClockEnabled(val)
        }
        return nil
    }
    res["timeOffReasons"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeOffReason() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeOffReason, len(val))
            for i, v := range val {
                res[i] = *(v.(*TimeOffReason))
            }
            m.SetTimeOffReasons(res)
        }
        return nil
    }
    res["timeOffRequests"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeOffRequest() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeOffRequest, len(val))
            for i, v := range val {
                res[i] = *(v.(*TimeOffRequest))
            }
            m.SetTimeOffRequests(res)
        }
        return nil
    }
    res["timeOffRequestsEnabled"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetBoolValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeOffRequestsEnabled(val)
        }
        return nil
    }
    res["timesOff"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfObjectValues(func () i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable { return NewTimeOff() })
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]TimeOff, len(val))
            for i, v := range val {
                res[i] = *(v.(*TimeOff))
            }
            m.SetTimesOff(res)
        }
        return nil
    }
    res["timeZone"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetStringValue()
        if err != nil {
            return err
        }
        if val != nil {
            m.SetTimeZone(val)
        }
        return nil
    }
    res["workforceIntegrationIds"] = func (o interface{}, n i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.ParseNode) error {
        val, err := n.GetCollectionOfPrimitiveValues("string")
        if err != nil {
            return err
        }
        if val != nil {
            res := make([]string, len(val))
            for i, v := range val {
                res[i] = *(v.(*string))
            }
            m.SetWorkforceIntegrationIds(res)
        }
        return nil
    }
    return res
}
func (m *Schedule) IsNil()(bool) {
    return m == nil
}
// Serialize serializes information the current object
func (m *Schedule) Serialize(writer i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.SerializationWriter)(error) {
    err := m.Entity.Serialize(writer)
    if err != nil {
        return err
    }
    {
        err = writer.WriteBoolValue("enabled", m.GetEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetOfferShiftRequests() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOfferShiftRequests()))
        for i, v := range m.GetOfferShiftRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("offerShiftRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("offerShiftRequestsEnabled", m.GetOfferShiftRequestsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetOpenShiftChangeRequests() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOpenShiftChangeRequests()))
        for i, v := range m.GetOpenShiftChangeRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("openShiftChangeRequests", cast)
        if err != nil {
            return err
        }
    }
    if m.GetOpenShifts() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetOpenShifts()))
        for i, v := range m.GetOpenShifts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("openShifts", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("openShiftsEnabled", m.GetOpenShiftsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetProvisionStatus() != nil {
        cast := (*m.GetProvisionStatus()).String()
        err = writer.WriteStringValue("provisionStatus", &cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("provisionStatusCode", m.GetProvisionStatusCode())
        if err != nil {
            return err
        }
    }
    if m.GetSchedulingGroups() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSchedulingGroups()))
        for i, v := range m.GetSchedulingGroups() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("schedulingGroups", cast)
        if err != nil {
            return err
        }
    }
    if m.GetShifts() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetShifts()))
        for i, v := range m.GetShifts() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("shifts", cast)
        if err != nil {
            return err
        }
    }
    if m.GetSwapShiftsChangeRequests() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetSwapShiftsChangeRequests()))
        for i, v := range m.GetSwapShiftsChangeRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("swapShiftsChangeRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("swapShiftsRequestsEnabled", m.GetSwapShiftsRequestsEnabled())
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("timeClockEnabled", m.GetTimeClockEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetTimeOffReasons() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTimeOffReasons()))
        for i, v := range m.GetTimeOffReasons() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("timeOffReasons", cast)
        if err != nil {
            return err
        }
    }
    if m.GetTimeOffRequests() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTimeOffRequests()))
        for i, v := range m.GetTimeOffRequests() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("timeOffRequests", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteBoolValue("timeOffRequestsEnabled", m.GetTimeOffRequestsEnabled())
        if err != nil {
            return err
        }
    }
    if m.GetTimesOff() != nil {
        cast := make([]i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable, len(m.GetTimesOff()))
        for i, v := range m.GetTimesOff() {
            temp := v
            cast[i] = i04eb5309aeaafadd28374d79c8471df9b267510b4dc2e3144c378c50f6fd7b55.Parsable(&temp)
        }
        err = writer.WriteCollectionOfObjectValues("timesOff", cast)
        if err != nil {
            return err
        }
    }
    {
        err = writer.WriteStringValue("timeZone", m.GetTimeZone())
        if err != nil {
            return err
        }
    }
    if m.GetWorkforceIntegrationIds() != nil {
        err = writer.WriteCollectionOfStringValues("workforceIntegrationIds", m.GetWorkforceIntegrationIds())
        if err != nil {
            return err
        }
    }
    return nil
}
// SetEnabled sets the enabled property value. Indicates whether the schedule is enabled for the team. Required.
func (m *Schedule) SetEnabled(value *bool)() {
    if m != nil {
        m.enabled = value
    }
}
// SetOfferShiftRequests sets the offerShiftRequests property value. 
func (m *Schedule) SetOfferShiftRequests(value []OfferShiftRequest)() {
    if m != nil {
        m.offerShiftRequests = value
    }
}
// SetOfferShiftRequestsEnabled sets the offerShiftRequestsEnabled property value. Indicates whether offer shift requests are enabled for the schedule.
func (m *Schedule) SetOfferShiftRequestsEnabled(value *bool)() {
    if m != nil {
        m.offerShiftRequestsEnabled = value
    }
}
// SetOpenShiftChangeRequests sets the openShiftChangeRequests property value. 
func (m *Schedule) SetOpenShiftChangeRequests(value []OpenShiftChangeRequest)() {
    if m != nil {
        m.openShiftChangeRequests = value
    }
}
// SetOpenShifts sets the openShifts property value. 
func (m *Schedule) SetOpenShifts(value []OpenShift)() {
    if m != nil {
        m.openShifts = value
    }
}
// SetOpenShiftsEnabled sets the openShiftsEnabled property value. Indicates whether open shifts are enabled for the schedule.
func (m *Schedule) SetOpenShiftsEnabled(value *bool)() {
    if m != nil {
        m.openShiftsEnabled = value
    }
}
// SetProvisionStatus sets the provisionStatus property value. The status of the schedule provisioning. The possible values are notStarted, running, completed, failed.
func (m *Schedule) SetProvisionStatus(value *OperationStatus)() {
    if m != nil {
        m.provisionStatus = value
    }
}
// SetProvisionStatusCode sets the provisionStatusCode property value. Additional information about why schedule provisioning failed.
func (m *Schedule) SetProvisionStatusCode(value *string)() {
    if m != nil {
        m.provisionStatusCode = value
    }
}
// SetSchedulingGroups sets the schedulingGroups property value. The logical grouping of users in the schedule (usually by role).
func (m *Schedule) SetSchedulingGroups(value []SchedulingGroup)() {
    if m != nil {
        m.schedulingGroups = value
    }
}
// SetShifts sets the shifts property value. The shifts in the schedule.
func (m *Schedule) SetShifts(value []Shift)() {
    if m != nil {
        m.shifts = value
    }
}
// SetSwapShiftsChangeRequests sets the swapShiftsChangeRequests property value. 
func (m *Schedule) SetSwapShiftsChangeRequests(value []SwapShiftsChangeRequest)() {
    if m != nil {
        m.swapShiftsChangeRequests = value
    }
}
// SetSwapShiftsRequestsEnabled sets the swapShiftsRequestsEnabled property value. Indicates whether swap shifts requests are enabled for the schedule.
func (m *Schedule) SetSwapShiftsRequestsEnabled(value *bool)() {
    if m != nil {
        m.swapShiftsRequestsEnabled = value
    }
}
// SetTimeClockEnabled sets the timeClockEnabled property value. Indicates whether time clock is enabled for the schedule.
func (m *Schedule) SetTimeClockEnabled(value *bool)() {
    if m != nil {
        m.timeClockEnabled = value
    }
}
// SetTimeOffReasons sets the timeOffReasons property value. The set of reasons for a time off in the schedule.
func (m *Schedule) SetTimeOffReasons(value []TimeOffReason)() {
    if m != nil {
        m.timeOffReasons = value
    }
}
// SetTimeOffRequests sets the timeOffRequests property value. 
func (m *Schedule) SetTimeOffRequests(value []TimeOffRequest)() {
    if m != nil {
        m.timeOffRequests = value
    }
}
// SetTimeOffRequestsEnabled sets the timeOffRequestsEnabled property value. Indicates whether time off requests are enabled for the schedule.
func (m *Schedule) SetTimeOffRequestsEnabled(value *bool)() {
    if m != nil {
        m.timeOffRequestsEnabled = value
    }
}
// SetTimesOff sets the timesOff property value. The instances of times off in the schedule.
func (m *Schedule) SetTimesOff(value []TimeOff)() {
    if m != nil {
        m.timesOff = value
    }
}
// SetTimeZone sets the timeZone property value. Indicates the time zone of the schedule team using tz database format. Required.
func (m *Schedule) SetTimeZone(value *string)() {
    if m != nil {
        m.timeZone = value
    }
}
// SetWorkforceIntegrationIds sets the workforceIntegrationIds property value. 
func (m *Schedule) SetWorkforceIntegrationIds(value []string)() {
    if m != nil {
        m.workforceIntegrationIds = value
    }
}
