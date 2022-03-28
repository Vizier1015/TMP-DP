package packed

import "github.com/gogf/gf/os/gres"

func init() {
	if err := gres.Add("H4sIAAAAAAAC/wrwZmYRYeBg4GA4aKsawoAERBg4GYrLE9PTU4v0obReVnF+XmgIKwPjLp+riaueBQUEBARo6fn7+PvrBm0KCjC6ZBxwqbEj7dnSp9qRmg5XGlwnff7sJTTZ26Fpa8LFqBWXmo5M6knqKNBVYGD4/z/Am53DrHPCQgcGBgaQzTCnMDAcQnMKO8IpYNv3+FxNBOlGVhPgzcgkwozwCrLJIK/AwP9GEInXYwijsDsFAgQY/jt2I4xCchgrG0ieiYGJoZOBgeEkWDUgAAD//37tiYRoAQAA"); err != nil {
		panic("add binary content to resource manager failed: " + err.Error())
	}
}
