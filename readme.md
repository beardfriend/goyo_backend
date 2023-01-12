
# Relax
✔️ 서비스 : 요가 검색 서비스  
✔️ 기여도 : 본인 100%  
✔️ URL :  [👆Click](http://52.78.0.155/)  
✔️ 백엔드 : [👆Click](https://github.com/beardfriend/goyo_backend)  
✔️ 프론트엔드 : [👆Click](https://github.com/beardfriend/goyo_front)  

# 1. 개요

## 1.1. 사용기술

🔎 프론트  
<div style="display:flex;">
  <img src="https://img.shields.io/badge/TYPESCRIPT-gray?style=flat&logo=TypeScript&logoColor=3178C6"/>
  <img src="https://img.shields.io/badge/React-white?style=flat&logo=React&logoColor=61DAFB"/>
  <img src="https://img.shields.io/badge/Redux-white?style=flat&logo=Redux&logoColor=764ABC"/>
  <img src="https://img.shields.io/badge/Styled_Component-white?style=flat&logo=styled-components&logoColor=DB7093"/>
</div>
<br>
🔎 백엔드  
<div style="display:flex;">
   <img src="https://img.shields.io/badge/GO-gray?style=flat&logo=Go&logoColor=00ADD8"/>
	<img src="https://img.shields.io/badge/GORM-white?style=flat"/>
	<img src="https://img.shields.io/badge/GIN-white?style=flat"/>
</div>
<br>

🔎 데이터베이스  
<div style="display:flex;">
  <img src="https://img.shields.io/badge/MariaDB-green?style=flat&logo=MariaDB&logoColor=003545"/>
  <img src="https://img.shields.io/badge/Redis-green?style=flat&logo=Redis&logoColor=DC382D"/>
</div>
<br>
🔎 인프라  
<div style="display:flex;">
  <img src="https://img.shields.io/badge/AWS_LightSail-green?style=flat"/>
  <img src="https://img.shields.io/badge/Docker-green?style=flat&logo=Docker&logoColor=2496ED"/>
  <img src="https://img.shields.io/badge/NGINX-green?style=flat&logo=NGINX&logoColor=009639"/>
</div>

## 1.2. API 명세


`학원`

- 요가원 리스트 네이버 크롤링 API
- 요가원 상세정보 크롤링 API
- 학원 검색 API

`요가 종류`

- 인기검색어 등록, 조회API
- 검색어 자동완성 등록,조회 API


`어드민`

- 요가원에서 진행중인 요가 CRUD API
- 요가원 리스트 API
  - 행정구역(자치구) 필터
  - 태그가 등록됐는지 여부 필터
  
- 행정구역 리스트 API


## 1.3. 페이지
- 어드민페이지 (나만 사용하는..)
 <img src="https://velog.velcdn.com/images/beardfriend/post/3e86ba83-0d61-463e-abd9-088c888e1b22/image.png"/> 
 
가장 왼쪽에는 요가원 리스트가 나옵니다.
구(행정구역)별로 필터링이 됩니다.
요가 태그가 등록됐는지 안됐는지 필터링이 됩니다.

요가원을 클릭하면,
가장 오른쪽 영역에 요가원의 네이버 플레이스 페이지가 나옵니다.

시간표를 보고 수동으로 요가원에 태그를 입력했습니다. (맨 아래에서 자동화 방법에 대한 고찰 설명...)

요가 종류를 등록할 때는 `,`로 멀티등록을 했습니다.







- 검색결과, 검색페이지 (유저가 보는 페이지)


| ![](https://velog.velcdn.com/images/beardfriend/post/4a79cb0f-3f7d-4b63-ab01-61b7a0df90f5/image.png) | <img src="https://velog.velcdn.com/images/beardfriend/post/b64779ca-77b8-42e9-ae72-03224835e50d/image.png"/> |
| :----------------------------:| :---------------------:| 
| <img src="https://velog.velcdn.com/images/beardfriend/post/61824d66-c5c3-48f3-841f-ddff4c96578f/image.png"  />     |     ![](https://velog.velcdn.com/images/beardfriend/post/82c17457-6c07-424a-a8be-e06827e3066f/image.png) |




## 1.4. 제작 이유
>이 프로젝트는, 
제가 좋아하는 취미활동인 요가를 알아보며,
개인적으로 느꼈던 불편함을 해소하기 위해 시작했습니다.

---
요가는 종류가 다양합니다.
(아쉬탕가, 하타, 플라잉 등) 

저는 다양한 요가를 체험해보고 싶었고,
네이버에서 종류별로 검색했을 때 ("양천구 하타 요가")
만족스러운 결과를 얻지는 못했습니다.

![](https://velog.velcdn.com/images/beardfriend/post/1a19dc97-959d-4cfd-9910-dc4f7eddd4c1/image.png)

네이버에서는 각 사업장마다 해시태그를 등록할 수 있도록 해놨습니다. 
검색할 단어가 해시태그에 포함되어 있으면,
검색결과에 노출이 됩니다.

![](https://velog.velcdn.com/images/beardfriend/post/e1668683-bdb0-421a-91c8-edd18fc1e240/image.png)

요가원 대부분은 
수업시간, 수업할 요가, 수업진행을 도와주는 선생님 이름을
시간표에 공지를 합니다.

**_시간표에는 "하타요가" 수업이 있다고 안내가 되어 있지만,  
해시태그에는 없어, 검색이 되지 않는 불편함이 있습니다._**




## 1.5. 목표
**`요가 종류`로 검색했을 때 

해당 요가를 진행하고 있는 학원을 노출시켜주는 검색 시스템 구축. **

### 1.5.1. 세부 계획
**동작**
- 네이버에 등록된 요가 학원을 크롤링해서 내 시스템에 등록하자. (서울에 소재한).
- 요가원마다 취급하는 요가를 내 시스템에 등록하자.
- 내 시스템에 등록된 요가종류를 검색할 때 자동완성이 되게끔 하자.


**시스템**
- 수요에 따라 시스템 크기를 정하기.
- react, golang, redis, mariadb
- 최대한 심플하게.



# 2. 기능 설명

## 2.1. 네이버에 등록된 요가원 크롤링
[CrawlNaverPlacesService](https://github.com/beardfriend/goyo_backend/blob/main/modules/academy/academy_service.go#L30)

**로직**

 1. `행정구역+요가` 키워드로 네이버 플레이스에 검색. (서울시에 존재하는 행정구)
 	ex) 양천구 요가 

2. 검색 결과 카테고리가 (요가, 명상)이면 DB존재여부 체크 후 수집.

**특이사항**
- 네이버플레이스 검색은 Graphql를 사용하고 있습니다.

## 2.2. 각 학원의 대표 이미지 긁어오기.

[CrawlMobileNaverPlaceService](https://github.com/beardfriend/goyo_backend/blob/main/modules/academy/academy_service.go#L104)

**로직**

1. DB에 등록된 네이버 고유 ID리스트 가져오기
2. `https://m.place.naver.com/place/고유ID/home` 로 접속
3. 이미지 크롤링.
4. 데이터베이스 업데이트


## 2.3. 검색 자동완성

요가 종류는 개수가 한정적입니다.  
자동완성 데이터가 그리 많지 않다는 점을 고려했습니다.  

엘라스틱 서치는 규모대비 손이 너무 많이 가기 때문에,  
익숙한 레디스를 선택했습니다.  

### 2.3.1. 키워드 등록하기

[CronYogaSortsService](https://github.com/beardfriend/goyo_backend/blob/main/modules/yoga/yoga_controller.go#L120)

하나의 키워드에 대해  
초성 + 중성 + 종성을 모두 조합해 (종성이 있는 경우는 초성 +중성 조합은 제외)  
Key에 등록하고 Value에는 키워드를 넣습니다.  

예를들면 빈야사는 이런식입니다.  
ㅂ, 빈, 빈ㅇ ... 빈야사  

|![](https://velog.velcdn.com/images/beardfriend/post/8094fe47-9b6f-4f12-8bfe-69f9687da567/image.png)|![](https://velog.velcdn.com/images/beardfriend/post/4b76b29d-d56e-41a6-b7f5-d1daec3e4b4d/image.png)|
|--|--|

### 2.3.2. 키워드 조회

[GetYogaSortsService](https://github.com/beardfriend/goyo_backend/blob/main/modules/yoga/yoga_controller.go#L60)



![](https://velog.velcdn.com/images/beardfriend/post/6078fd99-fa22-40ff-a7e6-b08cdea98a86/image.png)

레디스에서 검색 결과가 1개일 때, 1개만 노출


![](https://velog.velcdn.com/images/beardfriend/post/0cad7217-5d63-4ae0-8f1f-736263944eb1/image.png)

2개일 때,   
정확히 일치하는 검색결과가 있으면, 첫 번째 노출,  
연관된 건 두 번째 노출  


![](https://velog.velcdn.com/images/beardfriend/post/143931c9-ce8c-49e9-8c69-bb4b156aa8fd/image.png)

검색 결과가 3개 이상일 때는

검색 결과가 8개를 넘지 않으면  
가장 가깝게 일치하는 1개와 나머지.

검색 결과가 8개를 넘으면  
사람들이 검색 많이한 2개와 + 랜덤으로 6개 


에러가 많이 보여서, 좀 더 세밀하게 설계할 필요성을 느꼈습니다.  
아니면 이미 검증된 엘라스틱을 사용하는 편이 가장 좋아 보입니다.


## 2.4. 인기 검색어

### 2.4.1. 등록하기

[RankingService](https://github.com/beardfriend/goyo_backend/blob/main/modules/yoga/yoga_controller.go#L200)

![](https://velog.velcdn.com/images/beardfriend/post/2a6a3dba-a08f-4a80-9bf5-934e5589a468/image.png)

검색할 때마다, 레디스에 count가 되는데,  
5분마다 크론탭을 사용하여  
Redis에서 정보를 가져와 집계하여 결과를 DB에 넣습니다.

### 2.4.2. 가져오기

랭킹 순서대로 키워드를 가져옵니다.


# 3. 히스토리

## 3.1. OCR를 사용한 요가 태그 등록 자동화 구상
1. 요가원이 공식적으로 등록한 홈페이지 혹은 블로그에서 이미지URL 크롤링하여 출처와 URL 디비에 저장합니다.
2. 이미지로 된 시간표 정보를 텍스트화 합니다..
3. [월,화,수,목,금 시간표 .. sunday, monday] 필터를 거쳐 이 이미지가 시간표임을 확률로 확인합니다.
4. 영어, 한글로된 [아쉬탕가,'ashtanga',아쉬탕'...] 필터를 거쳐 필터에 걸리면 요가원에 등록합니다.



OCR를 사용하면 되겠다 싶어서 OCR를 사용했으나, 이미지 전처리과정이 꽤나 까다로워서 포기했습니다.  
네이버클라우드플랫폼에서 이미지를 텍스트로 바꿔주는 서비스가 있어서 사용해봤는데,  
너무 잘되지만, 비용이 너무 비싸서 시작할 엄두가 나지 않았습니다..

## 3.2.프로젝트에 관한 고찰

### 3.2.1. 배포
이 서비스는 공식적으로 사용하기는 어렵다는 판단하에,  
포트폴리오 용도로 사용하기 위해 

docker로 빌드하여 
aws lightsail에 2개의 인스턴스에  
프론트, (백,디비,레디스)  
scp, compose를 사용하여 수동으로 배포하였습니다.

nginx를 사용했고,  
백엔드는 systemd를 이용하여 어플리케이션을 띄웠습니다.




