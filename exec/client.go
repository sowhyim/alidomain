package exec

import (
	aliDNS "github.com/alibabacloud-go/alidns-20150109/v2/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/sowhyim/alidomain/config"
)

// TODO: 改写数据类型，可支持多客户端并存
var client *aliDNS.Client

// Init .
func Init() {
	var err error
	c := &openapi.Config{
		// 您的AccessKey ID
		AccessKeyId: tea.String(config.AccessKey.AccessKeyID),
		// 您的AccessKey Secret
		AccessKeySecret: tea.String(config.AccessKey.AccessKeySecret),
	}
	// 访问的域名
	c.Endpoint = tea.String("dns.aliyuncs.com")
	client, err = aliDNS.NewClient(c)
	if err != nil {
		panic(err)
	}
}
