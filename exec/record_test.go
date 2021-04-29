package exec

import (
	"github.com/alibabacloud-go/tea/tea"
	"github.com/sowhyim/alidomain/config"
	"github.com/sowhyim/alidomain/define"
	"testing"
)

func init() {
	// 此类接口均依赖阿里服务器，故不提供 key，见谅
	config.Init("../conf.d/access_key.yaml")
	Init()
}

func TestAddAndDeleteRecord2(t *testing.T) {
	ip := "113.110.201.204"
	domain := "imsowhy.com"
	rr := "test"
	res, err := AddDefaultRecord(tea.String(domain), tea.String(rr), tea.String(define.RecordTypeA), tea.String(ip))
	if err != nil {
		t.Fatalf("add default record failed, err: %v", err)
	}
	t.Logf("add record success!! res: %+v", res)

	records, err := GetRecords()
	if err != nil {
		t.Fatalf("add default record failed, err: %v", err)
	}
	t.Logf("got record success!! res: %+v", records)

	for i := range records.Body.DomainRecords.Record {
		if *records.Body.DomainRecords.Record[i].RR == rr && *records.Body.DomainRecords.Record[i].DomainName == domain {
			res, err := DeleteRecord(records.Body.DomainRecords.Record[i].RecordId)
			if err != nil {
				t.Fatalf("add default record failed, err: %v", err)
			}
			t.Logf("delete record success!! res: %+v", res)
		}
	}

	records, err = GetRecords()
	if err != nil {
		t.Fatalf("add default record failed, err: %v", err)
	}
	t.Logf("got record success!! final res: %+v", records)
}

func TestGetRecords(t *testing.T) {
	res, err := GetRecords()
	if err != nil {
		t.Fatalf("add default record failed, err: %v", err)
	}
	t.Logf("success!! res: %+v", res)
}

func TestUpdateDefaultRecord(t *testing.T) {
	rr := "host"
	domain := "imsowhy.com"
	records, err := GetRecords()
	if err != nil {
		t.Fatalf("add default record failed, err: %v", err)
	}
	t.Logf("got record success!! res: %+v", records)

	for i := range records.Body.DomainRecords.Record {
		if *records.Body.DomainRecords.Record[i].RR == rr && *records.Body.DomainRecords.Record[i].DomainName == domain {
			res, err := UpdateDefaultRecord(records.Body.DomainRecords.Record[i].RecordId, tea.String(rr), records.Body.DomainRecords.Record[i].Type, tea.String("113.110.200.41"))
			if err != nil {
				t.Fatalf("add default record failed, err: %v", err)
			}
			t.Logf("delete record success!! res: %+v", res)
		}
	}
}
