package visca

import (
	"context"
	"testing"
	"time"
)

func TestZoomTele(t *testing.T) {
	cam := New("10.13.34.8:52381")

	err := cam.ZoomTele(context.Background())
	if err != nil {
		t.Fatalf("failed to zoom tele: %s\n", err)
	}
}

func TestZoomWide(t *testing.T) {
	cam := New("10.13.34.8:52381")

	err := cam.ZoomWide(context.Background())
	if err != nil {
		t.Fatalf("failed to zoom wide: %s\n", err)
	}
}

func TestZoomWideTele(t *testing.T) {
	cam := New("10.13.34.8:52381")

	for i := 0; i < 10; i++ {
		err := cam.ZoomTele(context.Background())
		if err != nil {
			t.Fatalf("failed to zoom tele: %s\n", err)
		}

		time.Sleep(time.Duration(i) * time.Second)

		err = cam.ZoomWide(context.Background())
		if err != nil {
			t.Fatalf("failed to zoom wide: %s\n", err)
		}

		time.Sleep(time.Duration(10-i) * time.Second)
	}
}
