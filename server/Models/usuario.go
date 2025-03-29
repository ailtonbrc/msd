package models

type Usuario struct {
	ID                    int     `json:"id"`
	Nome                  string  `json:"nome"`
	Email                 string  `json:"email"`
	Senha                 string  `json:"senha"`
	Perfil                string  `json:"perfil"`
	ClinicaID             *int    `json:"clinica_id"`
	SupervisorID          *int    `json:"supervisor_id"`
	Ativo                 bool    `json:"ativo"`
	DataInicioInatividade *string `json:"data_inicio_inatividade"`
	DataFimInatividade    *string `json:"data_fim_inatividade"`
	MotivoInatividade     *string `json:"motivo_inatividade"`
}
