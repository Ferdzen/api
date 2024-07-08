package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositorio de usuários
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuário no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {

	statement, err := repositorio.db.Prepare("INSERT INTO USUARIOS (nome, nick, email, senha) VALUES($1, $2, $3, $4) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}

	fmt.Println(resultado)

	// captura o ID inserido
	var IDInserido int64
	err = repositorio.db.QueryRow("SELECT LASTVAL()").Scan(&IDInserido)
	if err != nil {
		return 0, err
	}

	return uint64(IDInserido), nil
}

// Buscar traz todos os usuários que atendem um filtro de nome ou nick
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	rows, err := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE $1 or nick LIKE $2",
		nomeOuNick, nomeOuNick,
	)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var usuarios []modelos.Usuario

	for rows.Next() {
		var usuario modelos.Usuario

		if err = rows.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

// BuscarPorID traz um usuário do banco de dados
func (repositorio Usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	rows, err := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id = $1",
		ID,
	)
	if err != nil {
		return modelos.Usuario{}, err
	}
	defer rows.Close()

	var usuario modelos.Usuario

	if rows.Next() {
		if err = rows.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return modelos.Usuario{}, err
		}
	}

	return usuario, nil
}

// Atualizar altera as informações de um usuario no banco de dados.
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, err := repositorio.db.Prepare("update usuarios set nome = $1, nick= $2, email = $3 where id = $4")

	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); err != nil {
		return err
	}
	return nil
}

// Deletar exclui as informações de um usuário no banco de dados
func (repositorio Usuarios) Deletar(ID uint64) error {
	statement, err := repositorio.db.Prepare("delete from usuarios where id = $1")
	if err != nil {
		return err
	}

	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}
