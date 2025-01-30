DROP TABLE IF EXISTS project;
CREATE TABLE project (
  id  serial PRIMARY KEY,
  summary varchar(128),
  description text,
  tags varchar(128)
);

INSERT INTO project (summary, description, tags)
VALUES
  ('Realtime Markdown App', 'A realtime collaborative markdown note taking webapp, similar to Obsidian', 'REST,Backend'),
  ('Currency Converter', 'Convert currency taking into account the year and inflaction', 'Backend');
