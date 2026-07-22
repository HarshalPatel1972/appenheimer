CREATE TYPE visibility_status AS ENUM ('draft', 'published', 'archived');
CREATE TYPE health_status AS ENUM ('operational', 'degraded', 'maintenance', 'outage', 'unknown');
CREATE TYPE verification_status AS ENUM ('verified', 'community', 'pending');
