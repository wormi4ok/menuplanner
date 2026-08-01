import token from '@/auth/token';
import { useUserStore } from '@/stores/user';
import { useErrorStore } from '@/stores/error';

export default {
  async initUser(to, from, next) {
    const user = useUserStore();

    if (token.getRefresh() && !user.isLoggedIn) {
      try {
        if (token.isExpired()) {
          await user.refreshToken();
        }
        await user.fetchCurrentUser();
        next();
      } catch (e) {
        useErrorStore().reportError('Authentication failed');
        next();
      }
    } else {
      next();
    }
  },
  checkAccess(to, from, next) {
    const isAuthRoute = to.matched.some((item) => item.meta.isAuth);

    if (isAuthRoute && !useUserStore().isLoggedIn) return next({ name: 'Login' });
    return next();
  },
};
