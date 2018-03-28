package models

import (
	"time"
	"database/sql"
)


type Radicacion struct {
	NumeroRadicacion string    `xorm:"varchar(45) primary key not null unique 'numero_radicacion'"` //`json:"numeroRadicacion"`
	FechaRadicacion  time.Time `xorm:"TIMESTAMP 'fecha_radicacion' not null default 'CURRENT_TIMESTAMP'"`
	FechaDocumento   time.Time `xorm:"DATE 'fecha_documento' not null"`
	Asunto           string    `xorm:"VARCHAR(60) not null 'asunto'"`
	IDRemitente      string    `xorm:"varchar(45) not null 'id_remitente'"`
	PdfIdPdf         sql.NullInt64     `xorm:"INT(11) 'pdf_idPDF' null"`
	TiempoRespuesta  time.Time `xorm:"TIMESTAMP 'tiempo_respuesta'"`
}
