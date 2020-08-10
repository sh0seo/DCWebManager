# DCWebManager

DigiCAP Web Manager Renewal

## TODO
- [X] 기존 프로젝트(spring) 기반으로 go 프로젝트 생성
- [ ] UI 재 개발(JSP -> HTML + template)
  - [X] login
  - [X] index
  - [ ] 사용자 등록
  - [ ] 사용자 수정/삭제
  - [ ] 카테고리/메뉴 관리
  - [ ] 주문취소 관리
  - [ ] 월말공제정산
  - [ ] 손김결제정산
  - [ ] DB 연결
  - [ ] User 서버, Api 서버 연동 

## Build
### go-bindata 설치
```
go get -u github.com/jteeuwen/go-bindata/...
```
### resource build
```
go-bindata static/... templates/...
go fmt bindata.go
```
### single binary build
```
go build
```

## 로그인
<img src="static/github/1.png" >

## 카페 메뉴 관리
<img src="static/github/2.png">
