package config

type Env string

const (
	Pro  Env = "pro"  // 生产环境
	Test Env = "test" // 测试环境
	Dev  Env = "dev"  // 开发环境
)

var EnvDesc = map[Env]string{
	Pro:  "生产环境",
	Test: "测试环境",
	Dev:  "本地开发环境",
}

func IsEnv(env string) (Env, bool) {
	for k := range EnvDesc {
		if string(k) == env {
			return k, true
		}
	}
	return "", false
}
