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

func TestZoomInOut(t *testing.T) {
	cam := New("10.13.34.8:52381")

	err := cam.ZoomTele(context.Background())
	if err != nil {
		t.Fatalf("failed to zoom tele: %s\n", err)
	}

	time.Sleep(1000 * time.Millisecond)

	err = cam.ZoomWide(context.Background())
	if err != nil {
		t.Fatalf("failed to zoom tele: %s\n", err)
	}

	time.Sleep(1000 * time.Millisecond)

	err = cam.ZoomStop(context.Background())
	if err != nil {
		t.Fatalf("failed to zoom tele: %s\n", err)
	}
}
