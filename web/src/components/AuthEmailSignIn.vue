<template>
  <form class="block" @submit.prevent="signIn">
    <b-notification
      v-if="error"
      type="is-danger is-light"
      aria-close-label="Close error message"
      role="alert">
      {{ error }}
    </b-notification>
    <b-field label="Email">
      <b-input
        v-model="email"
        icon="envelope"
        type="email"
        placeholder="e.g. natalia@flaningo.ru"
        required
      />
    </b-field>
    <b-field label="Password">
      <b-input
        v-model="password"
        type="password"
        icon="lock"
        placeholder="*******"
        required
        password-reveal
      />
    </b-field>
    <b-button class="is-primary" label="Continue" native-type="submit"/>
  </form>
</template>

<script>
export default {
  name: 'AuthEmailSignIn',
  data() {
    return {
      email: '',
      password: '',
      error: '',
    };
  },
  methods: {
    async signIn() {
      this.error = '';
      const { email, password } = this;
      try {
        await this.$store.dispatch('logIn', { email, password });
        await this.$router.push('/');
      } catch (e) {
        this.error = e.response.data;
      }
    },
  },
};
</script>

<style scoped>

</style>
