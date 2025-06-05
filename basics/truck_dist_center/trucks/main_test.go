package main

import (
	"context"
	"testing"
)

func TestMain(t *testing.T) {
	t.Run("processTruck", func(t *testing.T) {
		t.Run("should load and unload a truck cargo", func(t *testing.T) {
			nt := &NormalTruck{id: "1"}
			et := &EletricTruck{id: "2"}
			ctx := context.Background()

			err := processTruck(ctx, nt)
			if err != nil {
				t.Fatalf("Error processing truck, message:\n %s", err)
			}

			err = processTruck(ctx, et)
			if err != nil {
				t.Fatalf("Error processing truck, message:\n %s", err)
			}

			if nt.cargo != 0 {
				t.Fatal("Normal truck cargo should be 0")
			}

			if et.cargo != 0 {
				t.Fatal("Eletrick truck cargo should be 0")
			}

			if et.battery != -2 {
				t.Fatal("Eletrick truck battery should be -2")
			}
		})
	})
}
