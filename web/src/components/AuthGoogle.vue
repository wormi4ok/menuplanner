<template>
  <b-button id="loginButton" class="is-success" :loading="loading" @click="onSignIn">
    <span>Sign In with Google</span>
  </b-button>
</template>

<script>
import { getAuthCode } from '@/auth/google';

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
        await this.$store.dispatch('googleLogIn', authCode);
        await this.$router.push('/');
      } catch (e) {
        await this.$store.dispatch('reportError', e.response.data);
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
