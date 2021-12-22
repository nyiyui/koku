// This file is part of a program named Sisikyō or Sisikyo.
//
// Copyright (C) 2019 Ken Shibata <kenxshibata@gmail.com>
//
// License as published by the Free Software Foundation, either version 1 of the License, or (at your option) any later
// version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
// warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along with this program. If not, see
// <https://www.gnu.org/licenses/>.

package api

import (
	"net/http"
	"net/url"
	"sync"
	"time"
)

// eventsURL is the url for a Client.Events API call.
var eventsURL, _ = url.Parse("events")

// EventsReq has the options for a Client.Events API call.
type EventsReq struct {
	Start *time.Time
	End   *time.Time
}

func (req EventsReq) Req(c *Client) (*http.Request, error) {
	return queryReq(eventsURL, req, c)
}

// Query serializes the options in EventsReq to a new url.Values.
func (req EventsReq) Query() (v url.Values) {
	v = url.Values{}
	v.Add("format", "json")
	if req.Start != nil {
		v.Add("start", req.Start.Format(time.RFC3339))
	}
	if req.End != nil {
		v.Add("end", req.End.Format(time.RFC3339))
	}
	return
}

// EventsResp is the expected response from a Client.Events API call.
type EventsResp = []Event

// Deprecated: use Client.Do instead.
func (c *Client) Events(req EventsReq) (resp EventsResp, err error) {
	err = c.Do(req, &resp)
	return
}

type Events struct {
	_ [0]sync.Locker
	m map[time.Time][]Event
}

func NewEvents(resp EventsResp) *Events {
	evs := &Events{m: map[time.Time][]Event{}}
	for _, ev := range resp {
		truncated := ev.Start.Truncate(24 * time.Hour)
		evs.m[truncated] = append(evs.m[truncated], ev)
	}
	return evs
}
