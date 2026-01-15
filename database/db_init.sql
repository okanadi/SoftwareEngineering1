-- Erweiterung für UUIDs aktivieren
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- ==========================================
-- 1. SCHEMA DEFINITION (Tabellen)
-- ==========================================

-- ENUMS (verhindert Fehler bei Neustart)
DO $$ BEGIN
    CREATE TYPE user_role AS ENUM ('admin', 'innendienst', 'handwerker');
    CREATE TYPE status AS ENUM ('geplant', 'in_arbeit', 'fertiggestellt', 'verzögert', 'problem');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

-- USERS
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role user_role NOT NULL DEFAULT 'handwerker',
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- PROJECTS
CREATE TABLE IF NOT EXISTS projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    manager_id UUID REFERENCES users(id),
    customer_lastname TEXT NOT NULL,
    address TEXT NOT NULL,
    description TEXT,
    start_date DATE,
    end_date DATE,
    progress status DEFAULT 'geplant',
    created_at TIMESTAMPTZ DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS idx_projects_login ON projects(id, customer_lastname);

-- PROJECT_STEPS
CREATE TABLE IF NOT EXISTS project_steps (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID REFERENCES projects(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    description TEXT,
    start_date DATE,
    end_date DATE,
    progress status DEFAULT 'geplant',
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- PROJECT_HISTORY
CREATE TABLE IF NOT EXISTS project_history (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    step_id UUID REFERENCES project_steps(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id),
    new_status status NOT NULL,
    note TEXT,
    timestamp TIMESTAMPTZ DEFAULT NOW()
);

-- MEDIA
CREATE TABLE IF NOT EXISTS media (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    history_id UUID REFERENCES project_history(id) ON DELETE CASCADE,
    s3_key TEXT NOT NULL,
    file_type TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- ==========================================
-- 2. SEED DATA (Testdaten)
-- ==========================================

-- USERS
INSERT INTO users (name, email, password, role) VALUES 
    ('Adam Admin', 'admin@firma.de', '$2a$10$xwqHfMlOBReRFTBbUc8lk.p/xW/u2yLUOEHL9xOeweedQwZ13w4y.', 'admin'),
    ('Inga Innendienst', 'inga@firma.de', '$2a$10$xwqHfMlOBReRFTBbUc8lk.p/xW/u2yLUOEHL9xOeweedQwZ13w4y.', 'innendienst'),
    ('Hans Handwerker', 'hans@firma.de', '$2a$10$xwqHfMlOBReRFTBbUc8lk.p/xW/u2yLUOEHL9xOeweedQwZ13w4y.', 'handwerker'),
    ('Franz Fliese', 'franz@firma.de', '$2a$10$xwqHfMlOBReRFTBbUc8lk.p/xW/u2yLUOEHL9xOeweedQwZ13w4y.', 'handwerker'),
    ('Manni Maler', 'manni@firma.de', '$2a$10$xwqHfMlOBReRFTBbUc8lk.p/xW/u2yLUOEHL9xOeweedQwZ13w4y.', 'handwerker')
ON CONFLICT (email) DO NOTHING;

-- PROJECTS (Manager: Inga)
INSERT INTO projects (manager_id, customer_lastname, address, description, start_date, end_date, progress)
VALUES 
    ((SELECT id FROM users WHERE email = 'inga@firma.de'), 'Müller', 'Hauptstraße 10, Berlin', 'Badsanierung', CURRENT_DATE - 5, CURRENT_DATE + 10, 'in_arbeit'),
    ((SELECT id FROM users WHERE email = 'inga@firma.de'), 'Schmidt', 'Waldweg 5, München', 'Neubau Garage', CURRENT_DATE + 14, CURRENT_DATE + 30, 'geplant'),
    ((SELECT id FROM users WHERE email = 'inga@firma.de'), 'Weber', 'Hafenstraße 99, Hamburg', 'Rohrbruch Keller', CURRENT_DATE - 2, CURRENT_DATE + 5, 'problem');

-- STEPS (Für alle 3 Projekte)
-- Müller Steps
INSERT INTO project_steps (project_id, title, description, start_date, end_date, progress) VALUES 
    ((SELECT id FROM projects WHERE customer_lastname = 'Müller'), 'Demontage', 'Altbestand entfernen', CURRENT_DATE-5, CURRENT_DATE-3, 'fertiggestellt'),
    ((SELECT id FROM projects WHERE customer_lastname = 'Müller'), 'Rohrinstallation', 'Neue Leitungen', CURRENT_DATE-2, CURRENT_DATE, 'in_arbeit'),
    ((SELECT id FROM projects WHERE customer_lastname = 'Müller'), 'Fliesen', 'Wand und Boden', CURRENT_DATE+3, CURRENT_DATE+7, 'geplant');

-- Schmidt Steps
INSERT INTO project_steps (project_id, title, description, start_date, end_date, progress) VALUES 
    ((SELECT id FROM projects WHERE customer_lastname = 'Schmidt'), 'Erdaushub', 'Fundament graben', CURRENT_DATE+15, CURRENT_DATE+17, 'geplant'),
    ((SELECT id FROM projects WHERE customer_lastname = 'Schmidt'), 'Fundament', 'Betonieren', CURRENT_DATE+18, CURRENT_DATE+19, 'geplant');

-- Weber Steps (Problemfall)
INSERT INTO project_steps (project_id, title, description, start_date, end_date, progress) VALUES 
    ((SELECT id FROM projects WHERE customer_lastname = 'Weber'), 'Leckortung', 'Gefunden', CURRENT_DATE-2, CURRENT_DATE-2, 'fertiggestellt'),
    ((SELECT id FROM projects WHERE customer_lastname = 'Weber'), 'Rohraustausch', 'Kupferrohr ersetzen', CURRENT_DATE, CURRENT_DATE, 'problem');

-- ==========================================
-- OPTIONAL: Ein Test-User zum Einloggen (Passwort: "secret")
-- ==========================================
-- Das Passwort ist hier Klartext für Demo-Zwecke. 
-- In Produktion muss das ein Hash sein (Argon2)!
-- INSERT INTO users (name, email, password, role) 
-- VALUES ('Max Handwerker', 'max@smartbuilders.de', 'secret', 'handwerker')
-- ON CONFLICT DO NOTHING;