require_relative './constants/commands.rb'
require_relative './Map.rb'
require_relative './infrastructure/repository/Warship.rb'

class GameMaster
    def initialize(*players)
        @turn = 0
        @map = Map.new
        @playerList = players
        

        @playerList.each do |player|
            player.fleet.each do |warshipInfo|
                warship = WarshipRepository.get(player.userId, warshipInfo[:id])
                @map.setObject(warship, warshipInfo[:position][:x], warshipInfo[:position][:y])
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
        response = player.getVirtualRequest("game/action")
        command, params = parseResponseForAction(response)
        viewAction(player, command, params)

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

    def viewAction(player, command, params)
        puts "#{player.name}\##{player.userId}:
            \t command: #{command}
            \t params: #{params}"
    end

    def parseResponseForAction(response)
        return response[:command], response[:params]
    end

    def move(player, params)
        warship = player.warshipList[params[:warshipID]]
        @map.updateInfo(warship, params[:x], params[:y])
    end

    def attack(player, params)
        attacker = player.warshipList[params[:warshipID]]

        #@map.move(warship, params[:x], params[:y])
    end

    def gameEnd?
        @turn>=3
    end

    def gameEndProcess
        #
    end
end
