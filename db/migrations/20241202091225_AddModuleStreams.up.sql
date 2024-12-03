BEGIN;

CREATE TABLE IF NOT EXISTS module_streams (
      uuid UUID UNIQUE NOT NULL PRIMARY KEY,
      created_at TIMESTAMP WITH TIME ZONE,
      updated_at TIMESTAMP WITH TIME ZONE,
      Name text NOT NULL,
      Stream text NOT NULL,
      Version text NOT NULL,
      Context text NOT NULL,
      Arch text NOT NULL,
      Summary text NOT NULL,
      Description text NOT NULL,
      Package_names text[] NOT NULL,
      Hash_value text NOT NULL
);

CREATE TABLE IF NOT EXISTS repositories_module_streams (
       repository_uuid UUID NOT NULL,
       module_stream_uuid UUID NOT NULL
);

ALTER TABLE ONLY repositories_module_streams
DROP CONSTRAINT IF EXISTS repositories_module_streams_pkey,
ADD CONSTRAINT repositories_module_streams_pkey PRIMARY KEY (repository_uuid, module_stream_uuid);

ALTER TABLE ONLY repositories_module_streams
DROP CONSTRAINT IF EXISTS fk_repositories_module_streams_mstream,
ADD CONSTRAINT fk_repositories_module_streams_mstream
FOREIGN KEY (module_stream_uuid) REFERENCES module_streams(uuid)
ON DELETE CASCADE;

ALTER TABLE ONLY repositories_module_streams
DROP CONSTRAINT IF EXISTS fk_repositories_module_streams_repository,
ADD CONSTRAINT fk_repositories_module_streams_repository
FOREIGN KEY (repository_uuid) REFERENCES repositories(uuid)
ON DELETE CASCADE;

ALTER TABLE ONLY module_streams
DROP CONSTRAINT IF EXISTS fk_module_streams_uniq,
ADD CONSTRAINT fk_module_streams_uniq UNIQUE (hash_value);

COMMIT;
