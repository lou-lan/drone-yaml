package yaml

import "testing"

func TestCronUnmarshal(t *testing.T) {
	diff, err := diff("testdata/cron.yml")
	if err != nil {
		t.Error(err)
	}
	if diff != "" {
		t.Error("Failed to parse cron")
		t.Log(diff)
	}
}

func TestCronValidate(t *testing.T) {
	cron := new(Cron)
	cron.Spec.Branch = "master"
	if err := cron.Validate(); err != nil {
		t.Error(err)
		return
	}

	cron.Spec.Branch = ""
	if err := cron.Validate(); err == nil {
		t.Errorf("Expect invalid cron error")
	}
}
