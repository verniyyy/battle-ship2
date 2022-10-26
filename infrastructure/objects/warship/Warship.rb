require 'securerandom'

### Warshipを継承するクラスすべてはダックタイピングによって行動する
module WarshipObjects
    class Warship

        attr_reader :id, :userId, :warshipId, :individualId, :name, :type, :hp, :atk, :resource

        def self.Load(id, userId, warshipId, individualId, name, hp, atk, resource)
            return self.New(id, userId, warshipId, individualId, name).load(hp, atk, resource)
        end

        def initialize(id, userId, warshipId, individualId, name)
            @id             = id
            @userId         = userId
            @warshipId      = warshipId
            @individualId   = individualId
            @name           = name
            @type           = ""
            @hp             = -1
            @atk            = -1
            @resource       = -1
        end

        def load(hp, atk, resource)
            @hp             = hp
            @atk            = atk
            @resource       = resource
        end

        def getDamage(damage)
            @hp -= damage
        end

        def attack(command)
            #
        end

        def destroyed?
            @hp <= 0
        end
    end

    class Battleship < Warship
    end
end
