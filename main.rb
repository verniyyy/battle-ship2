#!/bin/ruby

require './Player.rb'
require './GameMaster.rb'

def main
    cliPlayer = CLIPlayer.new('000001', "cli-player")
    debugPlayer = DebugPlayer.new('000002', "debug-player")


    gameMaster = GameMaster.new(cliPlayer, debugPlayer)
    gameMaster.game
end


main