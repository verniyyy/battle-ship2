require 'securerandom'

class Fleet

    attr_reader :id, :name, :type, :hp, :bullet

    def initialize(name, type)
        @id = SecureRandom.uuid
        @name = name
        @type = type
        @hp = 100
        @atk = 40
        @bullets = 10
    end

    def getDamage(damage)
        #
    end

    def attack()
        #
    end

    def destroyed?
        #
    end
end