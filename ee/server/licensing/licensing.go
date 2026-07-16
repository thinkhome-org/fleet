package licensing

import (
	"time"

	"github.com/fleetdm/fleet/v4/server/fleet"
)

// LoadLicense enables Fleet Premium for every build.
func LoadLicense(_ string) (*fleet.LicenseInfo, error) {
	return &fleet.LicenseInfo{
		Tier:                  fleet.TierPremium,
		Organization:          "Local",
		DeviceCount:           1000,
		Expiration:            time.Now().AddDate(10, 0, 0),
		Note:                  "Built-in Fleet Premium license",
		AllowDisableTelemetry: true,
	}, nil
}
