package controllers

import "testing"

func TestAppWithoutRouterInstance(t *testing.T) {
	app := &App{}
	if err := app.Listen(); err == nil {
		t.Error("should return error when there isn't a router instantiated")
	}
}
