package userconfig

import (
	"zakirullin/stuffbot/internal/consts"
	"zakirullin/stuffbot/pkg/tg"
)

var AvailableQuickBtns = []tg.Btn{
	tg.NewBtn("Later", tg.NewCmd(consts.CmdLater, nil)),
	tg.NewBtn("Search", tg.NewCustomCmd(consts.CmdInlineQuerySearchEveryWhere, nil, tg.CmdTypeInlineQueryCurrentChat)),
	tg.NewBtn("Files", tg.NewCmd(consts.CmdShowFiles, nil)),
	tg.NewBtn("Checklists", tg.NewCmd(consts.CmdShowChecklists, nil)),
	tg.NewBtn("Postpone", tg.NewCmd(consts.CmdShowPostpone, nil)),
	tg.NewBtn("Read", tg.NewCmd(consts.CmdShowReadChecklist, nil)),
	tg.NewBtn("Watch", tg.NewCmd(consts.CmdShowWatchChecklist, nil)),
	tg.NewBtn("Shop", tg.NewCmd(consts.CmdShowShopChecklist, nil)),
	tg.NewBtn("Habits", tg.NewCustomCmd(consts.CmdWebAppHabits, nil, tg.CmdTypeWebApp)),
}

func (c *Config) AddQuickCmd(cmd string) bool {
	// Does this cmd already exist?
	for _, existingCmd := range c.raw.QuickCmds {
		if existingCmd == cmd {
			return false
		}
	}
	c.raw.QuickCmds = append(c.raw.QuickCmds, cmd)
	return true
}

func (c *Config) QuickCmds() []string {
	return c.raw.QuickCmds
}

func (c *Config) HasQuickCmd(cmd string) bool {
	for _, pref := range c.raw.QuickCmds {
		if cmd == pref {
			return true
		}
	}
	return false
}

func (c *Config) DelQuickCmd(cmd string) bool {
	var newButtons []string
	found := false // Was the target
	for _, curQuickCmd := range c.raw.QuickCmds {
		if curQuickCmd == cmd {
			found = true
		} else {
			newButtons = append(newButtons, curQuickCmd)
		}
	}
	c.raw.QuickCmds = newButtons
	return found
}
