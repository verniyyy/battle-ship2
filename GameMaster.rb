require './constants/commands.rb'
require './Map.rb'

class GameMaster
    def initialize(*players)
        @turn = 0
        @map = Map.new
        @playerList = players
        

        @playerList.each do |player|
            player.fleetList.each do |fleet|
                @map.setObject(fleet)
            end
        end
    end

    def game
        while !gameEnd? do
            turn
        end
        gameEndProcess
    end

    def turn
        @turn += 1

        puts "[TURN #{@turn}]"

        @playerList.each do |player|
            err = true
            while err do
                err = action(player)
            end
        end
    end

    def action(player)
        response = player.action
        command, params = parseResponseForAction(response)
        viewAction(command, params)

        case command
        when COMMANDS[:ATTACK]
            attack(player, params)

        when COMMANDS[:MOVE]
            move(player, params)

        else
            return "invalid command"
        end
        
        err = nil
        return err
    end

    def viewAction(command, params)
        puts "#{player.name}\##{player.id}:
            \t command: #{command}
            \t params: #{params}"
    end

    def parseResponseForAction(response)
        return response[:command], response[:params]
    end

    def move(player, params)
        fleet = player.fleetList[params[:fleetID]]
        @map.updateInfo(fleet, params[:x], params[:y])
    end

    def attack(player, params)
        attacker = player.fleetList[params[:fleetID]]

        #@map.move(fleet, params[:x], params[:y])
    end

    def gameEnd?
        @turn>=3
    end

    def gameEndProcess
        #
    end
end
