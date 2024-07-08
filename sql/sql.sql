-- Crie o banco de dados (execute apenas se você tiver privilégios de superusuário)
DO
$$
BEGIN
   IF NOT EXISTS (SELECT 1 FROM pg_database WHERE datname = 'turing') THEN
      PERFORM dblink_exec('dbname=postgres', 'CREATE DATABASE turing');
   END IF;
END
$$;

-- Conecte-se ao banco de dados turing
\c turing

-- Drope a tabela se ela já existir
DROP TABLE IF EXISTS usuarios;

-- Crie a tabela usuarios
CREATE TABLE usuarios (
    ID SERIAL PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    nick VARCHAR(50) NOT NULL UNIQUE,
    email VARCHAR(50) NOT NULL UNIQUE,
    senha VARCHAR(20) NOT NULL,
    criadoEm TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
