package services

import (
	"fmt"
	"sort"

	"github.com/japnoor/daelog/internal/db"
	"github.com/japnoor/daelog/internal/models"
)

const sessionGapMs int64 = 30 * 60 * 1000

type SessionService struct {
	events *EventService
}

func NewSessionService(db *db.MongoDB) *SessionService {
	return &SessionService{events: NewEventService(db)}
}

func (s *SessionService) GetByDate(from, to int64) ([]models.Session, error) {
	events, err := s.events.GetByDate(from, to)
	if err != nil {
		return nil, err
	}
	return groupSessions(events), nil
}

func groupSessions(events []models.Event) []models.Session {
	if len(events) == 0 {
		return []models.Session{}
	}

	sort.Slice(events, func(i, j int) bool {
		return events[i].Ts < events[j].Ts
	})

	var sessions []models.Session

	for _, event := range events {
		if len(sessions) == 0 || event.Ts-sessions[len(sessions)-1].EndTs > sessionGapMs {
			repos := []string{}
			if event.Repo != "" {
				repos = []string{event.Repo}
			}
			sessions = append(sessions, models.Session{
				ID:      fmt.Sprintf("session-%d", event.Ts),
				StartTs: event.Ts,
				EndTs:   event.Ts,
				Duration: 0,
				Events:  []models.Event{event},
				Repos:   repos,
			})
		} else {
			last := &sessions[len(sessions)-1]
			last.Events = append(last.Events, event)
			last.EndTs = event.Ts
			last.Duration = last.EndTs - last.StartTs
			if event.Repo != "" && !repoExists(last.Repos, event.Repo) {
				last.Repos = append(last.Repos, event.Repo)
			}
		}
	}

	// reverse so newest session comes first
	for i, j := 0, len(sessions)-1; i < j; i, j = i+1, j-1 {
		sessions[i], sessions[j] = sessions[j], sessions[i]
	}

	return sessions
}

func repoExists(repos []string, repo string) bool {
	for _, r := range repos {
		if r == repo {
			return true
		}
	}
	return false
}
