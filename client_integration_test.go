package wireguardctrl_test

import (
	"os"
	"testing"

	"github.com/mdlayher/wireguardctrl"
)

func TestClientIntegration(t *testing.T) {
	c, err := wireguardctrl.New()
	if err != nil {
		if os.IsNotExist(err) {
			t.Skip("skipping, wireguardctrl is not available on this system")
		}

		t.Fatalf("failed to open client: %v", err)
	}
	defer c.Close()

	// TODO(mdlayher): expand upon this.

	t.Run("devices", func(t *testing.T) {
		devices, err := c.Devices()
		if err != nil {
			if os.IsPermission(err) {
				t.Skip("skipping, wireguardctrl may require elevated privileges")
			}

			t.Fatalf("failed to get devices: %v", err)
		}

		for _, d := range devices {
			t.Logf("device: %s: %s", d.Name, d.PublicKey.String())
		}
	})
}