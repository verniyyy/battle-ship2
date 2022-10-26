require 'securerandom'

require_dir = "./infrastructure/objects/warship"
Dir["#{require_dir}/*.rb"].each do |file|
    require file
end

module WarshipRepositoryMock
    module_function

    def get(id, userId)
        WarshipObjects::Battleship.new(id, userId, '010000', "000000", "戦艦ブランク体")
    end

    def list(userId)
        list = []
        5.times do |i|
            # Battleship.new(id, userId, warshipId, individualId, name)
            list << WarshipObjects::Battleship.new(SecureRandom.uuid, userId, '010000', "00000#{i.to_s}", "戦艦ブランク体")
        end
        return list
    end
end

module WarshipRepository
    module_function

    extend WarshipRepositoryMock

    def get(id)
        #super
    end

    def get(id, userId)
        super
    end
    
    def list(userId)
        super
    end
end


userId = '000000'
#pp WarshipRepository.list(userId)

#pp WarshipRepository.get(SecureRandom.uuid)