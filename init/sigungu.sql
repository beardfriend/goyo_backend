CREATE TABLE `administration_division` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `si_do` varchar(20) NOT NULL,
  `si_gun_gu` varchar(20) NOT NULL,
  `ub_myun_dong` varchar(20) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=76 DEFAULT CHARSET=utf8mb4;


INSERT INTO `administration_division` (`si_do`, `si_gun_gu`, `ub_myun_dong`)
VALUES
	('서울', '양천구', ''),
	('서울', '강서구', NULL),
	('서울', '구로구', NULL),
	('서울', '영등포구', NULL),
	('서울', '금천구', NULL),
	('서울', '동작구', NULL),
	('서울', '관악구', NULL),
	('서울', '서초구', NULL),
	('서울', '강남구', NULL),
	('서울', '송파구', NULL),
	('서울', '강동구', NULL),
	('서울', '마포구', NULL),
	('서울', '용산구', NULL),
	('서울', '광진구', NULL),
	('서울', '성동구', NULL),
	('서울', '중구', NULL),
	('서울', '서대문구', NULL),
	('서울', '동대문구', NULL),
	('서울', '중랑구', NULL),
	('서울', '종로구', NULL),
	('서울', '성북구', NULL),
	('서울', '은평구', NULL),
	('서울', '강북구', NULL),
	('서울', '도봉구', NULL),
	('서울', '노원구', NULL);
