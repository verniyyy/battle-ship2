require_relative './Warship.rb'

class WarshipObjects::Battleship
    def initialize(id, userId, warshipId, individualId, name)
        super
        @type       = "Battleship"
        @hp         = 100
        @atk        = 40
        @resource   = 50
    end

    def attack(command)
        case command
        when false
        else
            return defaultAttack
        end        
    end

    def defaultAttack()
        @resource -= 5
        err = false
        return err
    end
end
