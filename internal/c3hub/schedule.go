package c3hub

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (service *service) GetMySchedule(
	ctx context.Context,
) ([]Event, error) {
	r, err := http.NewRequest(http.MethodGet, service.config.BaseURL+"/me/events", nil)
	if err != nil {
		return nil, err
	}
	r.Header.Add("Cookie", "38C3_SESSION="+service.config.SessionCookie)
	r = r.WithContext(ctx)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var favs []string
	err = json.NewDecoder(resp.Body).Decode(&favs)
	if err != nil {
		return nil, err
	}

	events, err := service.getEvents(ctx)
	if err != nil {
		return nil, err
	}

	favedEvents := make([]Event, 0)
	for _, ev := range events {
		for _, fav := range favs {
			if ev.ID == fav {
				favedEvents = append(favedEvents, ev)
				break
			}
		}
	}

	return favedEvents, nil
}

func (service *service) getEvents(
	ctx context.Context,
) ([]Event, error) {
	type HubEvent struct {
		ID       string
		Name     string
		Track    string
		Assembly string
		Room     *string
		Location *string
		//Description   string
		ScheduleStart string `json:"schedule_start"`
		ScheduleEnd   string `json:"schedule_end"`
	}

	r, err := http.NewRequest(http.MethodGet, service.config.BaseURL+"/events", nil)
	if err != nil {
		return nil, err
	}
	r.Header.Add("Cookie", "38C3_SESSION="+service.config.SessionCookie)
	r = r.WithContext(ctx)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var intEvents []HubEvent
	err = json.NewDecoder(resp.Body).Decode(&intEvents)
	if err != nil {
		return nil, err
	}

	rooms, err := service.getRooms(ctx)
	if err != nil {
		return nil, err
	}
	roomIDToRoom := make(map[string]Room, len(rooms))
	for _, room := range rooms {
		roomIDToRoom[room.ID] = room
	}

	events := make([]Event, len(intEvents))
	for i := range intEvents {
		start, err := time.Parse(time.RFC3339, intEvents[i].ScheduleStart)
		if err != nil {
			return nil, err
		}
		end, err := time.Parse(time.RFC3339, intEvents[i].ScheduleEnd)
		if err != nil {
			return nil, err
		}

		events[i] = Event{
			ID:        intEvents[i].ID,
			Name:      intEvents[i].Name,
			StartTime: start,
			EndTime:   end,
		}

		if intEvents[i].Room != nil {
			if room, ok := roomIDToRoom[*intEvents[i].Room]; ok {
				events[i].Room = room.Name
				if room.Assembly != "" {
					events[i].Room += " - " + room.Assembly
				}
				events[i].Room += " [" + room.Type + "]"
			}
		} else if intEvents[i].Location != nil {
			events[i].Room = *intEvents[i].Location
		}

	}

	return events, nil
}

func (service *service) getRooms(
	ctx context.Context,
) ([]Room, error) {
	type HubRoom struct {
		ID       string
		Name     string
		Assembly string
		RoomType string `json:"room_type"`
	}

	r, err := http.NewRequest(http.MethodGet, service.config.BaseURL+"/rooms", nil)
	if err != nil {
		return nil, err
	}
	r.Header.Add("Cookie", "38C3_SESSION="+service.config.SessionCookie)
	r = r.WithContext(ctx)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var intRooms []HubRoom
	err = json.NewDecoder(resp.Body).Decode(&intRooms)
	if err != nil {
		return nil, err
	}

	rooms := make([]Room, len(intRooms))
	for i := range intRooms {
		rooms[i] = Room{
			ID:       intRooms[i].ID,
			Name:     intRooms[i].Name,
			Assembly: intRooms[i].Assembly,
			Type:     intRooms[i].RoomType,
		}
	}

	return rooms, nil
}
