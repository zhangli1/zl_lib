package zl_lib

import (
	"database/sql"
	//"fmt"
	//"os"

	l4g "zl_lib/code.google.com/p/log4go"
	_ "zl_lib/github.com/go-sql-driver/mysql"
)

type Mysql struct {
	Cfg       string
	l4gLogger *l4g.Logger
	db        *sql.DB
}

func NewSQL(cfg string, logger *l4g.Logger) *Mysql {
	return &Mysql{Cfg: cfg, l4gLogger: logger}
}

func (s *Mysql) Init() {
	db, err := sql.Open("mysql", s.Cfg)
	if err != nil {
		s.l4gLogger.Error("connection db fail", err)
		//os.Exit(-1)
	}
	s.db = db
}

func (s *Mysql) Close() {
	s.db.Close()
}

func (s *Mysql) Query(q_sql string) interface{} {
	rows, err := s.db.Query(q_sql)
	//defer s.Close()
	defer rows.Close()

	if err != nil {
		s.l4gLogger.Error("query db fail", q_sql, err)
		//os.Exit(-1)
		return false
	}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			//			fmt.Println("get data fail", err)
			//			os.Exit(-1)
			s.l4gLogger.Error(q_sql, err)
			return false

		}
		return name
		//fmt.Printf("%s\n", name)
	}
	if err := rows.Err(); err != nil {
		//		fmt.Println("get data fail2", err)
		//		os.Exit(-1)
		s.l4gLogger.Error(q_sql, err)
		return false
	}
	return false
}

func (s *Mysql) Exec(e_sql string) bool {
	_, err := s.db.Exec(e_sql)
	//defer s.db.Close()
	if err != nil {
		//fmt.Println("insert into data fail:", err)
		//os.Exit(-1)
		s.l4gLogger.Error(e_sql, err)
		return false
	}
	return true
}
