package server

import (
	"testing"

	"github.com/HackYourCareer/SmartKickers/internal/config"
	"github.com/stretchr/testify/assert"
)

func Test_isValidTeamID(t *testing.T) {
	tests := []struct {
		name     string
		teamID   int
		expected bool
	}{
		// TODO: Add test cases.
		{
			name:     "Team white ID, should return true",
			teamID:   config.TeamWhite,
			expected: true,
		},
		{
			name:     "Team blue ID, should return true",
			teamID:   config.TeamBlue,
			expected: true,
		},
		{
			name:     "Incorrect ID, should return false",
			teamID:   -1,
			expected: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, isValidTeamID(tt.teamID))
		})
	}
}
