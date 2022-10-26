require './constants/commands.rb'

module CLIMode
    def action
        print "enter command: "
        command = gets.chomp

        response = {
            command: command,
            params: { fleetId: 1, x: 1, y: 1 }
        }
        return response
    end
end

module DebugMode
    def action
        response = {
            command: COMMANDS[:ATTACK],
            params: { fleetId: 0, x: 0, y: 0 }
        }
        return response
    end
end

class Player
    attr_reader :id, :name, :fleetList
    def initialize(name)
        @id = -1
        @name = name
        @fleetList = []
    end

    def addFleet(fleet)
        #
    end

    def setPositionForFleet
        #
    end
end

class CLIPlayer < Player
    include CLIMode

    def initialize(name)
        super
        @id = 0
    end

    def addFleet(fleet)
        #@fleetList << {fleet: fleet, x: -1, y: -1}
    end

    def action
        super
        #puts "#{@name} command : #{super}"
    end

    def setPositionForFleet(x, y)
        #
    end
end

class DebugPlayer < Player
    include DebugMode

    def initialize(name)
        super
        @id = 1
    end

    def addFleet(fleet)
        #@fleetList << {fleet: fleet, x: -1, y: -1}
    end

    def action
        super
        #puts "#{@name} command : #{super}"
    end

    def setPositionForFleet(x, y)
        #
    end
end