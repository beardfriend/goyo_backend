package validators

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

/**
1. 구조체를 선언할 때 gin 에서는 validate 대신 binding


2. gin에서 아래와 같은 방식으로 customValidation 등록

server := gin.Default()

if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	v.RegisterValidation("customUrl", Url())
	v.RegisterValidation("customPhone", Phone())
	v.RegisterValidation("customPassword", Password())
}

*/

type TestPasswordStruct struct {
	Password string `validate:"customPassword,min=5,max=20"`
}

// 1자 이상의 문자열(특수문자)과 숫자를 포함한 5자 20이하의 패스워드 ... 가능한 특수문자는 !@$%^&*
func TestPassword(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("customPassword", Password())

	t.Run("정상 입력된 경우", func(t *testing.T) {
		t.Run("소문자 + 숫자", func(t *testing.T) {
			s := TestPasswordStruct{Password: "awesome1"}
			if err := validate.Struct(s); err != nil {
				t.Error("에러")
			}
		})

		t.Run("대소문자 + 숫자", func(t *testing.T) {
			s := TestPasswordStruct{Password: "Awesome1"}
			if err := validate.Struct(s); err != nil {
				t.Error("에러")
			}
		})

		t.Run("소문자 + 숫자 + 특수문자", func(t *testing.T) {
			s := TestPasswordStruct{Password: "awesome1!"}
			if err := validate.Struct(s); err != nil {
				t.Error("에러")
			}
		})

		t.Run("대소문자 + 숫자 + 특수문자", func(t *testing.T) {
			s := TestPasswordStruct{Password: "Awesomeasfsafwq1!"}
			if err := validate.Struct(s); err != nil {
				t.Error("에러")
			}
		})

		t.Run("대소문자 + 숫자 + 특수문자", func(t *testing.T) {
			s := TestPasswordStruct{Password: "!@Awesome1!"}
			if err := validate.Struct(s); err != nil {
				t.Error("에러")
			}
		})
	})

	t.Run("잘못 입력된 경우", func(t *testing.T) {
		t.Run("5자 미만", func(t *testing.T) {
			s := TestPasswordStruct{Password: "asd1"}
			if err := validate.Struct(s); err == nil {
				t.Error("에러")
			}
		})

		t.Run("20자 이상", func(t *testing.T) {
			s := TestPasswordStruct{Password: "asdasfasfsafaffaffaffaffasfasfa123f"}
			if err := validate.Struct(s); err == nil {
				t.Error("에러")
			}
		})

		t.Run("한글", func(t *testing.T) {
			s := TestPasswordStruct{Password: "ㅁㄴㄹㄹㄴ"}
			if err := validate.Struct(s); err == nil {
				t.Error("에러")
			}
		})

		t.Run("소문자만", func(t *testing.T) {
			s := TestPasswordStruct{Password: "asdff"}
			if err := validate.Struct(s); err == nil {
				t.Error("에러")
			}
		})

		t.Run("대문자만", func(t *testing.T) {
			s := TestPasswordStruct{Password: "ASDFF"}
			if err := validate.Struct(s); err == nil {
				t.Error("에러")
			}
		})

		t.Run("숫자만", func(t *testing.T) {
			s := TestPasswordStruct{Password: "12345"}
			if err := validate.Struct(s); err == nil {
				t.Error("에러")
			}
		})

		t.Run("금지된 특수문자", func(t *testing.T) {
			s := TestPasswordStruct{Password: "asdf123#"}
			if err := validate.Struct(s); err == nil {
				t.Error("에러")
			}
		})
	})
}

type TestUrlStruct struct {
	Url string `validate:"customUrl"`
}

func TestUrl(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("customUrl", Url())
	t.Run("정상 입려된 경우", func(t *testing.T) {
		t.Run("http로 시작", func(t *testing.T) {
			s := TestUrlStruct{Url: "http://www.naver.com"}
			if err := validate.Struct(s); err != nil {
				t.Error("에러")
			}
		})

		t.Run("https로 시작", func(t *testing.T) {
			s := TestUrlStruct{Url: "https://www.naver.com"}
			if err := validate.Struct(s); err != nil {
				t.Error("에러")
			}
		})

		t.Run("www 없는 경우", func(t *testing.T) {
			s := TestUrlStruct{Url: "https://github.io"}
			if err := validate.Struct(s); err != nil {
				t.Error("에러")
			}
		})

		t.Run("쿼리스트링 및 하위경로", func(t *testing.T) {
			s := TestUrlStruct{Url: "https://www.naver.com/adsd?query=asd&ff=qwe"}
			if err := validate.Struct(s); err != nil {
				t.Error("에러")
			}
		})
		t.Run("쿼리스트링 및 하위경로", func(t *testing.T) {
			s := TestUrlStruct{Url: "https://www.naver.com/adsd/asd?ffb=asd"}
			if err := validate.Struct(s); err != nil {
				t.Error("에러")
			}
		})

		t.Run("#이 있을 때", func(t *testing.T) {
			s := TestUrlStruct{Url: "https://www.naver.com/adsd/asd#태그"}
			if err := validate.Struct(s); err != nil {
				t.Error("에러")
			}
		})
	})

	t.Run("잘못 입력된 경우", func(t *testing.T) {
		t.Run("/ 가 하나 없을 때", func(t *testing.T) {
			s := TestUrlStruct{Url: "http:/www.naver.com"}
			if err := validate.Struct(s); err == nil {
				t.Error("에러")
			}
		})

		t.Run("/ 가 두 개  없을 때", func(t *testing.T) {
			s := TestUrlStruct{Url: "http:www.naver.com"}
			if err := validate.Struct(s); err == nil {
				t.Error("에러")
			}
		})

		t.Run(": 이 없을 때", func(t *testing.T) {
			s := TestUrlStruct{Url: "httpwww.naver.com"}
			if err := validate.Struct(s); err == nil {
				t.Error("에러")
			}
		})

		t.Run("http or https 로 시작하지 않을 때", func(t *testing.T) {
			s := TestUrlStruct{Url: "htts://www.naver.com"}
			if err := validate.Struct(s); err == nil {
				t.Error("에러")
			}
		})
	})
}

type TestPhoneStruct struct {
	Phone string `validate:"customPhone"`
}

func TestPhone(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("customPhone", Phone())

	t.Run("정상 입력된 경우", func(t *testing.T) {
		t.Run("010", func(t *testing.T) {
			s := TestPhoneStruct{Phone: "01043226633"}
			if err := validate.Struct(s); err != nil {
				t.Error("에러")
			}
		})

		t.Run("011", func(t *testing.T) {
			s := TestPhoneStruct{Phone: "01143226633"}
			if err := validate.Struct(s); err != nil {
				t.Error("에러")
			}
		})

		t.Run("10자리", func(t *testing.T) {
			s := TestPhoneStruct{Phone: "0116338517"}
			if err := validate.Struct(s); err != nil {
				t.Error("에러")
			}
		})
	})
	t.Run("잘못 입력된 경우", func(t *testing.T) {
		t.Run("문자열", func(t *testing.T) {
			s := TestPhoneStruct{Phone: "0104322663a"}
			if err := validate.Struct(s); err == nil {
				t.Error("에러")
			}
		})

		t.Run("특수문자", func(t *testing.T) {
			s := TestPhoneStruct{Phone: "0114322663!"}
			if err := validate.Struct(s); err == nil {
				t.Error("에러")
			}
		})

		t.Run("13자리", func(t *testing.T) {
			s := TestPhoneStruct{Phone: "010432266333"}
			if err := validate.Struct(s); err == nil {
				t.Error("에러")
			}
		})

		t.Run("10자리", func(t *testing.T) {
			s := TestPhoneStruct{Phone: "010432266"}
			if err := validate.Struct(s); err == nil {
				t.Error("에러")
			}
		})
	})
}
