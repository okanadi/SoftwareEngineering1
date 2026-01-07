-- Erweiterung für UUIDs aktivieren
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- ==========================================
-- Enums (Datentypen)
-- ==========================================
CREATE TYPE user_role AS ENUM ('admin', 'innendienst', 'handwerker');
CREATE TYPE step_status AS ENUM ('geplant', 'in_arbeit', 'fertiggestellt', 'verzögert', 'problem');

-- ==========================================
-- 1. USERS
-- ==========================================
CREATE TABLE IF NOT EXISTS users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL,
    role user_role NOT NULL DEFAULT 'handwerker',
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- ==========================================
-- 2. PROJECTS
-- ==========================================
CREATE TABLE IF NOT EXISTS projects (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    manager_id UUID REFERENCES users(id),
    customer_lastname TEXT NOT NULL,
    address TEXT NOT NULL,
    description TEXT,
    start_date DATE,
    end_date DATE,
    progress TEXT DEFAULT 'geplant',
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- Index für schnellen Login
CREATE INDEX IF NOT EXISTS idx_projects_login ON projects(id, customer_lastname);

-- ==========================================
-- 3. PROJECT_STEPS
-- ==========================================
CREATE TABLE IF NOT EXISTS project_steps (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id UUID REFERENCES projects(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    description TEXT,
    start_date DATE,
    end_date DATE,
    progress step_status DEFAULT 'geplant',
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- ==========================================
-- 4. PROJECT_HISTORY (Updates)
-- ==========================================
CREATE TABLE IF NOT EXISTS project_history (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    step_id UUID REFERENCES project_steps(id) ON DELETE CASCADE,
    user_id UUID REFERENCES users(id),
    new_status step_status NOT NULL,
    note TEXT,
    timestamp TIMESTAMPTZ DEFAULT NOW()
);

-- ==========================================
-- 5. MEDIA (S3 Referenzen)
-- ==========================================
CREATE TABLE IF NOT EXISTS media (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    history_id UUID REFERENCES project_history(id) ON DELETE CASCADE,
    s3_key TEXT NOT NULL,
    file_type TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW()
);

-- ==========================================
-- OPTIONAL: Ein Test-User zum Einloggen (Passwort: "secret")
-- ==========================================
-- Das Passwort ist hier Klartext für Demo-Zwecke. 
-- In Produktion muss das ein Hash sein (Argon2)!
-- INSERT INTO users (name, email, password, role) 
-- VALUES ('Max Handwerker', 'max@smartbuilders.de', 'secret', 'handwerker')
-- ON CONFLICT DO NOTHING;