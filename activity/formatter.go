package activity

import (
	"time"
)

type ActivityFormatter struct {
	ID        int        `json:"id"`
	Email     string     `json:"email"`
	Title     string     `json:"title"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

func FormatActivities(activities []Activity) []ActivityFormatter {
	activitiesFormatter := []ActivityFormatter{}

	for _, activity := range activities {
		activityFormatter := FormatActivity(activity)
		activitiesFormatter = append(activitiesFormatter, activityFormatter)
	}

	return activitiesFormatter
}

func FormatActivity(activity Activity) ActivityFormatter {
	activityFormatter := ActivityFormatter{}
	activityFormatter.ID = activity.ID
	activityFormatter.Email = activity.Email
	activityFormatter.Title = activity.Title
	activityFormatter.CreatedAt = activity.CreatedAt
	activityFormatter.UpdatedAt = activity.UpdatedAt

	if activity.DeletedAt == nil {
		activityFormatter.DeletedAt = nil
	}

	return activityFormatter
}
