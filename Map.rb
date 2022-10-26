class Map
    def initialize(height=5, width=5)
        @width  = width
		@height = height
        @mapObjectList = MapObjectList.new
    end

    def find(objectID)
        x, y = @mapObjectList.find(objectID)
        return x, y
    end

    def find(x, y)
        mapObject = objectList.find(x, y)
        return mapObject
    end

    def updateInfo(objectID, x, y)
        mapObject = objectList.find(mapObject)
    end

    def setObject(obj, x, y)
        id = @mapObjectList.add(obj, x, y)
        return id
    end
    
end

class MapObjectList
    attr_reader :mapObjectList

    def initialize()
        @mapObjectList = {}
    end

    def add(obj, x, y)
        id = obj.id
        mapObject = MapObject.new(id, obj, x, y)
        @mapObjectList[id] = mapObject

        return id
    end

    def update(id, x, y)
        @mapObjectList[id].updatePosition(x, y)
    end

    def find(id)
        return @mapObjectList[id].x, @mapObjectList[id].y
    end

    def find(x, y)
        @mapObjectList.each do |mapObject|
            return mapObject if position?(x, y)
        end
    end

    def delete
        #
    end
end

class MapObject
    attr_reader :id
    attr_accessor :obj, :x, :y
    def initialize(id, obj, x, y)
        @id = id
        @obj = obj
        @x = x
        @y = y
    end

    def updatePosition(x, y)
        @x = x
        @y = y
    end

    def position?(x, y)
        return true if @x==x and @y==y
    end
end