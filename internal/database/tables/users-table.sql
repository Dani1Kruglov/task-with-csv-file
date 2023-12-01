CREATE TABLE users (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(100),
                       job_titles VARCHAR(255),
                       department VARCHAR(255),
                       full_or_part_time VARCHAR(1),
                       salary_or_hourly VARCHAR(6),
                       typical_hours INTEGER,
                       annual_salary DOUBLE PRECISION,
                       hourly_rate DOUBLE PRECISION
);
