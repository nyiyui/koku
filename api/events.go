package api

import (
	"fmt"
)

// CourseEvents generates Event s from the course schedule.
//
// TODO: use MeScheduleResp or MeTimetableResp to generate events for 2+ weeks.
func (c *Client) CourseEvents() (evs []Event, err error) {
	resp := MeScheduleWeekResp{}
	err = c.Do(MeScheduleWeekReq{}, &resp)
	if err != nil {
		return nil, fmt.Errorf("req: %w", err)
	}
	evs = make([]Event, 0, len(resp)*2)
	var ev Event
	for _, day := range resp {
		for _, sch := range day {
			ev, err = sch.Event(c)
			if err != nil {
				return nil, err
			}
			evs = append(evs, ev)
		}
	}
	return evs, nil
}
