import { defineStore } from 'pinia';
import api from '@/api';
import token from '@/auth/token';
import { useRecipesStore } from './recipes';
import { useWeekStore } from './week';

export const useUserStore = defineStore('user', {
  state: () => ({
    user: '',
  }),
  getters: {
    isLoggedIn: (state) => !!state.user,
    currentUser: (state) => state.user,
  },
  actions: {
    setAuthTokens(payload) {
      token.set(payload.access_token, payload.expires_in);
      token.setRefresh(payload.refresh_token);
    },
    async logIn({ email, password }) {
      const response = await api.auth.login(email, password);
      this.setAuthTokens(response.data);
    },
    async googleLogIn(authCode) {
      const response = await api.auth.loginViaGoogle(authCode);
      this.setAuthTokens(response.data);
    },
    async signUp({ email, password, passwordConfirm }) {
      const response = await api.auth.signup(email, password, passwordConfirm);
      this.setAuthTokens(response.data);
    },
    async fetchCurrentUser() {
      const response = await api.user.profile();
      this.user = response.data;
    },
    async refreshToken() {
      const response = await api.auth.tokenRefresh();
      this.setAuthTokens(response.data);
    },
    logOut() {
      this.setAuthTokens({ access_token: '', refresh_token: '', expires_in: 0 });
      this.user = '';
      useRecipesStore().$reset();
      useWeekStore().resetCurrentWeek();
    },
  },
});

export default useUserStore;
