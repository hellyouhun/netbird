package android

import (
	"path/filepath"
	"testing"

	"github.com/netbirdio/netbird/client/internal/profilemanager"
)

func TestPreferences_DefaultValues(t *testing.T) {
	cfgFile := filepath.Join(t.TempDir(), "netbird.json")
	p := NewPreferences(cfgFile)
	defaultVar, err := p.GetAdminURL()
	if err != nil {
		t.Fatalf("读取默认值失败: %s", err)
	}

	if defaultVar != profilemanager.DefaultAdminURL {
		t.Errorf("无效的管理员 URL 默认值: %s", defaultVar)
	}

	defaultVar, err = p.GetManagementURL()
	if err != nil {
		t.Fatalf("读取默认管理 URL 失败: %s", err)
	}

	if defaultVar != profilemanager.DefaultManagementURL {
		t.Errorf("无效的管理 URL 默认值: %s", defaultVar)
	}

	var preSharedKey string
	preSharedKey, err = p.GetPreSharedKey()
	if err != nil {
		t.Fatalf("读取默认预共享密钥失败: %s", err)
	}

	if preSharedKey != "" {
		t.Errorf("无效的预共享密钥: %s", preSharedKey)
	}
}

func TestPreferences_ReadUncommitedValues(t *testing.T) {
	exampleString := "exampleString"
	cfgFile := filepath.Join(t.TempDir(), "netbird.json")
	p := NewPreferences(cfgFile)

	p.SetAdminURL(exampleString)
	resp, err := p.GetAdminURL()
	if err != nil {
		t.Fatalf("读取管理员 URL 失败: %s", err)
	}

	if resp != exampleString {
		t.Errorf("意外的管理员 URL: %s", resp)
	}

	p.SetManagementURL(exampleString)
	resp, err = p.GetManagementURL()
	if err != nil {
		t.Fatalf("读取管理 URL 失败: %s", err)
	}

	if resp != exampleString {
		t.Errorf("意外的管理 URL: %s", resp)
	}

	p.SetPreSharedKey(exampleString)
	resp, err = p.GetPreSharedKey()
	if err != nil {
		t.Fatalf("读取预共享密钥失败: %s", err)
	}

	if resp != exampleString {
		t.Errorf("意外的预共享密钥: %s", resp)
	}
}

func TestPreferences_Commit(t *testing.T) {
	exampleURL := "https://myurl.com:443"
	examplePresharedKey := "topsecret"
	cfgFile := filepath.Join(t.TempDir(), "netbird.json")
	p := NewPreferences(cfgFile)

	p.SetAdminURL(exampleURL)
	p.SetManagementURL(exampleURL)
	p.SetPreSharedKey(examplePresharedKey)

	err := p.Commit()
	if err != nil {
		t.Fatalf("failed to save changes: %s", err)
	}

	p = NewPreferences(cfgFile)
	resp, err := p.GetAdminURL()
	if err != nil {
		t.Fatalf("failed to read admin url: %s", err)
	}

	if resp != exampleURL {
		t.Errorf("unexpected admin url: %s", resp)
	}

	resp, err = p.GetManagementURL()
	if err != nil {
		t.Fatalf("failed to read management url: %s", err)
	}

	if resp != exampleURL {
		t.Errorf("unexpected management url: %s", resp)
	}

	resp, err = p.GetPreSharedKey()
	if err != nil {
		t.Fatalf("failed to read preshared key: %s", err)
	}

	if resp != examplePresharedKey {
		t.Errorf("unexpected preshared key: %s", resp)
	}
}
