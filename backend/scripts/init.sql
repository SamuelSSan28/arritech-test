-- Create database if not exists
CREATE DATABASE IF NOT EXISTS arritech_users CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Use the database
USE arritech_users;

-- Grant privileges to the user
GRANT ALL PRIVILEGES ON arritech_users.* TO 'arritech'@'%';
FLUSH PRIVILEGES;

-- The tables will be created automatically by GORM AutoMigrate
-- This script adds the date_of_birth column and removes the age column

-- Migration to add date_of_birth column if it doesn't exist
SET @col_exists = (SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS 
                   WHERE TABLE_SCHEMA = 'arritech_users' 
                   AND TABLE_NAME = 'users' 
                   AND COLUMN_NAME = 'date_of_birth');

SET @sql = IF(@col_exists = 0,
    'ALTER TABLE users ADD COLUMN date_of_birth DATE NOT NULL DEFAULT "1990-01-01"',
    'SELECT "Column date_of_birth already exists"');

PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt;

-- Migration to remove age column if it exists
SET @col_exists = (SELECT COUNT(*) FROM INFORMATION_SCHEMA.COLUMNS 
                   WHERE TABLE_SCHEMA = 'arritech_users' 
                   AND TABLE_NAME = 'users' 
                   AND COLUMN_NAME = 'age');

SET @sql = IF(@col_exists > 0,
    'ALTER TABLE users DROP COLUMN age',
    'SELECT "Column age does not exist"');

PREPARE stmt FROM @sql;
EXECUTE stmt;
DEALLOCATE PREPARE stmt; 