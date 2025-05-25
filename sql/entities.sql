-- create a catalogue for fields and types
CREATE TABLE IF NOT EXISTS entities (
    "id" TEXT,
    "type" TEXT,
    properties JSONB NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    PRIMARY KEY ("id")
);
COMMENT ON COLUMN entities.id IS 'Unique identifier generated using UUID';
COMMENT ON COLUMN entities.type IS 'A categorical field describing the nature of an entity';
COMMENT ON COLUMN entities.properties IS 'Keys and values relevant to the type';
COMMENT ON COLUMN entities.created_at IS 'Timestamp at which the entity was created';
COMMENT ON COLUMN entities.updated_at IS 'Timestamp at which the entity was last updated';

-- create new entities
INSERT INTO entities ("id", "type", properties)
VALUES 
	('example-id', 'example', '{"key": "value"}')
;

-- retrieve an entity by "id"
SELECT *
FROM entities
WHERE "id" = 'example-id';

-- delete an entity by id
DELETE FROM entities
WHERE "id" = 'example-id';
