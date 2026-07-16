package licensing

import (
	"testing"
	"time"

	"github.com/fleetdm/fleet/v4/server/fleet"
	"github.com/stretchr/testify/require"
)

func TestLoadLicenseAlwaysPremium(t *testing.T) {
	for _, key := range []string{"", "invalid"} {
		license, err := LoadLicense(key)
		require.NoError(t, err)
		require.Equal(t, fleet.TierPremium, license.Tier)
		require.True(t, license.IsPremium())
		require.False(t, license.IsExpired())
		require.Equal(t, "Local", license.Organization)
		require.Equal(t, 1000, license.DeviceCount)
		require.WithinDuration(t, time.Now().AddDate(10, 0, 0), license.Expiration, time.Second)
		require.Equal(t, "Built-in Fleet Premium license", license.Note)
		require.True(t, license.IsAllowDisableTelemetry())
	}
}
