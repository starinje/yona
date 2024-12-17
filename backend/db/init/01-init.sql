-- Create the database if it doesn't exist
CREATE DATABASE yona;

-- Connect to the database
\c yona;

-- Create extension if needed
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Grant privileges
GRANT ALL PRIVILEGES ON DATABASE yona TO postgres; 