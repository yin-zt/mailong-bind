package model

import (
	"gorm.io/gorm"
	"time"
)

type Dns struct {
	gorm.Model
	DnsRecord   string    `gorm:"type:varchar(100);not null;comment:'域名记录'" json:"dns_record"`                    // 域名记录
	DnsType     uint      `gorm:"type:tinyint(1);default:1;;comment:'域名类型(1A记录， 2CNAME记录)'" json:"dns_type"`      // 域名类型
	DnsTarget   string    `gorm:"type:varchar(100);comment:'解析目标'" json:"dns_target"`                             // 解析目标
	DnsView     string    `gorm:"type:varchar(50);comment:'域名所在视图'" json:"dns_view"`                              // 域名所属视图
	DnsStatus   uint      `gorm:"type:tinyint(1);comment:'域名状态(1待上线，2已上线，3待下线，4，已下线)'" json:"dns_status"`         // 域名状态
	DnsProtocol string    `gorm:"type:varchar(20);not null;default:'TCP';comment:'域名服务所用协议'" json:"dns_protocol"` // 域名服务所用协议
	DnsPort     uint      `gorm:"type:int(5);not null;default:80;comment:'域名服务所用端口'" json:"dns_port"`             // 域名服务所用端口
	DnsSystem   string    `gorm:"type:varchar(255);comment:'域名服务所属系统'" json:"dns_system"`                         // 域名服务所在系统
	DnsCreated  time.Time `gorm:"type:datetime(3);comment:'创建时间'" json:"dns_created"`                             // 域名创建时间
	DnsUpdated  time.Time `gorm:"type:datetime(3);comment:'修改时间'" json:"dns_updated"`                             // 域名修改时间
	DnsDomain   []*View   `gorm:"ForeignKey:domain_name;AssociationForeignKey:id" json:"dns_domain"`              // 域名关联的
}
