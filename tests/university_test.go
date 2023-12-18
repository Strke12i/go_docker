package tests

import (
	"bytes"
	"net/http"
	"testing"
)

const url = "http://localhost:8080/api/universities"

var client = http.Client{}

func TestCreateUniversity(t *testing.T) {
	payloads := [][]byte{
		[]byte(`{"name": "Universidade Federal de Santa Maria"}`),
		[]byte(`{"name": "Universidade Federal do Rio Grande do Sul"}`),
		[]byte(`{"name": "Universidade Federal de Santa Catarina"}`),
	}

	for _, payload := range payloads {
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
		if err != nil {
			t.Error(err)
		}

		resp, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}

		if resp.StatusCode != 200 {
			t.Error("University was not created")
		}
	}
}

func TestGetUniversityByID(t *testing.T) {
	req, err := http.NewRequest("GET", url+"/1", nil)
	if err != nil {
		t.Error(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Error("University was not found")
	}
}

func TestGetAllUniversities(t *testing.T) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Error(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Error("Universities were not found")
	}
}

func TestUpdateUniversity(t *testing.T) {
	payload := []byte(`{"name": "Universidade Federal de Santa Maria, RS"}`)

	req, err := http.NewRequest("PUT", url+"/1", bytes.NewBuffer(payload))
	if err != nil {
		t.Error(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Error("University was not updated")
	}
}

func TestDeleteUniversity(t *testing.T) {
	req, err := http.NewRequest("DELETE", url+"/1", nil)
	if err != nil {
		t.Error(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Error("University was not deleted")
	}
}
