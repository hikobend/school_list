SQLite format 3   @                                                                     ._� � ���'��                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               ��tableclassesclassesCREATE TABLE classes(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		class_number STRING,
		school_id INTEGER,
		created_at DATETIME)�8�CtablesessionssessionsCREATE TABLE sessions(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		email STRING,
		operator_id INTEGER,
		created_at DATETIME)/C indexsqlite_autoindex_sessions_1sessions��tableschoolsschoolsCREATE TABLE schools(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name STRING,
		operator_id STRING,
		created_at DATETIME)P++Ytablesqlite_sequencesqlite_sequenceCREATE TABLE sqlite_sequence(name,seq)�F�[tableoperatorsoperatorsCREATE TABLE operators(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		uuid STRING NOT NULL UNIQUE,
		name STRING,
		email STRING,
		password STRING,
		created_at DATETIME)1E indexsqlite_autoindex_operators_1operators          � r�V�                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                        � U+]M27440802-3c7d-11ed-8192-c2cd8da24975operator1email@email.com5baa61e4c9b93f3f0682250b6cf8331b7ee68fd82022-09-25 11:52:54.572015+09:00� U+]Md001da88-3c7c-11ed-a0d2-c2cd8da24975operator1email@email.com5baa61e4c9b93f3f0682250b6cf8331b7ee68fd82022-09-25 11:50:28.176762+09:00� U+]M512bdc5e-3c7c-11ed-b459-c2cd8da24975operator1email@email.com5baa61e4c9b93f3f0682250b6cf8331b7ee68fd82022-09-25 11:46:55.381527+09:00� U+]M26732512-3c7c-11ed-af60-c2cd8da24975operator1email@email.com5baa61e4c9b93f3f0682250b6cf8331b7ee68fd82022-09-25 11:45:43.706343+09:00
   ] �]��                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                             (U27440802-3c7d-11ed-8192-c2cd8da24975(Ud001da88-3c7c-11ed-a0d2-c2cd8da24975(U512bdc5e-3c7c-11ed-b459-c2cd8da24975'U	26732512-3c7c-11ed-af60-c2cd8da24975   � ���                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         classesschoolsoperators   g ��g                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                         1 %	Mfirst school2022-09-25 11:52:54.582249+09:001 %	Mfirst school2022-09-25 11:50:28.187404+09:001 %	Mfirst school2022-09-25 11:46:55.391322+09:00                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              
                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              � 8 �j88                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               � 93	KUpdate class_number2022-09-25 11:50:28.18791+09:000 #Kthird class2022-09-25 13:06:29.88059+09:001 %Ksecond class2022-09-25 12:20:30.00483+09:000 #	Mfirst class2022-09-25 11:52:54.582687+09:00   1#	Kfirst class2022-09-25 11:50:28.18791+09:00