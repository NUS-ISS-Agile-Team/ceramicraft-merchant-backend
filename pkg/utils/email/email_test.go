package email

import (
	"os"
	"path/filepath"
	"testing"
)

func TestEmailSender_Send(t *testing.T) {
	// 设置正确的配置文件路径
	// 获取项目根目录
	currentDir, _ := os.Getwd()
	rootDir := filepath.Dir(filepath.Dir(filepath.Dir(currentDir)))
	
	// 手动设置配置文件路径
	os.Chdir(rootDir) // 切换到项目根目录
	
	// 测试代码
	sender := NewEmailSender()
	emailTo := "e1554065@u.nus.edu"
	subject := "陶瓷商城商家端测试邮件"
	content := "<h1>这是一封测试邮件</h1>"

	err := sender.Send(content, emailTo, subject)
	if err != nil {
		t.Errorf("发送邮件失败: %v", err)
	} else {
		t.Logf("邮件已成功发送到 %s", emailTo)
	}
}
