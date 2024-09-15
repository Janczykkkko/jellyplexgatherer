package jellyplexgatherer

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

// Function to get activity since provided time (current time - x minutes)
func GetJellyActivityLogData(jellyfinAddress, jellyfinApiKey string, minutesSinceNow int, maxRecords int32) (activityLog JellyActivityLog, err error) {
	// Calculate the time since the scrape interval
	timeSince := time.Now().Add(-time.Duration(minutesSinceNow) * time.Minute)
	timeSinceIso := timeSince.Format("2006-01-02T15:04:05")

	// Build the URL with maxRecords
	url := fmt.Sprintf("%s/System/ActivityLog/Entries?minDate=%s&limit=%d&api_key=%s", jellyfinAddress, timeSinceIso, maxRecords, jellyfinApiKey)

	// Make the GET request
	resp, err := http.Get(url)
	if err != nil {
		return JellyActivityLog{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return JellyActivityLog{}, fmt.Errorf("failed to fetch data: %s", resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(&activityLog)
	if err != nil {
		return JellyActivityLog{}, fmt.Errorf("failed to decode response: %v", err)
	}

	return activityLog, nil
}

// GetOnlineUsersFromLog generates a list of currently online users from the activity log
func GetOnlineUsersFromLog(logs JellyActivityLog) JellyOnlineUsers {
	onlineUsersMap := make(map[string]JellyUserStatus) // Map to track latest status for each user and device

	// Iterate in reverse to prioritize the latest events
	for i := len(logs.Items) - 1; i >= 0; i-- {
		item := logs.Items[i]

		// Extract username and device from the Name field (e.g., "jan is online from Chrome")
		parts := strings.Split(item.Name, " from ")
		if len(parts) < 2 {
			continue // Ignore entries that don't follow the expected format
		}

		username := strings.Fields(parts[0])[0] // Extract username
		device := parts[1]                      // Extract device name (after 'from')

		// Create a key for the user's session on this device
		key := fmt.Sprintf("%s:%s", username, device)

		// Update the user's online status based on the event type
		if item.Type == "SessionStarted" {
			onlineUsersMap[key] = JellyUserStatus{
				UserName: username,
				Device:   device,
				Online:   true,
			}
		} else if item.Type == "SessionEnded" {
			// Remove the user's status when a session ends
			delete(onlineUsersMap, key)
		}
	}

	// Convert map to list
	var onlineUsers JellyOnlineUsers
	for _, status := range onlineUsersMap {
		if status.Online {
			onlineUsers = append(onlineUsers, status)
		}
	}

	return onlineUsers
}

// Fetches the activity log data and returns the list of currently online users in struct format
func GetOnlineUsers(jellyfinAddress, jellyfinApiKey string, maxRecords int32, minutesSinceNow int) (JellyOnlineUsers, error) {
	// Fetch activity log data from Jellyfin API
	activityLog, err := GetJellyActivityLogData(jellyfinAddress, jellyfinApiKey, minutesSinceNow, maxRecords)
	if err != nil {
		log.Printf("Error fetching activity log: %v", err)
		return nil, err
	}

	// Generate and return the list of currently online users from the activity log
	return GetOnlineUsersFromLog(activityLog), nil
}
