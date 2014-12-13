package front

import (
	"log"
	"net/http"
	"testing"
)

func TestPaginationAtZero(t *testing.T) {
	req, reqErr := http.NewRequest("GET", "/", nil)
	if reqErr != nil {
		t.Error(reqErr)
	}
	total := 26
	perPage := 5

	paginator := GetPagination(req, total, perPage)
	log.Println("Offset:", paginator.Offset*paginator.Count, "Count:", paginator.Count)
	if paginator.HasNewer {
		t.Error("Expected to not have newer.")
	}
	if !paginator.HasOlder {
		t.Error("Expected to have older.")
	}
	if paginator.NewerIndex != 0 {
		t.Errorf("New Index Expected: %d got %d", 0, paginator.NewerIndex)
	}
	if paginator.OlderIndex != 2 {
		t.Errorf("Older Index Expected: %d got %d", 2, paginator.OlderIndex)
	}
	if paginator.Offset != 0 {
		t.Errorf("Offset Expected: %d got %d", 0, paginator.Offset)
	}
}

func TestPaginationAtOne(t *testing.T) {
	req, reqErr := http.NewRequest("GET", "/?page=1", nil)
	if reqErr != nil {
		t.Error(reqErr)
	}
	total := 26
	perPage := 5

	paginator := GetPagination(req, total, perPage)
	log.Println("Offset:", paginator.Offset*paginator.Count, "Count:", paginator.Count)
	if paginator.HasNewer {
		t.Error("Expected to not have newer.")
	}
	if !paginator.HasOlder {
		t.Error("Expected to have older.")
	}
	if paginator.NewerIndex != 0 {
		t.Errorf("New Index Expected: %d got %d", 0, paginator.NewerIndex)
	}
	if paginator.OlderIndex != 2 {
		t.Errorf("Older Index Expected: %d got %d", 2, paginator.OlderIndex)
	}
	if paginator.Offset != 0 {
		t.Errorf("Offset Expected: %d got %d", 0, paginator.Offset)
	}
}

func TestPaginationAtTwo(t *testing.T) {
	req, reqErr := http.NewRequest("GET", "/?page=2", nil)
	if reqErr != nil {
		t.Error(reqErr)
	}
	total := 26
	perPage := 5

	paginator := GetPagination(req, total, perPage)
	log.Println("Offset:", paginator.Offset*paginator.Count, "Count:", paginator.Count)
	if !paginator.HasNewer {
		t.Error("Expected to have newer.")
	}
	if !paginator.HasOlder {
		t.Error("Expected to have older.")
	}
	if paginator.NewerIndex != 1 {
		t.Errorf("Newer Index Expected: %d got %d", 1, paginator.NewerIndex)
	}
	if paginator.OlderIndex != 3 {
		t.Errorf("Older Index Expected: %d got %d", 3, paginator.OlderIndex)
	}
	if paginator.Offset != 1 {
		t.Errorf("Offset Expected: %d got %d", 1, paginator.Offset)
	}
}

func TestPaginationAtThree(t *testing.T) {
	req, reqErr := http.NewRequest("GET", "/?page=3", nil)
	if reqErr != nil {
		t.Error(reqErr)
	}
	total := 26
	perPage := 5

	paginator := GetPagination(req, total, perPage)
	log.Println("Offset:", paginator.Offset*paginator.Count, "Count:", paginator.Count)
	if !paginator.HasNewer {
		t.Error("Expected to have newer.")
	}
	if !paginator.HasOlder {
		t.Error("Expected to have older.")
	}
	if paginator.NewerIndex != 2 {
		t.Errorf("Newer Index Expected: %d got %d", 2, paginator.NewerIndex)
	}
	if paginator.OlderIndex != 4 {
		t.Errorf("Older Index Expected: %d got %d", 4, paginator.OlderIndex)
	}
	if paginator.Offset != 2 {
		t.Errorf("Offset Expected: %d got %d", 2, paginator.Offset)
	}
}
