package orgcreditcode

import "testing"

func TestGenOrgCheckCode(t *testing.T) {
	var tests = []struct{
		input string
		want  string
	}{
		{"34300329","343003297"},
		{"19220382","192203821"},
		{"70577159","705771593"},
		{"70309976","703099764"},
		{"68221916","682219164"},
		{"68221916","682219164"},
		{"66300312","663003122"},
		{"72719571","727195718"},
		{"MA4UMYGN","MA4UMYGN4"},
	}

	for _, test := range tests{
		if got , err:= GenOrgCheckCode(test.input); got != test.want || err != nil {
			t.Errorf("GenOrgCheckCode(%q) = %v, err = %s", test.input, got, err)
		}
	}
}

func TestGenCreditCheckCode(t *testing.T) {
	var tests = []struct{
		input string
		want  string
	}{
		{"91510100343003297","91510100343003297B"},
		{"91440300192203821","914403001922038216"},
		{"91360100705771593","91360100705771593J"},
		{"91310000703099764","91310000703099764Y"},
		{"91310115682219164","913101156822191640"},
		{"91510100663003122","915101006630031225"},
		{"91330108727195718","91330108727195718R"},
		{"91131001663670945","911310016636709451"},
		{"91230103MA18WL9A7","91230103MA18WL9A78"},
		{"91441900MA4UMYGN4","91441900MA4UMYGN4L"},
		{"91420102MA4KYP7E5","91420102MA4KYP7E5R"},
	}

	for _, test := range tests{
		if got , err:= GenCreditCheckCode(test.input); got != test.want || err != nil {
			t.Errorf("GenCreditCheckCode(%q) = %v, err = %s", test.input, test.want, err)
		}
	}
}

func TestGenCreditCode(t *testing.T) {
	var tests = []struct{
		manageType string
		orgType string
		areaCode string
		orgCode string
		want  string
	}{
		{"9","1","510100","34300329","91510100343003297B"},
		{"9","1","440300","19220382","914403001922038216"},
		{"9","1","360100","70577159","91360100705771593J"},
		{"9","1","310000","70309976","91310000703099764Y"},
		{"9","1","310115","68221916","913101156822191640"},
		{"9","1","510100","66300312","915101006630031225"},
		{"9","1","330108","72719571","91330108727195718R"},
		{"9","1","131001","66367094","911310016636709451"},
		{"9","1","230103","MA18WL9A","91230103MA18WL9A78"},
		{"9","1","441900","MA4UMYGN","91441900MA4UMYGN4L"},
	}

	for _, test := range tests{
		if got , err:= GenCreditCode(test.manageType,test.orgType,test.areaCode,test.orgCode); got != test.want || err != nil {
			t.Errorf("GenCreditCheckCode(%q,%q,%q,%q) = %v, err = %s", test.manageType, test.orgType, test.areaCode, test.orgCode,
				test.want,err)
		}
	}
}