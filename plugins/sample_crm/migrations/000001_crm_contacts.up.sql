CREATE TABLE IF NOT EXISTS plg_crm_contacts (
    id BIGSERIAL PRIMARY KEY,
    uuid VARCHAR(36) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255),
    phone VARCHAR(50),
    company VARCHAR(255),
    notes TEXT,
    created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMP(0) WITHOUT TIME ZONE
);

CREATE INDEX IF NOT EXISTS idx_plg_crm_contacts_email ON plg_crm_contacts(email);
CREATE INDEX IF NOT EXISTS idx_plg_crm_contacts_company ON plg_crm_contacts(company);
CREATE INDEX IF NOT EXISTS idx_plg_crm_contacts_deleted_at ON plg_crm_contacts(deleted_at);
