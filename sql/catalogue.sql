-- create a catalogue for fields and types
CREATE TABLE IF NOT EXISTS catalogue (
    field TEXT,
    "type" TEXT,
    data_type TEXT NOT NULL,
    optional BOOLEAN DEFAULT false,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
	updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    PRIMARY KEY (field, "type")
);
COMMENT ON COLUMN catalogue.field IS 'Field name unique to the type';
COMMENT ON COLUMN catalogue.type IS 'Unique name of the type';
COMMENT ON COLUMN catalogue.data_type IS 'Data type of the field for generic interpretation';
COMMENT ON COLUMN catalogue.optional IS 'Optional field for marking a field optional for a type';
COMMENT ON COLUMN catalogue.created_at IS 'Timestamp at which the field was created';
COMMENT ON COLUMN catalogue.updated_at IS 'Timestamp at which the field was last updated';


-- create new fields (or a type)
INSERT INTO catalogue (field, "type", data_type, optional)
VALUES 
	('name', 'example', 'text', false)
	,('description', 'example', 'text', false)
;

-- retrieve all fields for a type
SELECT *
FROM catalogue
WHERE "type" = 'example';

-- delete field for a type
DELETE FROM catalogue
WHERE "type" = 'example' AND field = 'description';

-- delete a type
DELETE FROM catalogue
WHERE "type" = 'example';
