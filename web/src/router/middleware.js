import $store from '@/store';
import token from '@/auth/token';

export default {
  async initUser(to, from, next) {
    if (token.getRefresh() && !$store.getters.isLoggedIn) {
      try {
        if (token.isExpired()) {
          await $store.dispatch('refreshToken');
        }
        await $store.dispatch('fetchCurrentUser');
        next();
      } catch (e) {
        await $store.dispatch('reportError', 'Authentication failed');
      }
    } else {
      next();
    }
  },
  checkAccess(to, from, next) {
    const isAuthRoute = to.matched.some((item) => item.meta.isAuth);

    if (isAuthRoute && !$store.getters.isLoggedIn) return next({ name: 'Login' });
    return next();
  },
};
