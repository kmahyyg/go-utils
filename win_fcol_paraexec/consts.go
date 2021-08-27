// Copyright (C) 2021 kmahyyg
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

var exec_collect_cmds = []string{
	"net user",
	"query user",
	"systeminfo",
	"arp -a",
	"ipconfig /all",
	"netstat -ano",
	"route print",
	"powershell -w hidden -ep bypass Get-WmiObject -Namespace root\\SecurityCenter2 -Class AntiVirusProduct",
	"wmic qfe get hotfixid,Description",
	"wmic product get name,version",
	"wmic startup list full",
	"REG QUERY \"HKLM\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\App Paths\"",
	"REG QUERT \"HKEY_CURRENT_USER\\Software\\Microsoft\\Terminal Server Client\\Servers\" /s",
	"at",
	"schtasks /query",
	"schtasks /query /fo list /v",
	"tasklist /v /fo table",
	"tasklist /svc",
}

var exec_collect_cmds_comment = []string{
	"用户账户信息", "当前登录会话信息", "系统信息", "ARP缓存", "网络配置", "当前网络连接情况",
	"当前网络路由表", "反病毒软件安装", "通过MSI安装的软件信息", "系统补丁情况", "自启动情况", "注册表App Path项", "MSTSC记录的RDP连接", "计划任务(老)",
	"计划任务(新)", "计划任务详情(新)", "进程列表(详)", "进程列表(含服务)",
}

type CmdOutput struct {
	CmdComment string
	CmdDetail  string
	CmdOutput  string
}

type finalResults struct {
	CmdOutputs []CmdOutput
}
