package constants

const (
	SchemaMTTS = `
	CREATE TABLE IF NOT EXISTS mtss_jobs
	(
		id         int PRIMARY KEY        NOT NULL,
		created_at TIMESTAMP DEFAULT now() NOT NULL,
		job  JSON
	);
	`
)
