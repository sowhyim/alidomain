package main

import (
	aliDNS "github.com/alibabacloud-go/alidns-20150109/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/sowhyim/alidomain/config"
	"log"
	"os"
)

/**
 * 使用AK&SK初始化账号Client
 * @param accessKeyId
 * @param accessKeySecret
 * @return Client
 * @throws Exception
 */
func CreateClient() (_result *aliDNS.Client, _err error) {
	c := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: tea.String(config.AccessKey.AccessKeyID),
		// 您的AccessKey Secret
		AccessKeySecret: tea.String(config.AccessKey.AccessKeySecret),
	}
	// 访问的域名
	c.Endpoint = tea.String("dns.aliyuncs.com")
	_result = &aliDNS.Client{}
	_result, _err = aliDNS.NewClient(c)
	return _result, _err
}

func _main(args []*string) (err error) {
	client, err := CreateClient()
	if err != nil {
		return err
	}

	describeDomainRecordsRequest := &aliDNS.DescribeDomainRecordsRequest{
		DomainName: tea.String("imsowhy.com"),
	}
	// 复制代码运行请自行打印 API 的返回值
	res, err := client.DescribeDomainRecords(describeDomainRecordsRequest)
	if err != nil {
		return err
	}
	log.Printf("%+v", res)
	return err
}

func main() {
	config.Init("")
	err := _main(tea.StringSlice(os.Args[1:]))
	if err != nil {
		panic(err)
	}
}
