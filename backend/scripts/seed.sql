-- Seed data for testing - Always adds users (ignores duplicates gracefully)
USE arritech_users;

-- Clear existing seed data to avoid conflicts
DELETE FROM users WHERE email LIKE '%@email.com';

-- Insert Brazilian users
INSERT INTO users (name, email, date_of_birth, phone, address, created_at, updated_at) VALUES
('Ana Silva', 'ana.silva@email.com', '1990-03-15', '+5511987654321', 'Rua das Flores, 123, São Paulo - SP', NOW(), NOW()),
('Carlos Santos', 'carlos.santos@email.com', '1985-07-22', '+5511876543210', 'Av. Paulista, 456, São Paulo - SP', NOW(), NOW()),
('Maria Oliveira', 'maria.oliveira@email.com', '1992-11-08', '+5511765432109', 'Rua Augusta, 789, São Paulo - SP', NOW(), NOW()),
('Pedro Costa', 'pedro.costa@email.com', '1988-05-03', '+5511654321098', 'Rua Oscar Freire, 321, São Paulo - SP', NOW(), NOW()),
('Lucia Ferreira', 'lucia.ferreira@email.com', '1995-09-17', '+5511543210987', 'Alameda Santos, 654, São Paulo - SP', NOW(), NOW()),
('Roberto Lima', 'roberto.lima@email.com', '1982-12-30', '+5511432109876', 'Rua Consolação, 987, São Paulo - SP', NOW(), NOW()),
('Fernanda Souza', 'fernanda.souza@email.com', '1991-06-25', '+5511321098765', 'Rua Haddock Lobo, 111, São Paulo - SP', NOW(), NOW()),
('Marcos Pereira', 'marcos.pereira@email.com', '1987-02-14', '+5511210987654', 'Rua Bela Cintra, 222, São Paulo - SP', NOW(), NOW()),
('Juliana Rodrigues', 'juliana.rodrigues@email.com', '1993-08-09', '+5511109876543', 'Rua da Consolação, 333, São Paulo - SP', NOW(), NOW()),
('André Almeida', 'andre.almeida@email.com', '1989-04-27', '+5511098765432', 'Av. Rebouças, 444, São Paulo - SP', NOW(), NOW()),
('Camila Barbosa', 'camila.barbosa@email.com', '1994-01-12', '+5511987654322', 'Rua Teodoro Sampaio, 555, São Paulo - SP', NOW(), NOW()),
('Bruno Martins', 'bruno.martins@email.com', '1986-10-18', '+5511876543211', 'Rua Cardeal Arcoverde, 666, São Paulo - SP', NOW(), NOW()),
('Patricia Gomes', 'patricia.gomes@email.com', '1990-12-05', '+5511765432100', 'Rua Estados Unidos, 777, São Paulo - SP', NOW(), NOW()),
('Ricardo Dias', 'ricardo.dias@email.com', '1984-03-21', '+5511654321099', 'Alameda Franca, 888, São Paulo - SP', NOW(), NOW()),
('Carla Nascimento', 'carla.nascimento@email.com', '1996-07-13', '+5511543210988', 'Rua Pamplona, 999, São Paulo - SP', NOW(), NOW());

-- Insert international users to showcase different country flags
INSERT INTO users (name, email, date_of_birth, phone, address, created_at, updated_at) VALUES
('John Smith', 'john.smith@email.com', '1988-06-10', '+15551234567', '123 Main St, New York, NY', NOW(), NOW()),
('Emma Wilson', 'emma.wilson@email.com', '1992-09-15', '+447911123456', '456 Oxford St, London, UK', NOW(), NOW()),
('Pierre Dubois', 'pierre.dubois@email.com', '1985-03-22', '+33123456789', '789 Champs-Élysées, Paris, France', NOW(), NOW()),
('Hans Mueller', 'hans.mueller@email.com', '1990-11-08', '+49123456789', '321 Unter den Linden, Berlin, Germany', NOW(), NOW()),
('Yuki Tanaka', 'yuki.tanaka@email.com', '1993-07-14', '+81901234567', '654 Ginza St, Tokyo, Japan', NOW(), NOW()),
('Li Wei', 'li.wei@email.com', '1987-12-03', '+86123456789', '987 Nanjing Rd, Shanghai, China', NOW(), NOW()),
('Carlos Rodriguez', 'carlos.rodriguez@email.com', '1989-05-18', '+34612345678', '147 Gran Via, Madrid, Spain', NOW(), NOW()),
('Giulia Rossi', 'giulia.rossi@email.com', '1991-08-25', '+39391234567', '258 Via del Corso, Rome, Italy', NOW(), NOW()),
('Alex Johnson', 'alex.johnson@email.com', '1986-01-30', '+16135551234', '369 Queen St, Toronto, Canada', NOW(), NOW()),
('Sarah Brown', 'sarah.brown@email.com', '1994-04-12', '+61412345678', '741 Collins St, Melbourne, Australia', NOW(), NOW());

-- Show summary
SELECT 
    COUNT(*) as total_users,
    COUNT(CASE WHEN phone LIKE '+55%' THEN 1 END) as brazilian_users,
    COUNT(CASE WHEN phone NOT LIKE '+55%' THEN 1 END) as international_users
FROM users; 