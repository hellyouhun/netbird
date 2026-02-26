package android

import "testing"

func TestDNSList_Get(t *testing.T) {
	l := DNSList{}

	// 添加有效的 DNS 地址
	err := l.Add("8.8.8.8")
	if err != nil {
		t.Errorf("意外的错误: %s", err)
	}

	// 测试获取有效索引
	addr, err := l.Get(0)
	if err != nil {
		t.Errorf("无效的错误: %s", err)
	}
	if addr != "8.8.8.8" {
		t.Errorf("期望 8.8.8.8，得到 %s", addr)
	}

	// 测试负索引
	_, err = l.Get(-1)
	if err == nil {
		t.Errorf("期望错误但得到 nil")
	}

	// 测试超出范围索引
	_, err = l.Get(1)
	if err == nil {
		t.Errorf("期望错误但得到 nil")
	}
}
