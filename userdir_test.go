package userdir

import "testing"

func TestHome(t *testing.T) {
  if Home() == "" {
    t.Error("No Home Directory Found.")
  }
}

func TestAuthorizedKeys(t *testing.T) {
  if AuthorizedKeys() == "" {
    t.Error("No Home Directory Found.")
  }
}

