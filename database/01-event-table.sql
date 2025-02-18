CREATE TABLE IF NOT EXISTS events (
    hostname               String,
    site_id                String,
    site_name              String,
    event_name             String,
    start_time             DateTime,
    pathname               String,
    navigation_from        String,
    entry_meta_key         Array(String),
    entry_meta_value       Array(String),
    utm_medium             Nullable(String),
    utm_source             Nullable(String),
    utm_campaign           Nullable(String),
    utm_content            Nullable(String),
    utm_term               Nullable(String),
    referrer               String,
    referrer_source        String,
    screen_size            String,
    device                 String,
    operating_system       String,
    operating_system_version String,
    browser                String,
    browser_version        String
) ENGINE = MergeTree() ORDER BY start_time

