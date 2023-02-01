
# Relax API Server

<div style="display:flex;">
   <img src="https://img.shields.io/badge/GO-gray?style=flat&logo=Go&logoColor=00ADD8"/>
	<img src="https://img.shields.io/badge/gORM-white?style=flat"/>
	<img src="https://img.shields.io/badge/gin-white?style=flat"/>
	  <img src="https://img.shields.io/badge/MariaDB-green?style=flat&logo=MariaDB&logoColor=003545"/>
  <img src="https://img.shields.io/badge/Redis-green?style=flat&logo=Redis&logoColor=DC382D"/>
    <img src="https://img.shields.io/badge/AWS_LightSail-green?style=flat"/>
  <img src="https://img.shields.io/badge/Docker-green?style=flat&logo=Docker&logoColor=2496ED"/>
  <img src="https://img.shields.io/badge/Nginx-green?style=flat&logo=NGINX&logoColor=009639"/>
</div>

<br/>
요가 종류로 요가학원을 검색하는 시스템 

## 동기


네이버 검색시스템 기준  
xx요가를 검색했을 때 학원을 노출시키기 위해서는  
각 사업장에서 해시태그에 요가 종류를 등록해야합니다. 

열 종류 이상의 요가수업을 진행하는 학원이 대부분이지만  
해시태그는 5개까지 등록할 수 있게 되어있습니다.  
이로 인해, 검색결과에 노출이 되지 않는 불편함이 있었습니다.  


## 구현

- 네이버에 등록된 요가원 중 서울에 소재한 곳 정보 크롤링
- Redis를 이용하여 자동완성 시스템 구축
- 요가 검색 시스템
- 인기검색어
- 요가 종류를 등록할 수 있는 어드민 시스템


# 사용법

## 프로덕션

주소 : http://52.78.0.155/

### 스크린샷


| ![](https://velog.velcdn.com/images/beardfriend/post/4a79cb0f-3f7d-4b63-ab01-61b7a0df90f5/image.png) | <img src="https://velog.velcdn.com/images/beardfriend/post/b64779ca-77b8-42e9-ae72-03224835e50d/image.png"/> |
| :----------------------------:| :---------------------:| 
| <img src="https://velog.velcdn.com/images/beardfriend/post/61824d66-c5c3-48f3-841f-ddff4c96578f/image.png"  />     |     ![](https://velog.velcdn.com/images/beardfriend/post/82c17457-6c07-424a-a8be-e06827e3066f/image.png) |

# 기타

프론트엔드 레포 : https://github.com/beardfriend/goyo_front

