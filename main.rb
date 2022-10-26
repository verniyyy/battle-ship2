#!/bin/ruby

require './Player.rb'
require './Fleet.rb'
require './GameMaster.rb'

def main
    cliPlayer = CLIPlayer.new("cli-player")
    cliPlayer.addFleet(Fleet.new("cli1", 0))
    cliPlayer.addFleet(Fleet.new("cli2", 1))
    cliPlayer.addFleet(Fleet.new("cli3", 2))

    debugPlayer = DebugPlayer.new("debug-player")
    debugPlayer.addFleet(Fleet.new("dbg1", 0))
    debugPlayer.addFleet(Fleet.new("dbg2", 1))
    debugPlayer.addFleet(Fleet.new("dbg3", 2))

    gameMaster = GameMaster.new(cliPlayer, debugPlayer)
    gameMaster.game
end

def getAny
    return 3, "aaa"
end



module Fuga
    def hello
        puts "hello from module"
    end
end

class Hoge
    include Fuga
    def hello
        super
        puts "hello from method"
    end
end

#hoge = Hoge.new
#hoge.hello

main