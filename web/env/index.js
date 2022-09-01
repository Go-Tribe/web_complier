import devEnv from './dev.env';
import proEnv from './pro.env';

// NODE_ENV='development' 开发环境
// NODE_ENV='production' 生产环境
export const isProd = () => (process.env.NODE_ENV === 'production');
const env = isProd() ? proEnv : devEnv;

export default env;