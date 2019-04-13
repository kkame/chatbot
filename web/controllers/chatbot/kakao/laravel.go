package kakao

const baseUrl = "https://laravel.kr/docs"

func laravel(sender string, msg string) (responseText string) {

	switch msg {

	case "@헬프":

		responseText = "라라벨 사이트를 조금 더 쉽게 쓰기 위한 자동응답 봇입니다\n 기본적으로 @헬프를 통해 관련 메뉴를 볼 수 있으며 '@구글 {키워드}' 를 통해 자동 구글링 링크. @메뉴얼 과 @{라라벨버전} @{라라벨 메뉴얼의 LNB 메뉴명} 등을 통해 해당하는 링크를 반환합니다"

	case "@메뉴얼":
		fallthrough
	case "@매뉴얼":

		responseText = baseUrl + ""
	case "@5.7":

		responseText = baseUrl + "/5.7"

	case "@5.6":

		responseText = baseUrl + "/5.6"

	case "@5.5":

		responseText = baseUrl + "/5.5"

	case "@5.4":

		responseText = baseUrl + "/5.4"

	case "@5.3":

		responseText = baseUrl + "/5.3"

	case "@5.2":

		responseText = baseUrl + "/5.2"

	case "@5.1":

		responseText = baseUrl + "/5.1"
	case "@설치하기":
		fallthrough

	case "@인스톨":

		responseText = baseUrl + "/installation"

	case "@업그레이드":

		responseText = baseUrl + "/upgrade"

	case "@라이프사이클":

		responseText = baseUrl + "/lifecycle"

	case "@컨테이너":
		fallthrough

	case "@서비스컨테이너":

		responseText = baseUrl + "/container"

	case "@프로바이더":
		fallthrough

	case "@서비스프로바이더":

		responseText = baseUrl + "/providers"

	case "@파사드":

		responseText = baseUrl + "/facades"

	case "@컨트랙":
		fallthrough

	case "@컨트랙트":

		responseText = baseUrl + "/contracts"

	case "@라우팅":

		responseText = baseUrl + "/routing"

	case "@미들웨어":

		responseText = baseUrl + "/midleware"

	case "@csrf":

		responseText = baseUrl + "/csrf"

	case "@컨트롤러":

		responseText = baseUrl + "/controllers"

	case "@리퀘스트":

		responseText = baseUrl + "/requests"

	case "@리스폰스":

		responseText = baseUrl + "/responses"

	case "@뷰":

		responseText = baseUrl + "/views"

	case "@url":
		fallthrough

	case "@urls":

		responseText = baseUrl + "/urls"

	case "@세션":

		responseText = baseUrl + "/sessions"

	case "@유효성검사":

		fallthrough
	case "@밸리데이션":

		fallthrough
	case "@벨리데이션":

		responseText = baseUrl + "/validation"

	case "@에러핸들링":
		fallthrough

	case "@에러":

		responseText = baseUrl + "/errors"

	case "@로깅":

		responseText = baseUrl + "/logging"

	case "@블레이드":

		responseText = baseUrl + "/blade"

	case "@다국어":

		fallthrough
	case "@국제화":

		responseText = baseUrl + "/localization"

	case "@프론트엔드":

		fallthrough
	case "@스케폴딩":

		responseText = baseUrl + "/frontend"

	case "@mix":

		fallthrough
	case "@믹스":
		fallthrough

	case "@assets":

		responseText = baseUrl + "/mix"

	case "@인증":

		responseText = baseUrl + "/authentication"

	case "@인가":

		fallthrough
	case "@승인":

		responseText = baseUrl + "/authorization"

	case "@패스포트":

		fallthrough
	case "@passport":

		fallthrough
	case "@oauth":
		fallthrough

	case "@oauth2":

		responseText = baseUrl + "/passport"

	case "@검증":
		fallthrough

	case "@이메일":

		responseText = baseUrl + "/verification"

	case "@암호화":

		responseText = baseUrl + "/encription"

	case "@해싱":

		responseText = baseUrl + "/hashing"

	case "@패스워드재설정":
		fallthrough

	case "@재설정":
		fallthrough

	case "@패스워드":

		responseText = baseUrl + "/passwords"

	case "@아티즌 콘솔":

		responseText = baseUrl + "/artisan"

	case "@브로드캐스팅":

		responseText = baseUrl + "/broadcasting"

	case "@캐시":

		responseText = baseUrl + "/cache"

	case "@컬렉션":
		fallthrough

	case "@collection":
		fallthrough

	case "@collections":

		responseText = baseUrl + "/collections"

	case "@이벤트":

		fallthrough
	case "@event":

		fallthrough
	case "@events":

		responseText = baseUrl + "/events"

	case "@파일 스토리지":

		fallthrough
	case "@파일":

		fallthrough
	case "@파일시스템":

		fallthrough
	case "@filesystem":

		responseText = baseUrl + "/filesystem"

	case "@헬퍼 함수들":

		fallthrough
	case "@함수들":
		fallthrough

	case "@헬퍼":

		responseText = baseUrl + "/helpers"

	case "@메일":

		fallthrough
	case "@mail":

		responseText = baseUrl + "/mail"

	case "@notifications-알림":

		fallthrough
	case "@알림":
		fallthrough

	case "@notifications":

		responseText = baseUrl + "/notifications"

	case "@패키지 개발":
		fallthrough

	case "@패키지":

		responseText = baseUrl + "/packages"

	case "@queues-큐":
		fallthrough

	case "@큐":
		fallthrough

	case "@queues":

		responseText = baseUrl + "/queues"

	case "@작업 스케줄링":

		fallthrough
	case "@스케줄링":

		fallthrough
	case "@scheduling":
		responseText = baseUrl + "/scheduling"
	case "@데이터베이스":
		fallthrough

	case "@database":

		responseText = baseUrl + "/database"

	case "@쿼리 빌더":
		fallthrough

	case "@쿼리":
		fallthrough

	case "@queries":
		fallthrough

	case "@query":

		responseText = baseUrl + "/queries"

	case "@페이지네이션":
		fallthrough

	case "@pagination":

		fallthrough
	case "@페이징":

		responseText = baseUrl + "/pagination"

	case "@마이그레이션":
		fallthrough

	case "@migration":

		fallthrough
	case "@migrations":

		responseText = baseUrl + "/migrations"

	case "@seeding-시딩":
		fallthrough

	case "@시딩":
		fallthrough

	case "@seeding":

		responseText = baseUrl + "/seeding"

	case "@redis-레디스":
		fallthrough

	case "@레디스":

		fallthrough
	case "@redis":

		responseText = baseUrl + "/redis"

	case "@엘로퀀트":
		fallthrough

	case "@엘로퀸트":

		fallthrough
	case "@엘퀸":
		fallthrough

	case "@orm":
		fallthrough

	case "@eloquent":

		responseText = baseUrl + "/eloquent"

	case "@relationships-관계":
		fallthrough

	case "@관계":
		fallthrough

	case "@relationships":

		responseText = baseUrl + "/eloquent-relationships"

	case "@eloquent-collections":

		responseText = baseUrl + "/eloquent-collections"

	case "@mutators":

		responseText = baseUrl + "/eloquent-mutators"

	case "@eloquent resources":
	case "@eloquent-resources":
		fallthrough
	case "@api resources":

		responseText = baseUrl + "/eloquent-resources"

	case "@serialization":

		responseText = baseUrl + "/eloquent-serialization"
	case "@테스트":
		fallthrough

	case "@테스팅":
		fallthrough

	case "@testing":

		responseText = baseUrl + "/testing"

	case "@http테스트":

		responseText = baseUrl + "/http-tests"

	case "@콘솔테스트":

		responseText = baseUrl + "/console-tests"

	case "@브라우저테스트":

		fallthrough
	case "@dusk":

		responseText = baseUrl + "/dusk"

	case "@데이터베이스테스트":

		responseText = baseUrl + "/database-testing"

	case "@모킹":

		fallthrough
	case "@목킹":

		fallthrough
	case "@mocking":

		responseText = baseUrl + "/mocking"
	case "@캐셔":

		responseText = baseUrl + "/billing"

	case "@envoy-task runner":

		fallthrough
	case "@envoy":
		fallthrough

	case "@엔보이":

		responseText = baseUrl + "/envoy"

	case "@horizon":

		fallthrough
	case "@호라이즌":

		responseText = baseUrl + "/horizon"

	case "@scout-검색":

		fallthrough
	case "@scout":
		fallthrough

	case "@스카우트":

		responseText = baseUrl + "/scout"

	case "@socialite-소셜로그인":

		fallthrough
	case "@socialite":
		fallthrough

	case "@소셜로그인":

		fallthrough
	case "@소셜라이트":

		responseText = baseUrl + "/socialite"

	case "@telescope":
		fallthrough

	case "@텔레스코프":

		responseText = baseUrl + "/telescope"
	case "@홈스테드":

		fallthrough
	case "@homestead":

		responseText = baseUrl + "/homestead"

	case "@발렛":

		fallthrough
	case "@valet":

		responseText = baseUrl + "/valet"
	default:
		responseText = ""

	}
	//toast('여기까진 실행됨');
	return responseText
}
