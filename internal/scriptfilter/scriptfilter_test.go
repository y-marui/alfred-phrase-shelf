package scriptfilter

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestResponseWriteOmitsAbsentFields(t *testing.T) {
	resp := Response{Items: []Item{{Title: "Hello"}}}

	var buf bytes.Buffer
	if err := resp.Write(&buf); err != nil {
		t.Fatalf("Write: %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(buf.Bytes(), &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}

	items, ok := decoded["items"].([]any)
	if !ok || len(items) != 1 {
		t.Fatalf("items = %v, want 1 item", decoded["items"])
	}
	item, ok := items[0].(map[string]any)
	if !ok {
		t.Fatalf("item = %v, want object", items[0])
	}
	if item["title"] != "Hello" {
		t.Errorf("title = %v, want %q", item["title"], "Hello")
	}
	for _, absent := range []string{"uid", "subtitle", "arg", "valid", "autocomplete", "icon", "variables", "mods"} {
		if _, present := item[absent]; present {
			t.Errorf("field %q should be omitted when unset, got %v", absent, item[absent])
		}
	}
	for _, absent := range []string{"variables", "skipknowledge"} {
		if _, present := decoded[absent]; present {
			t.Errorf("top-level field %q should be omitted when unset, got %v", absent, decoded[absent])
		}
	}
}

func TestResponseWriteIncludesExplicitFalse(t *testing.T) {
	resp := Response{Items: []Item{{Title: "Info", Valid: BoolPtr(false)}}}

	var buf bytes.Buffer
	if err := resp.Write(&buf); err != nil {
		t.Fatalf("Write: %v", err)
	}

	var decoded struct {
		Items []struct {
			Valid *bool `json:"valid"`
		} `json:"items"`
	}
	if err := json.Unmarshal(buf.Bytes(), &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if len(decoded.Items) != 1 || decoded.Items[0].Valid == nil || *decoded.Items[0].Valid {
		t.Errorf("decoded valid = %v, want explicit false", decoded.Items[0].Valid)
	}
}

func TestResponseWriteIncludesVariablesIconAndMods(t *testing.T) {
	resp := Response{Items: []Item{{
		Title:        "X",
		Autocomplete: "x ",
		Icon:         &Icon{Path: "icons/x.png"},
		Variables:    map[string]string{"k": "v"},
		Mods: map[string]Mod{
			"cmd": {Arg: "cmd-arg", Subtitle: "Held cmd"},
		},
	}}}

	var buf bytes.Buffer
	if err := resp.Write(&buf); err != nil {
		t.Fatalf("Write: %v", err)
	}

	var decoded struct {
		Items []struct {
			Autocomplete string                `json:"autocomplete"`
			Icon         struct{ Path string } `json:"icon"`
			Variables    map[string]string     `json:"variables"`
			Mods         map[string]struct {
				Arg      string `json:"arg"`
				Subtitle string `json:"subtitle"`
			} `json:"mods"`
		} `json:"items"`
	}
	if err := json.Unmarshal(buf.Bytes(), &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	got := decoded.Items[0]
	if got.Autocomplete != "x " {
		t.Errorf("autocomplete = %q, want %q", got.Autocomplete, "x ")
	}
	if got.Icon.Path != "icons/x.png" {
		t.Errorf("icon.path = %q, want %q", got.Icon.Path, "icons/x.png")
	}
	if got.Variables["k"] != "v" {
		t.Errorf("variables[k] = %q, want %q", got.Variables["k"], "v")
	}
	if got.Mods["cmd"].Arg != "cmd-arg" || got.Mods["cmd"].Subtitle != "Held cmd" {
		t.Errorf("mods[cmd] = %+v, want arg=cmd-arg subtitle=%q", got.Mods["cmd"], "Held cmd")
	}
}

func TestResponseWriteSkipKnowledge(t *testing.T) {
	resp := Response{Items: []Item{{Title: "X"}}, SkipKnowledge: true}

	var buf bytes.Buffer
	if err := resp.Write(&buf); err != nil {
		t.Fatalf("Write: %v", err)
	}

	var decoded struct {
		SkipKnowledge bool `json:"skipknowledge"`
	}
	if err := json.Unmarshal(buf.Bytes(), &decoded); err != nil {
		t.Fatalf("Unmarshal: %v", err)
	}
	if !decoded.SkipKnowledge {
		t.Error("skipknowledge = false, want true")
	}
}
