<template>
  <b-button id="loginButton" class="is-success" v-if="!isSignIn" @click="onSignIn">
    <span>Sign In with Google</span>
  </b-button>
  <b-button v-else @click="onSignOut">
    <span>Sign Out</span>
  </b-button>
</template>

<script>
export default {
  name: 'AuthGoogle',
  data() {
    return {
      isSignIn: false,
    };
  },
  methods: {
    async onSignIn() {
      try {
        const authCode = await this.$gAuth.getAuthCode();
        console.log('authCode', authCode);
        this.isSignIn = this.$gAuth.isAuthorized;
      } catch (error) {
        // On fail do something
        console.error(error);
      }
    },
    onSignOut() {
      this.isSignIn = false;
    },
  },
};
</script>

<style scoped>
#loginButton {
  width: 100%;
}
</style>
