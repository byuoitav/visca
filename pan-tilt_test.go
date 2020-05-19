package visca

import (
	"context"
	"testing"
	"time"
)

func TestDirections(t *testing.T) {
	cam := New("10.13.34.8:52381")

	err := cam.TiltUp(context.Background(), 0x0e)
	if err != nil {
		t.Fatalf("failed to pan tilt up: %s\n", err)
	}

	time.Sleep(1500 * time.Millisecond)

	err = cam.PanRight(context.Background(), 0x0b)
	if err != nil {
		t.Fatalf("failed to pan right up: %s\n", err)
	}

	time.Sleep(1500 * time.Millisecond)

	err = cam.TiltDown(context.Background(), 0x0e)
	if err != nil {
		t.Fatalf("failed to pan right down: %s\n", err)
	}

	time.Sleep(1500 * time.Millisecond)

	err = cam.PanLeft(context.Background(), 0x0b)
	if err != nil {
		t.Fatalf("failed to pan right left: %s\n", err)
	}

	time.Sleep(1500 * time.Millisecond)

	err = cam.PanTiltStop(context.Background())
	if err != nil {
		t.Fatalf("failed to pan stop: %s\n", err)
	}
}
