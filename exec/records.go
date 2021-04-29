package exec

import (
	"errors"
	"fmt"
	aliDNS "github.com/alibabacloud-go/alidns-20150109/v2/client"
	"github.com/alibabacloud-go/tea/tea"
)

// GetRecords .
func GetRecords() (*aliDNS.DescribeDomainRecordsResponse, error) {
	describeDomainRecordsRequest := &aliDNS.DescribeDomainRecordsRequest{
		DomainName: tea.String("imsowhy.com"),
	}

	return client.DescribeDomainRecords(describeDomainRecordsRequest)
}

// AddDefaultRecord .
func AddDefaultRecord(domain *string, rr *string, recordType *string, value *string) (*aliDNS.AddDomainRecordResponse, error) {
	if domain == nil || rr == nil || recordType == nil || value == nil {
		return nil, errors.New(fmt.Sprintf("param error, domain: %v, rr: %v, type: %v, value: %v",
			domain, rr, recordType, value))
	}

	addDomainRecord := &aliDNS.AddDomainRecordRequest{
		DomainName: domain,
		RR:         rr,
		Type:       recordType,
		Value:      value,
	}

	return client.AddDomainRecord(addDomainRecord)
}

// UpdateRecord .
func UpdateDefaultRecord(recordID *string, rr *string, recordType *string, value *string) (*aliDNS.UpdateDomainRecordResponse, error) {
	if recordID == nil || rr == nil || recordType == nil || value == nil {
		return nil, errors.New(fmt.Sprintf("param error, record_id: %v, rr: %v, type: %v, value: %v",
			recordID, rr, recordType, value))
	}

	updateDomainRecord := &aliDNS.UpdateDomainRecordRequest{
		RecordId: recordID,
		RR:       rr,
		Type:     recordType,
		Value:    value,
		TTL:      tea.Int64(1),
	}

	return client.UpdateDomainRecord(updateDomainRecord)
}

// DeleteRecord .
func DeleteRecord(recordID *string) (*aliDNS.DeleteDomainRecordResponse, error) {
	if recordID == nil {
		return nil, errors.New(fmt.Sprintf("param error, record_id: %v", recordID))
	}

	deleteRecord := &aliDNS.DeleteDomainRecordRequest{
		RecordId: recordID,
	}

	return client.DeleteDomainRecord(deleteRecord)
}
