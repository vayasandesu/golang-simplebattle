package combat

import (
	//combatCommand "simplebattle/combat/command"
	combat "simplebattle/combat/commands"
	"simplebattle/command"
)

type Arena struct {
	turn       int
	Characters *[]interface{}
}

func (arena *Arena) ExecuteCommand(command command.ICommand) {
	command.Execute()
	for _, character := range *arena.Characters {
		character.(ILogable).GetInformation()
	}
}

func (arena Arena) GetCommand() []command.ICommand {
	var commands []command.ICommand
	var character interface{} = (*arena.Characters)[arena.turn]
	commands = AddAttackCommand(character, commands)
	commands = AddDefenseCommand(character, commands)
	commands = AddRestoreCommand(character, commands)

	return commands
}

func AddAttackCommand(character interface{}, commands []command.ICommand) []command.ICommand {
	attacker, ok := character.(IAttacker)
	if ok {
		command := combat.AttackCommand{
			Attacker: attacker,
			Target:   nil,
		}
		return append(commands, command)
	}

	return commands
}

func AddDefenseCommand(character interface{}, commands []command.ICommand) []command.ICommand {
	component, ok := character.(IDefender)
	if ok {
		command := combat.DefenseCommand{
			Self: component,
		}
		return append(commands, command)
	}

	return commands
}

func AddRestoreCommand(character interface{}, commands []command.ICommand) []command.ICommand {
	component, ok := character.(IRestorer)
	if ok {
		command := combat.RestoreCommand{
			Self: component,
		}
		return append(commands, command)
	}

	return commands
}
