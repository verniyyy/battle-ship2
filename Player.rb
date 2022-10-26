require 'securerandom'
require './constants/commands.rb'


class Player
    attr_reader :userId, :name, :fleet
    def initialize(userId, name)
        @userId = userId
        @name = name
        @fleet = []
    end

    def addWarship(warship)
        #
    end

    def setPositionForWarship
        #
    end
end

class CLIPlayer < Player
    def initialize(userId, name)
        super
        @fleet = [
            { id: SecureRandom.uuid, flagShip: true,  position: {x: 0, y: 0} },
            { id: SecureRandom.uuid, flagShip: false, position: {x: 1, y: 0} },
            { id: SecureRandom.uuid, flagShip: false, position: {x: 1, y: 1} }
        ]
    end

    def getVirtualRequest(api)
        case api
        when "warship/detail"
        when "warship/list"
        when "game/inqueue"
            response = gameInqueue
        when "game/ready"
            response = gameReady
        when "game/action"
            response = gameAction
        else
        end
        
        return response
    end

    def gameInqueue
        return {
            userId: @userId,
            fleet: [
                { id: SecureRandom.uuid, flagShip: true},
                { id: SecureRandom.uuid, flagShip: false},
                { id: SecureRandom.uuid, flagShip: false},
            ]
        }
    end

    def gameReady
        return {
            userId: @userId,
            fleet: @fleet
        }
    end

    def gameAction
        print "enter command: "
        command = gets.chomp

        return {
            command: command,
            params: { id: @fleet[0][:id], x: 1, y: 1 }
        }
    end
end

class DebugPlayer < CLIPlayer
    def action
        response = {
            userId: @userId,
            command: COMMANDS[:ATTACK],
            params: { warshipId: 0, x: 0, y: 0 }
        }
        return response
    end
end
