-- Ensure the uuid-ossp extension is available
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Add new UUID columns with default uuid_generate_v4()
ALTER TABLE owners ADD COLUMN new_id UUID DEFAULT uuid_generate_v4();
ALTER TABLE fields ADD COLUMN new_id UUID DEFAULT uuid_generate_v4();
ALTER TABLE reservations ADD COLUMN new_id UUID DEFAULT uuid_generate_v4();
-- ... rest of the migration script ...