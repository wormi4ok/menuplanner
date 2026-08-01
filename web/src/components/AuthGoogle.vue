<template>
  <b-button id="loginButton" class="is-success" :loading="loading" @click="onSignIn">
    <span>Sign In with Google</span>
  </b-button>
</template>

<script>
import { getAuthCode } from '@/auth/google';
import { useUserStore } from '@/stores/user';
import { useErrorStore } from '@/stores/error';

export default {
  name: 'AuthGoogle',
  data() {
    return {
      loading: false,
    };
  },
  methods: {
    async onSignIn() {
      try {
        this.loading = true;
        const authCode = await getAuthCode();
        await useUserStore().googleLogIn(authCode);
        await this.$router.push('/');
      } catch (e) {
        useErrorStore().reportError(e.response.data);
      } finally {
        this.loading = false;
      }
    },
  },
};
</script>

<style scoped>
#loginButton {
  width: 100%;
}
</style>
